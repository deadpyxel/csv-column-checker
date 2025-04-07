package main

import (
	"fmt"
	"os"
)

func CheckEmptyColumn(filePath string, delimiter string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	return nil, nil
}
