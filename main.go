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
		outputNames bool
	)

	flag.StringVar(&filePath, "file", "", "Path to the CSV file (required)")
	flag.StringVar(&delimiter, "delimiter", ",", "Column delimiter (default: comma)")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit")
	flag.BoolVar(&outputNames, "names", false, "Output column names instead of the indexes")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -file <csv_file> [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if showVersion {
		fmt.Println("csv-column-checker v0.2.0")
		os.Exit(0)
	}

	if filePath == "" {
		fmt.Fprintln(os.Stderr, "Error: -file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Call the checker function
	emptyCols, emptyNames, err := CheckEmptyColumn(filePath, delimiter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(emptyCols) > 0 {
		if outputNames {
			fmt.Println("Empty columns:", emptyNames)
		} else {
			fmt.Println("Empty columns found on position:", emptyCols)
		}
		os.Exit(1) // Indicates an error if empty columns are found
	}

	fmt.Println("No completely empty columns found.")
	os.Exit(0) // Indicate success
}
