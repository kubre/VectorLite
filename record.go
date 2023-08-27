package main

import "math"

type Record struct {
	index  []float64
	normal float64
}

func (r Record) cosine_similarity(w Record) float64 {
	return r.dot(w) / (r.norm() * w.norm())
}

func (r Record) dot(w Record) float64 {
	sum := 0.0
	for i := 0; i < len(r.index); i++ {
		sum += r.index[i] * w.index[i]
	}
	return sum
}

func (r Record) norm() float64 {
	sum := 0.0
	for _, elem := range r.index {
		sum += elem * elem
	}

	// TODO: Need to check how much perf gain is to return cached version
	return math.Sqrt(sum)
}
