package main

import (
	"os"
	"testing"
)

func TestFilterRowsAndColumns(t *testing.T) {
	records := [][]string{
		{"A1", "B1", "C1", "D1"},
		{"A2", "B2", "C2", "D2"},
		{"A3", "B3", "C3", "D3"},
		{"A4", "B4", "C4", "D4"},
	}

	tests := []struct {
		startColumn, endColumn, startRow, endRow int
		expected                                 [][]string
	}{
		{1, 2, 1, 2, [][]string{{"A1", "B1"}, {"A2", "B2"}}},
		{2, 4, 2, 3, [][]string{{"B2", "C2", "D2"}, {"B3", "C3", "D3"}}},
		{1, 1, 1, 4, [][]string{{"A1"}, {"A2"}, {"A3"}, {"A4"}}},
		{3, 3, 3, 3, [][]string{{"C3"}}},
	}

	for _, test := range tests {
		result := filterRowsAndColumns(records, test.startColumn, test.endColumn, test.startRow, test.endRow)
		if !equal(result, test.expected) {
			t.Errorf("For range (%d-%d, %d-%d), expected %v but got %v", test.startColumn, test.endColumn, test.startRow, test.endRow, test.expected, result)
		}
	}
}

func TestMainFunction(t *testing.T) {
	input := "A1,B1,C1,D1\nA2,B2,C2,D2\nA3,B3,C3,D3\nA4,B4,C4,D4\n"
	expectedOutput := "B2,C2,D2\nB3,C3,D3\n"

	inputFile, err := os.CreateTemp("", "input.csv")
	if err != nil {
		t.Fatalf("Failed to create temp input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	_, err = inputFile.WriteString(input)
	if err != nil {
		t.Fatalf("Failed to write to temp input file: %v", err)
	}
	inputFile.Close()

	outputFile, err := os.CreateTemp("", "output.csv")
	if err != nil {
		t.Fatalf("Failed to create temp output file: %v", err)
	}
	defer os.Remove(outputFile.Name())
	outputFile.Close()

	os.Args = []string{"cmd", "--input", inputFile.Name(), "--output", outputFile.Name(), "--start-column", "2", "--end-column", "4", "--start-row", "2", "--end-row", "3"}

	main()

	output, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if string(output) != expectedOutput {
		t.Errorf("Expected output %q but got %q", expectedOutput, string(output))
	}
}

func equal(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
