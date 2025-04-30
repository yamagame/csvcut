package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	inputFile, outputFile, startColumn, endColumn, startRow, endRow := parseFlags()

	records, err := readCSV(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	filteredRecords := filterRowsAndColumns(records, *startColumn, *endColumn, *startRow, *endRow)

	if *outputFile == "" {
		writeCSVToStdout(filteredRecords)
	} else {
		if err := writeCSVToFile(filteredRecords, *outputFile); err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
		}
	}
}

func parseFlags() (*string, *string, *int, *int, *int, *int) {
	help := flag.Bool("help", false, "Show usage information")
	flag.BoolVar(help, "h", false, "Show usage information (shorthand)")

	inputFile := flag.String("input", "", "Path to the input CSV file")
	outputFile := flag.String("output", "", "Path to the output CSV file")
	startColumn := flag.Int("start-column", 1, "Start column (1-based index)")
	endColumn := flag.Int("end-column", 1, "End column (1-based index)")
	startRow := flag.Int("start-row", 1, "Start row (1-based index)")
	endRow := flag.Int("end-row", 1, "End row (1-based index)")

	// Add shorthand options
	flag.StringVar(inputFile, "i", "", "Path to the input CSV file (shorthand)")
	flag.StringVar(outputFile, "o", "", "Path to the output CSV file (shorthand)")
	flag.IntVar(startColumn, "sc", 1, "Start column (1-based index) (shorthand)")
	flag.IntVar(endColumn, "ec", 1, "End column (1-based index) (shorthand)")
	flag.IntVar(startRow, "sr", 1, "Start row (1-based index) (shorthand)")
	flag.IntVar(endRow, "er", 1, "End row (1-based index) (shorthand)")

	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	args := flag.Args()
	if *outputFile == "" && len(args) > 0 {
		*outputFile = args[0]
	}

	return inputFile, outputFile, startColumn, endColumn, startRow, endRow
}

func printUsage() {
	fmt.Println("Usage: csvcut --input <input.csv> [--output <output.csv>] --start-column <start> --end-column <end> --start-row <start> --end-row <end>")
}

func readCSV(filePath string) ([][]string, error) {
	var file *os.File
	var err error

	if filePath == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func writeCSVToStdout(records [][]string) {
	writer := csv.NewWriter(os.Stdout)
	if err := writer.WriteAll(records); err != nil {
		fmt.Printf("Error writing to standard output: %v\n", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Printf("Error flushing to standard output: %v\n", err)
	}
}

func writeCSVToFile(records [][]string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return err
	}
	writer.Flush()
	return writer.Error()
}

func filterRowsAndColumns(records [][]string, startColumn, endColumn, startRow, endRow int) [][]string {
	if startRow < 1 || endRow > len(records) || startRow > endRow {
		fmt.Println("Invalid row range.")
		return nil
	}

	var filtered [][]string
	for i, record := range records {
		if i+1 < startRow || i+1 > endRow {
			continue
		}
		if startColumn < 1 || endColumn > len(record) || startColumn > endColumn {
			fmt.Println("Invalid column range.")
			return nil
		}
		filtered = append(filtered, record[startColumn-1:endColumn])
	}
	return filtered
}
