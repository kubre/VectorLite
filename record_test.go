package main

import (
	"testing"
)

func TestDotProductFor2Vectors(t *testing.T) {
	v := Record{
		index: []float64{1, 2, 3},
	}
	w := Record{
		index: []float64{4, 5, 6},
	}

	dotProduct := v.dot(w)

	if dotProduct != 32 {
		t.Fatal("Dot product of", v.index, "and", w.index, "should be 3")
	}
}

func TestEmptyVectors(t *testing.T) {
	v := Record{}
	w := Record{}

	dotProduct := v.dot(w)

	if dotProduct != 0 {
		t.Fatal("Dot product of", v.index, "and", w.index, "should be 3")
	}
}

func TestNormAVector(t *testing.T) {
	v := Record{
		index: []float64{1, 2, 3},
	}

	normal := v.norm()

	if normal != 3.7416573867739413 {
		t.Fatal("Norm of vector", v.index, "should be", 3.7416573867739413, "but it is", normal)
	}
}
