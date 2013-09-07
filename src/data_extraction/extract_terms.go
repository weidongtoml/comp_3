// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"github.com/weidongtoml/go_common/util"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

//Expected Input of the form:
//	"Id","Title","Body","Tags"

type tf_idf_T struct {
	tf  int
	idf int
}

func LoadLabels(label_file string) (map[string]bool, error) {
	label_map := make(map[string]bool)
	err := util.ForEachLineInFile(label_file, func(line string) (bool, error) {
		fields := strings.Fields(line)
		label_map[fields[0]] = true
		return true, nil
	})
	return label_map, err
}

func TermExtractionMain(csv_file, label_file, term_freq_file, tf_idf_file string) {
	label_map, err := LoadLabels(label_file)
	if err != nil {
		log.Printf("Failed to load label: %s", err)
		return
	}

	term_freq_fd, err := os.Create(term_freq_file)
	if err != nil {
		log.Printf("Failed to create Term Frequeny File: %s, %s.", term_freq_file, err)
		return
	}
	defer term_freq_fd.Close()
	term_freq_writer := bufio.NewWriter(term_freq_fd)

	tf_idf_map := make(map[string]tf_idf_T)
	ProcessCSVFieldsWith(csv_file, func(fields []string) error {
		text := fields[1] + fields[2]

		text_only := strings.ToLower(StripCodeSectionFromHTML(text))
		//fmt.Printf("Process: %s\n", text_only)
		words := strings.FieldsFunc(text_only, func(r rune) bool {
			return (unicode.IsSpace(r) ||
				(unicode.IsPunct(r) && r != '#' && r != '.' && r != '-' && r != '+'))
		})
		sort.Strings(words)
		prev := ""
		cnt := 0
		output_line := ""
		for _, w := range words {
			if w == "" {
				continue
			}
			//trim trailing .
			last_index := len(w) - 1
			for ; last_index >= 0 && w[last_index] == '.'; last_index-- {
			}
			if last_index < 0 {
				continue
			}
			w = w[0 : last_index+1]

			if !label_map[w] {
				w = string(Stem([]byte(w)))
			}
			if w == prev {
				cnt++
			} else {
				if prev != "" {
					cur_tf_idf := tf_idf_map[prev]
					tf_idf_map[prev] = tf_idf_T{cur_tf_idf.tf + cnt, cur_tf_idf.idf + 1}

					output_line += fmt.Sprintf("%s:%d ", prev, cnt)
				}
				cnt = 1
				prev = w
			}
		}
		if prev != "" {
			cur_tf_idf := tf_idf_map[prev]
			tf_idf_map[prev] = tf_idf_T{cur_tf_idf.tf + cnt, cur_tf_idf.idf + 1}

			output_line += fmt.Sprintf("%s:%d", prev, cnt)
		}
		output_line = fields[3] + " | " + output_line
		//fmt.Printf("Out: %s\n\n", output_line)
		term_freq_writer.WriteString(output_line + "\n")
		return nil
	})

	tf_idf_fd, err := os.Create(tf_idf_file)
	if err != nil {
		log.Printf("Failed to create TF-IDF file: %s, %s.", tf_idf_file, err)
		return
	}
	defer tf_idf_fd.Close()
	tf_idf_writer := bufio.NewWriter(tf_idf_fd)

	for k, v := range tf_idf_map {
		tf_idf_writer.WriteString(fmt.Sprintf("%s\t%f\t%d\t%d\n",
			k, float64(v.tf)/float64(v.idf), v.tf, v.idf))
	}

	return
}
