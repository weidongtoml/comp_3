// Copyright Weidong Liang 2013

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

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Converts CSV fields to tab separted records.\n")
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
	for {
		fields, err := csv_reader.Read()
		if err == nil {
			for index, value := range fields {
				if index == 0 {
					tab_writer.WriteString(value)
				} else if index == 1 || index == 2 {
					proper_value := strings.Map(func(r rune) rune {
						var v rune
						switch r {
						case rune('\n'), rune('\t'), rune('\f'):
							v = rune(' ')
							break
						default:
							v = r
							break
						}
						return v
					}, value)
					tab_writer.WriteString("\t" + proper_value)
				} else {
					tab_writer.WriteString("\t" + value)
				}
			}
			tab_writer.WriteByte('\n')
		} else if err == io.EOF {
			break
		} else {
			fmt.Errorf("Failed to process record: %s.\n", err)
		}
	}

}
