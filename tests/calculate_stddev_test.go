package tests

import (
	"math"
	"math-skills/pipeline" // Import the pipeline package
	"testing"              // Go's standard testing package
)

// TestCalculateStdDev_PositiveVariance checks standard deviation for a positive variance
func TestCalculateStdDev_PositiveVariance(t *testing.T) {
	variance := 200
	expected := int(math.Sqrt(200) + 0.5) // Rounded square root of 200

	stddev := pipeline.CalculateStdDev(variance) // Call the function

	if stddev != expected {
		t.Errorf("Expected stddev %d, got %d", expected, stddev)
	}
}

// TestCalculateStdDev_ZeroVariance checks standard deviation when variance is 0
func TestCalculateStdDev_ZeroVariance(t *testing.T) {
	variance := 0
	expected := 0 // Convention: stddev of 0 variance is 0

	stddev := pipeline.CalculateStdDev(variance)

	if stddev != expected {
		t.Errorf("Expected stddev %d, got %d", expected, stddev)
	}
}

// TestCalculateStdDev_SmallVariance checks standard deviation for a small variance
func TestCalculateStdDev_SmallVariance(t *testing.T) {
	variance := 2
	expected := int(math.Sqrt(2) + 0.5) // Rounded square root of 2

	stddev := pipeline.CalculateStdDev(variance)

	if stddev != expected {
		t.Errorf("Expected stddev %d, got %d", expected, stddev)
	}
}
