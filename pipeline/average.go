package pipeline

// CalculateAverage takes a slice of ints and returns the rounded average as int
func CalculateAverage(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Return 0 if slice is empty
	}

	sum := 0 // Sum all numbers
	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(len(numbers)) // Compute average
	return int(average + 0.5)                       // Round to nearest integer
}
