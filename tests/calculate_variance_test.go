package tests

import (
	"math-skills/pipeline" // Import the pipeline package
	"testing"              // Go's standard testing package
)

// TestCalculateVariance_PositiveNumbers checks variance for positive numbers
func TestCalculateVariance_PositiveNumbers(t *testing.T) {
	numbers := []int{10, 20, 30, 40, 50}
	mean := pipeline.CalculateAverage(numbers) // Use the average we already implemented
	expected := 200                            // Variance: ((10-30)^2 + (20-30)^2 + ...)/5 = 200

	variance := pipeline.CalculateVariance(numbers, mean) // Call the function

	if variance != expected {
		t.Errorf("Expected variance %d, got %d", expected, variance)
	}
}

// TestCalculateVariance_NegativeNumbers checks variance for negative numbers
func TestCalculateVariance_NegativeNumbers(t *testing.T) {
	numbers := []int{-10, 0, 10}
	mean := pipeline.CalculateAverage(numbers) // Should be 0
	expected := 67                             // Variance: ((-10-0)^2 + (0-0)^2 + (10-0)^2)/3 = 66.66 → rounded 67

	variance := pipeline.CalculateVariance(numbers, mean)

	if variance != expected {
		t.Errorf("Expected variance %d, got %d", expected, variance)
	}
}

// TestCalculateVariance_EmptySlice checks how variance handles an empty slice
func TestCalculateVariance_EmptySlice(t *testing.T) {
	numbers := []int{}
	mean := 0
	expected := 0 // Convention: empty slice returns 0

	variance := pipeline.CalculateVariance(numbers, mean)

	if variance != expected {
		t.Errorf("Expected variance %d, got %d", expected, variance)
	}
}
