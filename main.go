package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	dataset := readCsvIntoSlice("dataset.csv")
	questions := readCsvIntoSlice("questions.csv")

	rankings := make([][]float64, len(questions))
	for i, question := range questions {
		temp := make([]float64, len(dataset))
		for j, sentence := range dataset {
			score := question.cosine_similarity(sentence)
			temp[j] = score
		}
		rankings[i] = temp
	}
	fmt.Println(rankings)
}

func readCsvIntoSlice(name string) []Vector {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Cannot read the file", err)
	}

	r := csv.NewReader(file)
	embeddings := []Vector{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		tempEmbedding := convertStringListToFloats(record)

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

func convertStringListToFloats(record []string) Vector {
	tempEmbedding := Vector{
		data: make([]float64, len(record)),
	}

	for i, value := range record {
		point, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		tempEmbedding.data[i] = point
	}
	return tempEmbedding
}
