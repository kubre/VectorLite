package main

import "math"

type Vector struct {
	data   []float64
	normal float64
}

func (v Vector) cosine_similarity(w Vector) float64 {
	return v.dot(w) / (v.norm() * w.norm())
}

func (v Vector) dot(w Vector) float64 {
	sum := 0.0
	for i := 0; i < len(v.data); i++ {
		sum += v.data[i] * w.data[i]
	}
	return sum
}

func (v Vector) norm() float64 {
	sum := 0.0
	for _, vi := range v.data {
		sum += vi * vi
	}

	// TODO: Need to check how much perf gain is to return cached version
	return math.Sqrt(sum)
}
