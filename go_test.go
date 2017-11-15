package main

import (
	"math/rand"
	"testing"
)

const (
	dataSize = 1e6 // 1mil
)

func Benchmark1milSort(b *testing.B) {
	// To be fair to the node test, we will use float64
	// rather than faster integers
	data := make([]float64, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = rand.Float64()*1000 + 1
	}
	greater := make([]float64, dataSize)
	lesser := make([]float64, dataSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var numGreater int
		var numLesser int

		for _, v := range data {
			if v < 500 {
				lesser[numLesser] = v
				numLesser++
			} else if v > 500 {
				greater[numGreater] = v
				numGreater++
			}
		}

		//	lesser = lesser[:numLesser]
		//	greater = greater[:numGreater]
	}
}

func Benchmark1milSortOnlyIntegers(b *testing.B) {
	data := make([]uint64, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = rand.Uint64()*1000 + 1
	}
	greater := make([]uint64, dataSize)
	lesser := make([]uint64, dataSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var numGreater int
		var numLesser int

		for _, v := range data {
			if v < 500 {
				lesser[numLesser] = v
				numLesser++
			} else if v > 500 {
				greater[numGreater] = v
				numGreater++
			}
		}
		// lesser = lesser[:numLesser]
		// greater = greater[:numGreater]
	}

}
