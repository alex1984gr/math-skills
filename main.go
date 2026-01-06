package main

import (
	"fmt"
	"log"
	"math-skills/pipeline"
	"os"
)

func main() {
	// Check if the user provided a file path as a command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <data_file>")
		return
	}

	path := os.Args[1] // Get the path to the data file from command line

	// 1️⃣ Read numbers from file
	numbers, err := pipeline.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 2️⃣ Calculate average
	avg := pipeline.CalculateAverage(numbers)

	// 3️⃣ Calculate median
	median := pipeline.CalculateMedian(numbers)

	// 4️⃣ Calculate variance (requires mean)
	variance := pipeline.CalculateVariance(numbers, avg)

	// 5️⃣ Calculate standard deviation
	stddev := pipeline.CalculateStdDev(variance)

	// 6️⃣ Print all results
	pipeline.PrintResults(avg, median, variance, stddev)
}
