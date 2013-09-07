// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Extract all the labels and the corresponding frequencies from
// the given csv_input file of the format:
//	"Id","Title","Body","Tags"
func ExtractLabelsMain(csv_input, output string) error {
	out_fp, err := os.Create(os.Args[2])
	if err != nil {
		return fmt.Errorf("Failed to create output file %s: %s.\n", os.Args[2], err)
	}
	defer out_fp.Close()

	label_freq := make(map[string]int)
	err = ProcessCSVFieldsWith(csv_input, func(fields []string) error {
		if len(fields) != 4 {
			log.Printf("Expected number of fields to be 4 but got: %v.", fields)
		} else {
			labels := strings.Split(fields[3], " ")
			for _, label := range labels {
				label_freq[label]++
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	tab_writer := bufio.NewWriter(out_fp)
	for name, freq := range label_freq {
		tab_writer.WriteString(fmt.Sprintf("%s\t%d\n", name, freq))
	}

	return nil
}
