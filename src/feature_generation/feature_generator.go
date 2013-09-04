// Copyright 2013 Weidong Liang. All rights reserved.

package feature_generation

import (
	"bufio"
	"encoding/csv"
	feagen "feature_generator"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type generatorT struct {
	name      string
	generator feagen.FeatureGenerator
}

type FeagenManager struct {
	generators []generatorT
}

// Method Initialize create and initialize the feature generators as specified
// by the feature_list, return nil if all generators have been successfully
// created and initialized, error otherwise.
func (mgr *FeagenManager) Initialize(feature_list string) error {
	// Extract the list of available feature generators.

	// Use only the specified.
	generator_names := strings.Split(feature_list, ";")
	for i, name := range generator_names {
		if gen, ok := feagen.KAvailableGenearators[name]; ok {
			(*mgr).generators = append((*mgr).generators, generatorT{name, gen})
		} else {
			return fmt.Errorf("Cannot locate feature generator #%d: %s in the registry.", i, name)
		}
	}
	return nil
}

// Method GenerateFeatures generates features from the src_file and output 
// the result to out_file.
func (mgr *FeagenManager) GenerateFeatures(src_file, out_file string) error {
	in_fp, err := os.Open(src_file)
	if err != nil {
		return fmt.Errorf("Failed to open file %s: %s.\n", os.Args[1], err)
	}
	defer in_fp.Close()

	out_fp, err := os.Create(out_file)
	if err != nil {
		return fmt.Errorf("Failed to create output file %s: %s.\n", os.Args[2], err)
	}
	defer out_fp.Close()

	csv_reader := csv.NewReader(in_fp)
	tab_writer := bufio.NewWriter(out_fp)
	for {
		fields, err := csv_reader.Read()
		if err == nil {
			//TODO(weidoliang): add field preprocessor and parallel feature
			//generation using go routine?
			for _, gen := range mgr.generators {
				features := gen.generator.GenerateFeature(fields)
				for _, f := range features {
					tab_writer.WriteString(f + "\t")
				}
			}
			tab_writer.WriteByte('\n')
		} else if err == io.EOF {
			break
		} else {
			log.Printf("Failed to process record: %s.\n", err)
		}
	}
	return nil
}
