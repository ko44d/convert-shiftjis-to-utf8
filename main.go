package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <inputFile> <outputFile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputFilePath, err := filepath.Abs(inputFile)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get absolute path for input file: %v", err))
		return
	}

	outputFilePath, err := filepath.Abs(outputFile)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get absolute path for output file: %v", err))
		return
	}

	i, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open input file: %v", err))
		return
	}
	defer i.Close()

	o, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create output file: %v", err))
		return
	}
	defer o.Close()

	r := transform.NewReader(bufio.NewReader(i), japanese.ShiftJIS.NewDecoder())

	w := bufio.NewWriter(o)
	defer w.Flush()

	cr := csv.NewReader(r)
	cw := csv.NewWriter(w)

	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(fmt.Errorf("failed to read from input file: %v", err))
			return
		}

		if err := cw.Write(record); err != nil {
			fmt.Println(fmt.Errorf("failed to write to output file: %v", err))
			return
		}
	}

	cw.Flush()
	if err := cw.Error(); err != nil {
		fmt.Println(fmt.Errorf("failed to flush output file: %v", err))
		return
	}

	fmt.Println("File converted successfully")
}
