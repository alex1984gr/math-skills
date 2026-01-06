Glossary – math-skills

This glossary explains the main terms and concepts used in the math-skills project. It is intended for developers, graders, or anyone reviewing the code to understand key concepts quickly.

A

Average (Mean)

The arithmetic mean of a set of numbers.

Calculated by summing all numbers and dividing by the total count.

Function: CalculateAverage(numbers []int) int

D

Data File

A plain text file (.txt) containing integers, one per line.

Example: data1.txt

Input to the program.

Data Slice

A Go slice ([]int) containing numbers read from the file.

Used as input to all calculation functions.

M

Median

The middle value of a sorted list of numbers.

If the list has an even number of elements, it is the average of the two middle numbers.

Function: CalculateMedian(numbers []int) int

P

Pipeline

A linear sequence of functions, where each function performs one task and passes its output to the next.

Makes code modular, testable, and easy to maintain.

PrintResults()

Function that outputs all computed statistics to the console.

Ensures grading-friendly output format.

S

Standard Deviation

A measure of how spread out numbers are from the mean.

Calculated as the square root of variance.

Function: CalculateStdDev(variance int) int

V

Variance

Measures the average squared deviation of each number from the mean.

Function: CalculateVariance(numbers []int, mean int) int

T

TDD (Test-Driven Development)

Development approach where tests are written before the actual code.

Ensures that every function has a corresponding unit test to verify correctness.

E

Error Handling

Mechanism to deal with invalid or missing files.

ReadFile() returns an error if the file cannot be read or contains invalid data.