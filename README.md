# graphic_formate

# Transaction Formatter

This Go program formats transactions based on different time intervals such as hour, day, week, and month.

## Overview

The program defines a `Transaction` struct to represent individual transactions with a value and a timestamp. It provides a `FormatTransactions` function that takes a slice of transactions and an interval string as input, and returns a slice of transactions formatted based on the specified interval.

## Usage

To use the program, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the directory containing the `main.go` file.
3. Run the program by executing `go run main.go`.
4. Modify the `main` function in `main.go` to customize the transactions and interval as needed.
5. View the formatted transactions printed to the console.

# Transaction Formatter Tests

This document contains test cases for the `FormatTransactions` function in the Transaction Formatter program.

## Test Cases

### TestHourly

**Description:** 

Test checks if the function outputs data correctly if the user enters: month, week, day, hour.

## Test Execution

To execute the tests:

1. Clone the repository to your local machine.
2. Navigate to the directory containing the test file.
3. Run the tests using the `go test -v .` command.
4. View the test results.
