package pipeline

import (
	"fmt"
	"io"
	"os"
)

// PrintResultsWithWriter prints all statistics to a given writer (e.g., os.Stdout)
func PrintResultsWithWriter(avg, median, variance, stddev int, w io.Writer) {
	fmt.Fprintf(w, "Average: %d\n", avg)
	fmt.Fprintf(w, "Median: %d\n", median)
	fmt.Fprintf(w, "Variance: %d\n", variance)
	fmt.Fprintf(w, "Standard Deviation: %d\n", stddev)
}

// PrintResults prints all statistics to standard output
func PrintResults(avg, median, variance, stddev int) {
	PrintResultsWithWriter(avg, median, variance, stddev, os.Stdout)
}
