// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

//Expected Input of the form:
//	"Id","Title","Body","Tags"

func main1() {
	if len(os.Args) != 3 {
		fmt.Printf("Extract labels from the records.\n")
		fmt.Printf("Usage: %s csv_file output_tab_file\n", os.Args[0])
		return
	}

	in_fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to open file %s: %s.\n", os.Args[1], err)
		return
	}
	defer in_fp.Close()

	out_fp, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Print("Failed to create output file %s: %s.\n", os.Args[2], err)
		return
	}
	defer out_fp.Close()

	csv_reader := csv.NewReader(in_fp)
	tab_writer := bufio.NewWriter(out_fp)

	label_freq := make(map[string]int)
	for {
		fields, err := csv_reader.Read()
		if err == nil {
			labels := strings.Split(fields[3], " ")
			for _, label := range labels {
				label_freq[label]++
			}
		} else if err == io.EOF {
			break
		} else {
			fmt.Errorf("Failed to process record: %s.\n", err)
		}
	}

	for name, freq := range label_freq {
		tab_writer.WriteString(fmt.Sprintf("%s\t%d\n", name, freq))
	}

	return
}
