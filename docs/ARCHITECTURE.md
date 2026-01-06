Architecture of math-skills

This document explains the architecture and design decisions of the math-skills Go project.
The project is implemented in a pipeline structure, separating each stage of computation into its own function for clarity, maintainability, and testability.

1. Overview

math-skills reads a list of integers from a file and calculates four main statistics:

Average – mean of the numbers

Median – middle value after sorting

Variance – mean squared deviation from the average

Standard Deviation – square root of variance

The architecture is designed to:

Follow TDD principles: every function has unit tests

Be modular: each function does one thing only

Be grader-friendly: output matches exactly what the 01‑edu grader expects

Handle file reading errors gracefully

2. Pipeline Structure

The project is divided into the following core pipeline stages:

Stage / Function	Responsibility
ReadFile(path string)	Reads integers from a file into a slice. Handles errors for missing files or invalid data.
CalculateAverage(numbers []int)	Computes the arithmetic mean of the list. Returns 0 for empty lists.
CalculateMedian(numbers []int)	Computes the median. Works for odd and even-length lists.
CalculateVariance(numbers []int, mean int)	Computes the variance (average squared deviation).
CalculateStdDev(variance int)	Computes standard deviation (square root of variance).
PrintResults(avg, median, variance, stddev int)	Prints results in the exact format expected by the grader.

Data flow:

data.txt → ReadFile → CalculateAverage → CalculateMedian
                                   → CalculateVariance → CalculateStdDev
                                               ↓
                                       PrintResults → Console Output

3. File Organization
math-skills/
├── pipeline/      # Core functions of the pipeline
├── tests/         # Unit tests for each function
├── data/          # Sample input files
├── main.go        # Program entry point
├── go.mod
├── README.md
└── ARCHITECTURE.md


pipeline/: Each file contains a single function or closely related functions

tests/: Follows the naming convention <function>_test.go

data/: Contains example .txt files for testing and validation

4. Design Decisions

Pipeline Design

Each stage has one responsibility → easier to test, maintain, and extend.

Makes debugging simple because you can test each function in isolation.

Error Handling

ReadFile returns errors for missing or invalid files.

All other functions assume valid integer slices.

Testing

Unit tests cover all edge cases: empty files, single numbers, negative numbers, duplicates.

Ensures correctness before submission to 01‑edu grader.

Grader Compliance

Output formatting strictly matches grader requirements:

Average: <integer>
Median: <integer>
Variance: <integer>
Standard Deviation: <integer>


No extra print statements or debug logs.

5. Example Execution Flow

Input file (data1.txt):

10
20
30
40
50


Execution:

go run main.go data1.txt


Internal Processing:

ReadFile → [10, 20, 30, 40, 50]

CalculateAverage → 30

CalculateMedian → 30

CalculateVariance → 200

CalculateStdDev → 14

PrintResults → console output

Output:

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14

6. Extensibility

Adding new statistics (e.g., min, max) requires adding one new function in the pipeline.

Easy to integrate with other programs or graders due to modular design.

Unit tests can be extended without changing the main pipeline.