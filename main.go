package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <inputFile> <outputFile>")
		os.Exit(1)
	}

	inputPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get absolute path for input file: %v\n", err)
		os.Exit(1)
	}

	outputPath, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get absolute path for output file: %v\n", err)
		os.Exit(1)
	}

	if err := convertFile(inputPath, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error during conversion: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File converted successfully")
}

func convertFile(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	transformedReader := transform.NewReader(bufio.NewReader(inputFile), japanese.ShiftJIS.NewDecoder())
	csvReader := csv.NewReader(transformedReader)

	bufferedWriter := bufio.NewWriter(outputFile)
	csvWriter := csv.NewWriter(bufferedWriter)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read from input file: %v", err)
		}

		if err := csvWriter.Write(record); err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("failed to flush csv writer: %v", err)
	}

	if err := bufferedWriter.Flush(); err != nil {
		return fmt.Errorf("failed to flush output file: %v", err)
	}

	return nil
}
