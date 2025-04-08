package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Helper function to create a large CSV file for benchmarking
func createLargeCSVFile(numRows, numCols int) (string, error) {
	tmpfile, err := os.CreateTemp("", "large_test.csv")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	header := make([]string, numCols)
	for i := 0; i < numCols; i++ {
		header[i] = fmt.Sprintf("col%d", i)
	}
	headerRow := strings.Join(header, ",") + "\n"
	_, err = tmpfile.WriteString(headerRow)
	if err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		return "", err
	}

	for i := 0; i < numRows; i++ {
		row := make([]string, numCols)
		for j := 0; j < numCols; j++ {
			if i%100 == 0 && j == 1 { // Introduce some empty columns
				row[j] = ""
			} else {
				row[j] = fmt.Sprintf("value%d", j)
			}
		}
		rowString := strings.Join(row, ",") + "\n"
		_, err = tmpfile.WriteString(rowString)
		if err != nil {
			tmpfile.Close()
			os.Remove(tmpfile.Name())
			return "", err
		}
	}

	err = tmpfile.Close()
	if err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}

	return tmpfile.Name(), nil
}

func BenchmarkCheckEmptyColumns(b *testing.B) {
	// Define different benchmark cases
	benchmarkCases := []struct {
		name    string
		numRows int
		numCols int
	}{
		{name: "100 rows, 10 cols", numRows: 100, numCols: 10},
		{name: "1000 rows, 10 cols", numRows: 1000, numCols: 10},
		{name: "100 rows, 100 cols", numRows: 100, numCols: 100},
		{name: "1000 rows, 100 cols", numRows: 1000, numCols: 100},
		{name: "5000 rows, 50 cols", numRows: 5000, numCols: 50}, // Added new test case
	}

	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			// Create a large CSV file for benchmarking
			filePath, err := createLargeCSVFile(bc.numRows, bc.numCols)
			if err != nil {
				b.Fatalf("Failed to create large CSV file: %v", err)
			}
			defer os.Remove(filePath)

			// Run the benchmark
			b.ResetTimer() // Reset the timer to exclude file creation time
			for i := 0; i < b.N; i++ {
				_, _, err := CheckEmptyColumn(filePath, ",")
				if err != nil {
					b.Fatalf("Error during CheckEmptyColumns: %v", err)
				}
			}
		})
	}
}
