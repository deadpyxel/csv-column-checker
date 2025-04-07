package main

import (
	"testing"
)

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
			res, err := CheckEmptyColumn(tc.filePath, tc.delimiter)
			if res != nil {
				t.Errorf("Expected result to be nil, got %v instead", res)
			}
			if err == nil {
				t.Errorf("expected error, got nil instead with args (filkePath=%s, delimiter=%s)", tc.filePath, tc.delimiter)
			}
		})
	}
}
