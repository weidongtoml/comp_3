package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// ProcessCSVFieldsWith process each records in the given csv_file
// using the processor function.
func ProcessCSVFieldsWith(csv_file string, processor func([]string) error) error {
	in_fp, err := os.Open(csv_file)
	if err != nil {
		return fmt.Errorf("Failed to open file %s: %s.\n", csv_file, err)
	}
	defer in_fp.Close()

	csv_reader := csv.NewReader(in_fp)
	for {
		fields, err := csv_reader.Read()
		if err == nil {
			err = processor(fields)
			if err != nil {
				return err
			}
		} else if err == io.EOF {
			break
		} else {
			log.Printf("Failed to process record: %s.\n", err)
		}
	}
	return nil
}
