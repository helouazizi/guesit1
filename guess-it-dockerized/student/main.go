package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		println("Error: Usage is: ", "go run .")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	slices := []float64{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			println("Error: unable to parse line: ", line)
			continue
		}
		slices = append(slices, num)
		if len(slices) > 1 {
			min, max := Guess_it(slices)
			fmt.Printf("%.0f %.0f\n", min, max)
		}

	}
	if err := scanner.Err(); err != nil {
		println("Error: unable to read from stdin: ", err)
		return
	}
}

func Average(data []float64) float64 {
	var sum float64
	length := len(data)
	var average float64
	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	average = sum / float64(length)
	return average
}

func Variance(data []float64) float64 {
	var mean float64
	var variance float64
	length := len(data)
	var sub float64
	var sq float64

	mean = Average(data)
	for _, n := range data {
		sub = n - mean
		sq += sub * sub
	}
	variance = sq / float64(length)
	return variance
}

func Standard_Deviation(data []float64) float64 {
	deviation := math.Sqrt(Variance(data))
	return deviation
}

func Guess_it(data []float64) (float64, float64) {
	var min float64
	var max float64

	start := len(data) - 4
	if start < 0 {
		start = 0
	}
	preciseData := data[start:]
	average := Average(preciseData)
	sd := Standard_Deviation(preciseData)

	min = average - sd
	max = average + sd

	return min, max
}
