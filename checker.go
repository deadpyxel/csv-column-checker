package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// CheckEmptyColumn checks for empty columns in a CSV file.
//
// Parameters:
// - filePath: a string representing the path to the CSV file to be checked
// - delimiter: a string representing the delimiter used in the CSV file (default is comma)
//
// Returns:
// - []int: a slice of integers representing the indices of empty columns in the CSV file
// - error: an error if any issues occur during file reading or processing
//
// This function opens the specified CSV file, reads the header row to determine the number of columns,
// and then iterates through each row to check for empty values in each column. It returns a slice of integers
// representing the indices of columns that are empty. If the file is empty or no empty columns are found,
// it returns nil for the slice of integers.
//
// Possible errors that can be returned include:
// - "failed to open file: " followed by the specific error encountered when opening the file
// - "failed to read header row: " followed by the specific error encountered when reading the header row
// - "failed to read row: " followed by the specific error encountered when reading a row
//
// Example usage:
//
//	emptyCols, err := CheckEmptyColumn("data.csv", ",")
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Empty columns:", emptyCols)
//	}
func CheckEmptyColumn(filePath string, delimiter string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if delimiter == "" {
		delimiter = ","
	}
	reader.Comma = rune(delimiter[0])

	header, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, nil // Empty file, no empty columns
		}
		return nil, fmt.Errorf("failed to read header row: %w", err)
	}
	numCol := len(header)

	// Create a slide to track if a column has any non-empty value
	hasValue := make([]bool, numCol)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			return nil, fmt.Errorf("failed to read row: %w", err)
		}
		// Iterate through the columns in the row
		for i := 0; i < numCol; i++ {
			// Check if the column has a value (i.e, it is not empty)
			if i < len(row) && row[i] != "" {
				hasValue[i] = true
			}
		}
	}
	// Identify empty columns
	var emptyCols []int
	for i := 0; i < numCol; i++ {
		if !hasValue[i] {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols, nil
}
