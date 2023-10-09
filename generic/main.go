package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"time"
)

func callerName(skip int) string {
	const unknown = "unknown"
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return unknown
	}
	frame, _ := runtime.CallersFrames(pcs).Next()
	if frame.Function == "" {
		return unknown
	}
	return frame.Function
}

func timer() func() {
	name := callerName(1)
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
		stringDuration := fmt.Sprintf("%f", time.Since(start).Seconds())
		f, err := os.OpenFile("../runs/generic_runs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		w := csv.NewWriter(f)
		w.Write([]string{stringDuration})

		w.Flush()
	}
}

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
	defer timer()()
	var holder n
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
