package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Vector struct {
	data []float64

	dim int
}

func main() {

	dataset := readCsvIntoSlice("dataset.csv")
	questions := readCsvIntoSlice("questions.csv")

	for _, question := range questions {
		for _, sentence := range dataset {
			cosine_similarity(question, sentence)
		}
	}
}

func cosine_similarity(Vector, Vector) float64 {
	return 0
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
		dim:  len(record),
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
