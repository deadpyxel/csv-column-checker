package main

import (
	"os"
	"slices"
	"testing"
)

func createTempFile(content string) (string, error) {
	tmpfile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	_, err = tmpfile.WriteString(content)
	if err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}

func TestCheckEmptyColumnErrors(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		delimiter string
	}{
		{name: "when path is invalid returns error", filePath: "notreal.csv", delimiter: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, _, err := CheckEmptyColumn(tc.filePath, tc.delimiter)
			if res != nil {
				t.Errorf("Expected result to be nil, got %v instead", res)
			}
			if err == nil {
				t.Errorf("expected error, got nil instead with args (filkePath=%s, delimiter=%s)", tc.filePath, tc.delimiter)
			}
		})
	}
}

func TestCheckEmptyColumn(t *testing.T) {
	tests := []struct {
		name          string   // testcase description
		csvContent    string   // temp csv contents
		delimiter     string   //which delimiter to use
		expectedIdx   []int    // resulting indeces
		expectedNames []string // resulting column names
		expectedErr   bool     // do we expect an error
	}{
		{
			name:          "when no empty columns returns empty slice and no errors",
			csvContent:    "h1,h2,h3\nv1,v2,v3",
			delimiter:     ",",
			expectedIdx:   []int{},
			expectedNames: []string{},
			expectedErr:   false,
		},
		{
			name:          "when one empty column returns single element slice and no errors",
			csvContent:    "h1,h2,h3\nv1,,v3",
			delimiter:     ",",
			expectedIdx:   []int{1},
			expectedNames: []string{"h2"},
			expectedErr:   false,
		},
		{
			name:          "when multiple empty columns returns multi-element slice and no errors",
			csvContent:    "h1,h2,h3\nv1,,",
			delimiter:     ",",
			expectedIdx:   []int{1, 2},
			expectedNames: []string{"h2", "h3"},
			expectedErr:   false,
		},
		{
			name:          "when all columns are empty returns all indices",
			csvContent:    "h1,h2,h3\n,,",
			delimiter:     ",",
			expectedIdx:   []int{0, 1, 2},
			expectedNames: []string{"h1", "h2", "h3"},
			expectedErr:   false,
		},
		{
			name:          "when delimiter is non-standard function handles without errors",
			csvContent:    "h1;h2;h3\nv1;;v3",
			delimiter:     ";",
			expectedIdx:   []int{1},
			expectedNames: []string{"h2"},
			expectedErr:   false,
		},
		{
			name:          "when delimiter is empty string takes default value",
			csvContent:    "h1,h2,h3\nv1,,v3",
			delimiter:     "",
			expectedIdx:   []int{1},
			expectedNames: []string{"h2"},
			expectedErr:   false,
		},
		{
			name:          "when empty file returns empty slice",
			csvContent:    "",
			delimiter:     ",",
			expectedIdx:   []int{},
			expectedNames: []string{},
			expectedErr:   false,
		},
		{
			name:          "when only header is present returns empty slice",
			csvContent:    "h1,h2,h3",
			delimiter:     ",",
			expectedIdx:   []int{},
			expectedNames: []string{},
			expectedErr:   false,
		},
		{
			name:          "when rows have more columns than header returns error",
			csvContent:    "h1,h2\nv1,v2,v3", // Content doesn't matter in this case
			delimiter:     ",",
			expectedIdx:   []int{},
			expectedNames: []string{},
			expectedErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := createTempFile(tc.csvContent)
			if err != nil {
				t.Fatalf("Error creating temp file for tests: %v", err)
			}
			defer os.Remove(f)

			resIdx, resNames, err := CheckEmptyColumn(f, tc.delimiter)
			if tc.expectedErr {
				if err == nil {
					t.Errorf("expected error, got nil instead of args (filePath=%s, delimiter=%s)", f, tc.delimiter)
				}
				return
			}
			if err != nil {
				t.Errorf("Error checking file: %v", err)
			}
			if !slices.Equal(resIdx, tc.expectedIdx) {
				t.Errorf("Expected resulting indices to be %v, got %v instead", tc.expectedIdx, resIdx)
			}
			if !slices.Equal(resNames, tc.expectedNames) {
				t.Errorf("Expected resulting names to be %v, got %v instead", tc.expectedNames, resNames)
			}
		})
	}
}
