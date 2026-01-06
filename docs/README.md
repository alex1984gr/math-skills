math-skills

math-skills is a Go project that reads a list of integers from a file and calculates basic statistics:

Average

Median

Variance

Standard Deviation

The project is structured as a pipeline, with each step implemented in a separate function. It also includes unit tests following TDD principles.

Installation

Clone the repository:

git clone <your-repo-url>
cd math-skills


Install dependencies:

go mod tidy


Run the program with a data file:

go run main.go data1.txt

Project Structure
math-skills/
├── pipeline/          # All core functions
│   ├── read_file.go
│   ├── average.go
│   ├── median.go
│   ├── variance.go
│   ├── stddev.go
│   └── print.go
├── tests/             # Unit tests for each function
│   ├── read_file_test.go
│   ├── calculate_average_test.go
│   ├── calculate_median_test.go
│   ├── calculate_variance_test.go
│   ├── calculate_stddev_test.go
│   └── print_results_test.go
├── data/              # Example input files
│   ├── data1.txt
│   ├── data2.txt
│   ├── data3.txt
│   ├── data4.txt
│   └── data5.txt
├── go.mod
└── main.go

Functions
ReadFile(path string) ([]int, error)

Reads integers from a file. Each line must contain a single integer.
Returns a slice of integers and an error if the file cannot be read.

Example:

numbers, err := pipeline.ReadFile("data1.txt")

CalculateAverage(numbers []int) int

Calculates the average of a slice of integers.
Returns 0 if the slice is empty.

avg := pipeline.CalculateAverage(numbers)

CalculateMedian(numbers []int) int

Calculates the median of a slice of integers.
Works for both odd and even-length slices.

median := pipeline.CalculateMedian(numbers)

CalculateVariance(numbers []int, mean int) int

Calculates the variance of a slice given the mean.

variance := pipeline.CalculateVariance(numbers, avg)

CalculateStdDev(variance int) int

Calculates the standard deviation from the variance.

stddev := pipeline.CalculateStdDev(variance)

PrintResults(avg, median, variance, stddev int)

Prints all statistics in the expected format for the grader.

pipeline.PrintResults(avg, median, variance, stddev)


Example Output:

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14

Testing

All functions have corresponding unit tests in the tests/ folder.

Run all tests with:

go test ./tests/...


Tests cover:

Reading valid and non-existent files

Calculating average, median, variance, and stddev for various datasets

Correct output formatting

Example Data

data1.txt

10
20
30
40
50


Expected Output

Average: 30
Median: 30
Variance: 200
Standard Deviation: 14


Other examples:

File	Average	Median	Variance	StdDev
data2.txt	5	5	8	3
data3.txt	5	5	8	3
data4.txt	42	42	0	0
data5.txt	0	0	67	8

Use these files to test your program before submitting to ensure correct results.