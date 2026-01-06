package pipeline

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads integers from a file line by line and returns them as a slice of ints
func ReadFile(path string) ([]int, error) {
	file, err := os.Open(path) // Open the file
	if err != nil {
		return nil, err // Return error if file cannot be opened
	}
	defer file.Close() // Ensure file is closed when function ends

	var numbers []int                 // Slice to store numbers
	scanner := bufio.NewScanner(file) // Create a scanner to read lines
	for scanner.Scan() {              // Loop through each line
		line := strings.TrimSpace(scanner.Text()) // Remove spaces
		if line == "" {                           // Skip empty lines
			continue
		}
		num, err := strconv.Atoi(line) // Convert line to int
		if err != nil {
			return nil, err // Return error if conversion fails
		}
		numbers = append(numbers, num) // Add number to slice
	}

	if err := scanner.Err(); err != nil { // Check for scanning errors
		return nil, err
	}

	return numbers, nil // Return the slice of numbers
}
