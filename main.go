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

	data := readCsvIntoSlice("test_data.csv")

	fmt.Println(data[0][0:4])

}

func readCsvIntoSlice(name string) [][]float64 {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Cannot read the file", err)
	}

	r := csv.NewReader(file)
	embeddings := [][]float64{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		tempEmbedding := convertStringListToFloats(record, 384)

		if err != nil {
			log.Fatal("Error reading file", err)
		}

		embeddings = append(embeddings, tempEmbedding[:])
	}

	if err := file.Close(); err != nil {
		log.Fatal("Not able to close the file")
	}
	return embeddings
}

func convertStringListToFloats(record []string, dimension int) []float64 {
	tempEmbedding := make([]float64, dimension)
	for i, value := range record {
		point, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}
		tempEmbedding[i] = point
	}
	return tempEmbedding
}
