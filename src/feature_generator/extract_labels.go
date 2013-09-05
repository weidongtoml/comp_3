// Copyright 2013 Weidong Liang. All rights reserved.

package feature_generator

import (
	"strings"
)

type ExtractLabelsGenerator struct{}

func (g ExtractLabelsGenerator) Init() bool {
	return true
}

func (g ExtractLabelsGenerator) GenerateFeature(fields []string) []string {
	label_field := fields[3]
	return strings.Split(lable_field)
}

func (g ExtractLabelsGenerator) UnInit() {

}
