// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	//	"github.com/reiver/go-porterstemmer"
	"io"
	"log"
	"os"
	"strings"
)

//Expected Input of the form:
//	"Id","Title","Body","Tags"

func ProcessCSVFieldsWith(csv_file string, processor func([]string)) {
	in_fp, err := os.Open(csv_file)
	if err != nil {
		log.Printf("Failed to open file %s: %s.\n", csv_file, err)
		return
	}
	defer in_fp.Close()

	csv_reader := csv.NewReader(in_fp)
	for {
		fields, err := csv_reader.Read()
		if err == nil {
			processor(fields)
		} else if err == io.EOF {
			break
		} else {
			log.Printf("Failed to process record: %s.\n", err)
		}
	}

}

type tf_idf_T struct {
	tf  int
	idf int
}

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Extract term frequencies from the records.\n")
		fmt.Printf("Usage: %s csv_file non_stemmer_file term_frequencies_file tf_idf_file\n", os.Args[0])
		return
	}

	csv_file := os.Args[1]
	non_stemmer_file := os.Args[2]
	term_freq_file := os.Args[3]
	tf_idf_file := os.Args[4]

	tf_idf_map := make(map[string]tf_idf_T)

	do_title := true
	do_body := true

	ProcessCSVFieldsWith(csv_file, func(fields []string) {
		for i := 1; i <= 2; i++ {
			if i == 1 && !do_title {
				continue
			}
			if i == 2 && !do_body {
				continue
			}
			text := fields[i]
			//Filter text out of <code></code> <pre></pre> sections and remove html
			//tags, such as <p></p>, <strong></strong> etc.
			filtered_test := ""
			const (
				kTagStart = iota
				kTagBody
				kTagEnd
				kBody
			)
			state := kBody
			cur_text := ""
			for _, c := range text {
				switch c {
				case '<':
					state := kTagBody
					break
				case '>':
					state := kTagEnd
					break
				default:
					if state == kTagBody {
						cur_text += c
					} else if state == kTagEnd {

					}

					break
				}
			}
		}
	})

	out_fp, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Print("Failed to create output file %s: %s.\n", os.Args[2], err)
		return
	}
	defer out_fp.Close()

	tab_writer := bufio.NewWriter(out_fp)

	return
}
