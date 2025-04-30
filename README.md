# csvcut

`csvcut` is a command-line tool for filtering rows and columns from CSV files. It allows you to specify a range of rows and columns to extract from an input CSV file and save the result to an output file or print it to the standard output.

## Features

- Extract specific rows and columns from a CSV file.
- Output the filtered data to a file or standard output.
- Simple and intuitive command-line interface.

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```bash
   git clone <repository-url>
   cd csvcut
   ```
3. Build the project:
   ```bash
   go build -o csvcut main.go
   ```

## Usage

Run the `csvcut` command with the following options:

```bash
csvcut --input <input.csv> [--output <output.csv>] --start-column <start> --end-column <end> --start-row <start> --end-row <end>
```

### Options

- `--input` or `-i`: Path to the input CSV file (optional). If omitted, the tool reads from the standard input.
- `--output` or `-o`: Path to the output CSV file (optional). If omitted, the result is printed to the standard output.
- `--start-column` or `-sc`: Start column (1-based index).
- `--end-column` or `-ec`: End column (1-based index).
- `--start-row` or `-sr`: Start row (1-based index).
- `--end-row` or `-er`: End row (1-based index).

### Example

Extract columns 2 to 4 and rows 2 to 3 from `input.csv` and save the result to `output.csv`:

```bash
csvcut --input input.csv --output output.csv --start-column 2 --end-column 4 --start-row 2 --end-row 3
```

## Updates

### April 30, 2025

- Added support for reading input from standard input when the `-i` option is omitted.
- Refactored the code for better modularity and maintainability.

## Testing

Run the tests using the following command:

```bash
go test ./...
```

## License

This project is licensed under the MIT License.
