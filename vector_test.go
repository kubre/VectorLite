package main

import (
	"testing"
)

func TestDotProductFor2Vectors(t *testing.T) {
	v := Vector{
		data: []float64{1, 2, 3},
	}
	w := Vector{
		data: []float64{4, 5, 6},
	}

	dotProduct := v.dot(w)

	if dotProduct != 32 {
		t.Fatal("Dot product of", v.data, "and", w.data, "should be 3")
	}
}

func TestEmptyVectors(t *testing.T) {
	v := Vector{}
	w := Vector{}

	dotProduct := v.dot(w)

	if dotProduct != 0 {
		t.Fatal("Dot product of", v.data, "and", w.data, "should be 3")
	}
}

func TestNormAVector(t *testing.T) {
	v := Vector{
		data: []float64{1, 2, 3},
	}

	normal := v.norm()

	if normal != 3.7416573867739413 {
		t.Fatal("Norm of vector", v.data, "should be", 3.7416573867739413, "but it is", normal)
	}
}
