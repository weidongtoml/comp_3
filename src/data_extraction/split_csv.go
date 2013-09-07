// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"
)

//Expected Input of the form:
//	"Id","Title","Body","Tags"

// Split the given CSV file into multiple smaller files, with each file
// having at most max_records. The splitted files will be named with the given
// prefix followed by a number starting from 0.
// Also, all space characters in the records will be converted to single space.
func SplitCSVMain(src_file, prefix string, max_records int) error {
	record_counter := 0
	file_counter := 0
	var record_fd *os.File
	var record_writer *csv.Writer
	defer func() {
		if record_fd != nil {
			record_fd.Close()
		}
	}()
	return ProcessCSVFieldsWith(src_file, func(fields []string) error {
		if record_writer == nil {
			var err error
			fname := fmt.Sprintf("%s_%d.csv", prefix, file_counter)
			record_fd, err = os.Create(fname)
			if err != nil {
				return fmt.Errorf("Failed to create file %s: %s.", fname, err)
			}
			record_writer = csv.NewWriter(record_fd)
		}
		for index, value := range fields {
			fields[index] = strings.Map(func(r rune) rune {
				if unicode.IsSpace(r) {
					return ' '
				}
				return r
			}, value)
		}
		record_writer.Write(fields)
		record_counter++
		if record_counter > max_records {
			record_fd.Close()
			record_writer = nil
			record_fd = nil
			record_counter = 0
			file_counter++
		}
		return nil
	})
}
