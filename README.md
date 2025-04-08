# csv-column-checker

A command-line tool to check for completely empty columns in CSV files.

## Overview

`csv-column-checker` is a simple, cross-platform utility that analyzes CSV files and identifies columns where all data entries are empty (i.e., contain no non-whitespace characters). It adheres to the Unix philosophy of doing one thing well, providing a straightforward way to validate CSV data integrity.

## Features

*   **Identifies Empty Columns:** Accurately detects columns with no data beyond the header row.
*   **Custom Delimiter:** Supports specifying a custom column delimiter (defaults to comma).
*   **Column Name Output:**  Optionally outputs the names of the empty columns instead of their indices.
*   **Clear Error Reporting:** Provides informative error messages for common issues like missing files or invalid delimiters.
*   **Unix Philosophy:** Designed to be used in command-line pipelines, with appropriate exit codes for success and failure.
*   **Cross-Platform:**  Works on Linux, macOS, and Windows.
*   **Well-Tested:** Includes comprehensive unit and integration tests to ensure reliability.

## Installation

1.  **Prerequisites:** Go 1.20 or later must be installed.

2.  **Download and Install:**

    ```bash
    go install github.com/deadpyxel/csv-column-checker@latest
    ```

    Make sure your `$GOPATH/bin` or `$HOME/go/bin` directory is in your system's `$PATH` to access the `csv-column-checker` executable.

## Usage

```bash
csv-column-checker -file <csv_file> [options]
```

### Options

*   `-file <csv_file>`: (Required) The path to the CSV file to check.
*   `-delimiter <delimiter>`: (Optional) The column delimiter. Defaults to `,` (comma).
*   `-names`: (Optional) Output the column names instead of indices.
*   `-version`: (Optional) Show the version number and exit.

### Examples

1.  **Check a CSV file with the default delimiter:**

    ```bash
    csv-column-checker -file data.csv
    ```

2.  **Check a CSV file with a semicolon delimiter:**

    ```bash
    csv-column-checker -file data.csv -delimiter ";"
    ```

3.  **Output the names of the empty columns:**

    ```bash
    csv-column-checker -file data.csv -names
    ```

### Exit Codes

*   `0`: Success.  No completely empty columns were found.
*   `1`: Error or empty columns found.  The tool will print an error message to stderr, or the list of empty columns to stdout.

### Usage in Pipelines

You can use `csv-column-checker` in shell pipelines to automate CSV validation. For example:

```bash
csv-column-checker -file data.csv | grep "Empty columns" && echo "Validation Failed" || echo "Validation Passed"
```


## Contributing

Contributions are welcome!

Please read our [Code of Conduct](CODE_OF_CONDUCT.md) and [Contributing Guidelines](CONTRIBUTING.md) for instructions in how to proceed.

## Development

### Building from Source

```bash
go build -o csv-column-checker
```

### Running Tests

```bash
go test
```

### Running Benchmarks

```bash
go test -bench=.
```
