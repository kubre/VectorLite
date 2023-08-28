package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Output struct {
	text string
	rank float32
}

func main() {
	dataset := readCsvIntoRecords("sentences.csv")
	questions := readCsvIntoRecords("questions.csv")

	rankings := make([][]Output, len(questions))
	for i, question := range questions {
		temp := make([]Output, len(dataset))
		for j, sentence := range dataset {
			score := question.cosine_similarity(sentence)
			temp[j] = Output{text: sentence.metadata["text"].(string), rank: float32(score)}
		}
		rankings[i] = temp
	}

	// Sorting for testing
	for _, ranks := range rankings {
		sort.Slice(ranks, func(i, j int) bool {
			return ranks[i].rank > ranks[j].rank
		})
	}

	fmt.Println(rankings)
}

func readCsvIntoRecords(name string) []Record {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Cannot read the file", err)
	}

	r := csv.NewReader(file)
	embeddings := []Record{}

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		tempEmbedding := readRecord(row)

		if err != nil {
			log.Fatal("Error reading file", err)
		}

		embeddings = append(embeddings, tempEmbedding)
	}

	if err := file.Close(); err != nil {
		log.Fatal("Not able to close the file")
	}

	return embeddings
}

func readRecord(row []string) Record {
	tempEmbedding := Record{
		index:    make([]float64, len(row)-1),
		metadata: map[string]interface{}{"text": row[0]},
	}

	for i, value := range row[1:] {
		point, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		tempEmbedding.index[i] = point
	}
	return tempEmbedding
}
