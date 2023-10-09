package main

import (
	"fmt"
)

func makeIntRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func makeFloatRange(min, max float64) []float64 {
	a := make([]float64, int(max)-int(min)+1)
	for i := range a {
		a[i] = min + float64(i)
	}
	return a
}

type Number interface {
	int | int64 | float32 | float64
}

func sumNumbers[n Number](nums []n) n {
	var holder n
	for _, num := range nums {
		holder += num
	}
	return holder
}

func sumInts(nums []int) int {
	var holder int
	for _, num := range nums {
		holder += num
	}
	return holder
}

func sumFloat64(nums []float64) float64 {
	var holder float64
	for _, num := range nums {
		holder += num
	}
	return holder
}

func main() {
	intsToAdd := makeIntRange(1, 10000000)
	floatsToAdd := makeFloatRange(1, 10000000)

	fmt.Println("ints", sumNumbers(intsToAdd))
	fmt.Println("floats", sumNumbers(floatsToAdd))
}
