package pipeline

import "math"

// CalculateStdDev takes a variance and returns the rounded standard deviation as int
func CalculateStdDev(variance int) int {
	if variance == 0 {
		return 0 // Stddev of 0 variance is 0
	}

	stddev := math.Sqrt(float64(variance)) // Square root
	return int(stddev + 0.5)               // Round to nearest int
}
