package tests

import (
	"math-skills/pipeline" // Import our pipeline package where ReadFile() is defined
	"testing"              // Go's testing package for writing unit tests
)

// TestReadFileSuccess checks that ReadFile correctly reads numbers from a valid file
func TestReadFileSuccess(t *testing.T) {
	path := "../data/data1.txt" // Path to a sample file with numbers

	expected := []int{10, 20, 30, 40, 50} // What we expect to get from ReadFile

	numbers, err := pipeline.ReadFile(path) // Call ReadFile to read numbers from the file

	// Check if there was an unexpected error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the number of numbers read matches the expected number
	if len(numbers) != len(expected) {
		t.Fatalf("Expected %d numbers, got %d", len(expected), len(numbers))
	}

	// Compare each number to the expected value
	for i, v := range expected {
		if numbers[i] != v {
			t.Errorf("Expected numbers[%d] = %d, got %d", i, v, numbers[i])
		}
	}
}

// TestReadFileNotFound checks that ReadFile returns an error when the file does not exist
func TestReadFileNotFound(t *testing.T) {
	path := "../data/does_not_exist.txt" // Path to a file that does NOT exist

	_, err := pipeline.ReadFile(path) // Try to read a non-existent file

	// The function should return an error in this case
	if err == nil {
		t.Fatal("Expected an error for non-existent file, got nil")
	}
}
