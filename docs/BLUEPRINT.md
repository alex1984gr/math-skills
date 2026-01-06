Blueprint – math-skills

This document describes the pipeline blueprint of the math-skills project, showing how data flows through the program and how each function interacts.

1. Pipeline Overview

The program is designed as a linear pipeline:

Input File → ReadFile → CalculateAverage → CalculateMedian → CalculateVariance → CalculateStdDev → PrintResults → Console


Each step is modular and testable independently.

The output of one function is the input to the next function.

This design allows TDD (Test-Driven Development) and easy debugging.

2. Pipeline Steps
Step	Function	Input	Output	Description
1	ReadFile(path string)	File path	[]int	Reads all integers from the file. Returns error if the file cannot be read or data is invalid.
2	CalculateAverage(numbers []int)	[]int	int	Calculates the average of the numbers. Returns 0 for empty input.
3	CalculateMedian(numbers []int)	[]int	int	Calculates the median. Works for both odd and even-length slices.
4	CalculateVariance(numbers []int, mean int)	[]int, mean	int	Computes the variance from the mean.
5	CalculateStdDev(variance int)	int	int	Computes the standard deviation from the variance.
6	PrintResults(avg, median, variance, stddev int)	int, int, int, int	Console output	Prints all results in grader-compliant format.
3. Data Flow Diagram
┌────────────┐
│  data.txt  │
└─────┬──────┘
      │
      ▼
┌──────────────┐
│  ReadFile()  │
│  ([]int)     │
└─────┬────────┘
      │
      ▼
┌──────────────┐
│ CalculateAvg │
└─────┬────────┘
      │
      ▼
┌──────────────┐
│ CalculateMed │
└─────┬────────┘
      │
      ▼
┌──────────────┐
│ CalculateVar │
└─────┬────────┘
      │
      ▼
┌──────────────┐
│ CalculateSD  │
└─────┬────────┘
      │
      ▼
┌──────────────┐
│ PrintResults │
└──────────────┘
      │
      ▼
  Console Output

4. Key Design Principles

Single Responsibility

Each function does exactly one task.

Easy to test and debug.

Modularity

Functions can be reused or replaced without affecting others.

Supports future extensions like min, max, range, etc.

TDD-Friendly

Every step has its own unit test.

Ensures correctness at each stage.

Grader Compliance

Output is strictly formatted to match 01‑edu requirements.

No extra logs or debug print statements.

5. Example Execution

Input File (data1.txt):

10
20
30
40
50


Pipeline Execution:

ReadFile("data1.txt") → [10, 20, 30, 40, 50]

CalculateAverage() → 30

CalculateMedian() → 30

CalculateVariance() → 200

CalculateStdDev() → 14

PrintResults() → Console

Output:

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
