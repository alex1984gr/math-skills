How It Works – math-skills

math-skills is a Go project that calculates basic statistics from a file containing integers. This document explains how the program works step by step, from reading the input file to printing the results.

1. Program Entry Point

The program starts in main.go, which expects a file path as a command-line argument:

go run main.go data1.txt


The program reads the file using the ReadFile() function.

If the file does not exist or contains invalid data, the program prints an error and exits.

2. Reading the File

Function: ReadFile(path string) ([]int, error)

Opens the file at the specified path.

Reads all lines sequentially.

Converts each line into an integer using strconv.Atoi.

Returns a slice of integers if successful, otherwise an error.

Example Input (data1.txt):

10
20
30
40
50


Example Output (slice):

[10 20 30 40 50]

3. Calculating Statistics

Once the data is read, the program calculates the following statistics in order:

3.1 Average

Function: CalculateAverage(numbers []int) int

Sums all integers.

Divides by the number of integers.

Returns the average, rounded to the nearest integer.

3.2 Median

Function: CalculateMedian(numbers []int) int

Sorts the numbers in ascending order.

For odd-length slices, returns the middle value.

For even-length slices, returns the average of the two middle values, rounded.

3.3 Variance

Function: CalculateVariance(numbers []int, mean int) int

Computes the squared difference of each number from the mean.

Calculates the average of these squared differences.

Returns the variance, rounded to the nearest integer.

3.4 Standard Deviation

Function: CalculateStdDev(variance int) int

Computes the square root of the variance.

Returns the result, rounded to the nearest integer.

4. Printing Results

Function: PrintResults(avg, median, variance, stddev int)

Prints all statistics in the format expected by the 01‑edu grader:

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14


The program does not print any extra information or debug logs to ensure grader compliance.

5. Example Execution Flow

Command:

go run main.go data1.txt


Internal Steps:

ReadFile("data1.txt") → [10, 20, 30, 40, 50]

CalculateAverage() → 30

CalculateMedian() → 30

CalculateVariance() → 200

CalculateStdDev() → 14

PrintResults() → prints results to console

Console Output:

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14

6. Testing

All functions have unit tests in the tests/ folder.

Tests cover:

Valid and invalid files

Empty lists, single numbers, negative numbers

Correct calculation for each statistic

Proper formatting of printed output

Run tests:

go test ./tests/...

7. Notes

The program uses a pipeline design: each step is independent and can be tested individually.

The code is grader-ready: output matches exactly what the 01‑edu grader expects.

Adding new statistics in the future is easy: just add a new function and call it in main.go.