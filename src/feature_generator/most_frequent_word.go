// Copyright 2013 Weidong Liang. All rights reserved.

package feature_generator

type MostFrequentWordGenerator struct{}

func (g MostFrequentWordGenerator) Init() bool {
	return true
}

func (g MostFrequentWordGenerator) GenerateFeature(fields []string) []string {
	return []string{"Hello"}
}

func (g MostFrequentWordGenerator) UnInit() {

}
