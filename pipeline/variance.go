package pipeline

// CalculateVariance takes a slice of ints and the mean, returns variance as int
func CalculateVariance(numbers []int, mean int) int {
	if len(numbers) == 0 {
		return 0 // Return 0 if slice is empty
	}

	sum := 0
	for _, num := range numbers {
		diff := num - mean
		sum += diff * diff // Square difference
	}

	variance := float64(sum) / float64(len(numbers)) // Compute variance
	return int(variance + 0.5)                       // Round to nearest int
}
