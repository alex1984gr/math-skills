package tests

import (
	"math-skills/pipeline" // Import the pipeline package
	"testing"              // Go's testing package
)

// TestCalculateMedian_OddLength checks median for a slice with odd number of elements
func TestCalculateMedian_OddLength(t *testing.T) {
	numbers := []int{10, 30, 20} // Unsorted slice with odd length
	expected := 20               // Median after sorting [10,20,30] is 20

	median := pipeline.CalculateMedian(numbers) // Call the function

	if median != expected {
		t.Errorf("Expected median %d, got %d", expected, median) // Fail if result is wrong
	}
}

// TestCalculateMedian_EvenLength checks median for a slice with even number of elements
func TestCalculateMedian_EvenLength(t *testing.T) {
	numbers := []int{10, 40, 30, 20} // Unsorted slice with even length
	expected := 25                   // Median: average of middle two after sorting [10,20,30,40] → (20+30)/2=25

	median := pipeline.CalculateMedian(numbers) // Call the function

	if median != expected {
		t.Errorf("Expected median %d, got %d", expected, median)
	}
}

// TestCalculateMedian_EmptySlice checks how median handles an empty slice
func TestCalculateMedian_EmptySlice(t *testing.T) {
	numbers := []int{} // Empty slice
	expected := 0      // Convention: empty slice returns 0

	median := pipeline.CalculateMedian(numbers) // Call the function

	if median != expected {
		t.Errorf("Expected median %d, got %d", expected, median)
	}
}
