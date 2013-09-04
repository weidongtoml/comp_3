// Copyright 2013 Weidong Liang. All rights reserved.

package feature_generator

type FeatureGenerator interface {
	Init() bool
	GenerateFeature(fields []string) []string
	UnInit()
}

var KAvailableGenearators map[string]FeatureGenerator

func Init() {
	KAvailableGenearators["MostFrequentWordGenerator"] = new(MostFrequentWordGenerator)
}
