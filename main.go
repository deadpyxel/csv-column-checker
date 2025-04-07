package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		filePath    string
		delimiter   string
		showVersion bool
	)

	flag.StringVar(&filePath, "file", "", "Path to the CSV file (required)")
	flag.StringVar(&delimiter, "delimiter", ",", "Column delimiter (default: comma)")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -file <csv_file> [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if showVersion {
		fmt.Println("csv-column-checker v0.1.0")
		os.Exit(0)
	}

	if filePath == "" {
		fmt.Fprintln(os.Stderr, "Error: -file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Call the checker function
	emptyCols, err := CheckEmptyColumn(filePath, delimiter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(emptyCols) > 0 {
		fmt.Println("Empty columns found:", emptyCols)
		os.Exit(1) // Indicates an error if empty columns are found
	}

	fmt.Println("No completely empty columns found.")
	os.Exit(0) // Indicate success
}
