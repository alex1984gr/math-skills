package tests

import (
	"bytes"
	"math-skills/pipeline"
	"testing"
)

// TestPrintResults checks that print_results outputs the correct format
func TestPrintResults(t *testing.T) {
	var buf bytes.Buffer // Create a buffer to capture printed output

	// Temporarily replace stdout with our buffer
	pipeline.PrintResultsWithWriter(35, 4, 5, 65, &buf)

	expected := "Average: 35\nMedian: 4\nVariance: 5\nStandard Deviation: 65\n"

	output := buf.String() // Read buffer content

	if output != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output)
	}
}
