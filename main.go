package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Vector struct {
	data []float64

	// dim int
}

func (v Vector) cosine_similarity(w Vector) float64 {
	// return v.dot(w).divide(v.norm().multiply(w.norm()))
	return v.dot(w)
}

func (v Vector) dot(w Vector) float64 {
	sum := 0.0
	for i := 0; i < len(v.data); i++ {
		sum += v.data[i] * w.data[i]
	}
	return sum
}

func main() {

	dataset := readCsvIntoSlice("dataset.csv")
	questions := readCsvIntoSlice("questions.csv")

	rankings := make([][]float64, len(questions))
	for _, question := range questions {
		temp := make([]float64, len(dataset))
		for _, sentence := range dataset {
			score := question.cosine_similarity(sentence)
			temp = append(temp, score)
		}
		rankings = append(rankings, temp)
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
		// dim:  len(record),
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
