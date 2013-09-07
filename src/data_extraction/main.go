// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Data extraction utilities.\n")
		fmt.Printf("Usage: %s command options.\n")
		fmt.Printf("Available commands:\n")
		fmt.Printf("\textract_terms\n")
		fmt.Printf("\tsplit_records\n")
		fmt.Printf("\textract_labels\n")
		return
	}
	switch os.Args[1] {
	case "extract_terms":
		if len(os.Args) != 6 {
			fmt.Printf("Extract term frequencies from the records.\n")
			fmt.Printf("Usage: %s extract_terms csv_input non_stemmer_file term_frequencies_file tf_idf_file\n", os.Args[0])
			return
		}

		csv_file := os.Args[2]
		label_file := os.Args[3]
		term_freq_file := os.Args[4]
		tf_idf_file := os.Args[5]

		TermExtractionMain(csv_file, label_file, term_freq_file, tf_idf_file)

		break
	case "split_records":
		if len(os.Args) != 5 {
			fmt.Printf("Split csv records into multiple files.\n")
			fmt.Printf("Usage: %s split_records csv_input prefix max_records.\n", os.Args[0])
			return
		}
		csv_file := os.Args[2]
		prefix := os.Args[3]
		max_records, err := strconv.ParseInt(os.Args[4], 10, 32)
		if err != nil {
			fmt.Printf("Invalid parameter max_records: %s.\n", os.Args[4])
			return
		}
		err = SplitCSVMain(csv_file, prefix, int(max_records))
		if err != nil {
			fmt.Printf("Error: %s.\n", err)
		}
		break
	case "extract_labels":
		if len(os.Args) != 4 {
			fmt.Printf("Extract labels and their corresponding frequencies.\n")
			fmt.Printf("Usage: %s extract_labels csv_input labels_output.\n")
			return
		}
		csv_input := os.Args[2]
		label_output := os.Args[3]
		err := ExtractLabelsMain(csv_input, label_output)
		if err != nil {
			fmt.Printf("Error: %s.\n", err)
		}
		break
	default:
		fmt.Printf("Unknown command: %s.\n", os.Args[1])
		break
	}
}
