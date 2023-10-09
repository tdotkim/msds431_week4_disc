package main

import (
	"testing"
)

var intsToAdd = makeIntRange(1, 10000000)
var floatsToAdd = makeFloatRange(1, 10000000)

func BenchmarkSumInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumInts(intsToAdd)
	}
}

func BenchmarkSumFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumFloat64(floatsToAdd)
	}
}

func BenchmarkSumNumInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNumbers(intsToAdd)
	}
}

func BenchmarkSumNumFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNumbers(floatsToAdd)
	}
}
