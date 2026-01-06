package tests

import (
	"math-skills/pipeline" // Import our pipeline package where CalculateAverage() is defined
	"testing"              // Go's standard testing package
)

// TestCalculateAverage_PositiveNumbers checks that CalculateAverage correctly computes the average of positive numbers
func TestCalculateAverage_PositiveNumbers(t *testing.T) {
	numbers := []int{10, 20, 30, 40, 50} // Define a slice of positive numbers
	expected := 30                       // The expected average: (10+20+30+40+50)/5 = 30

	avg := pipeline.CalculateAverage(numbers) // Call the function we are testing

	// Check if the result matches the expected value
	if avg != expected {
		t.Errorf("Expected average %d, got %d", expected, avg) // Report error if test fails
	}
}

// TestCalculateAverage_NegativeNumbers checks that CalculateAverage works with negative numbers
func TestCalculateAverage_NegativeNumbers(t *testing.T) {
	numbers := []int{-10, 0, 10} // Slice contains negative, zero, and positive numbers
	expected := 0                // The expected average: (-10+0+10)/3 = 0

	avg := pipeline.CalculateAverage(numbers) // Call the function

	// Check if the result matches the expected value
	if avg != expected {
		t.Errorf("Expected average %d, got %d", expected, avg)
	}
}

// TestCalculateAverage_EmptyList checks how CalculateAverage handles an empty slice
func TestCalculateAverage_EmptyList(t *testing.T) {
	numbers := []int{} // Empty slice
	expected := 0      // Convention: empty list returns 0

	avg := pipeline.CalculateAverage(numbers) // Call the function

	// Check if the result matches the expected value
	if avg != expected {
		t.Errorf("Expected average %d, got %d", expected, avg)
	}
}
