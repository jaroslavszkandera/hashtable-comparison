package main

import (
	"hashtablecomparison/linearprobing"
	"hashtablecomparison/robinhood"
	"hashtablecomparison/separatechaining"
	"hashtablecomparison/testdata"
	"testing"
)

const TEST_DATA_COUNT = 10000000

func simpleHash_(val string) uint {
	var hash uint
	for i, char := range val {
		hash += uint(char) * uint(i*i*i*i)
	}
	return hash
}

func BenchmarkRobinhoodInsert(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	robinhood := robinhood.NewHashTable(simpleHash_)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			robinhood.Insert(val)
		}
	}
}

func BenchmarkLinearProbingInsert(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	linearprobing := linearprobing.NewHashTable(simpleHash_)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			linearprobing.Insert(val)
		}
	}
}

func BenchmarkSeparateChainingInsert(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	separatechaining := separatechaining.NewHashTable(simpleHash_)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			separatechaining.Insert(val)
		}
	}
}

func BenchmarkRobinhoodSearch(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	robinhood := robinhood.NewHashTable(simpleHash_)
	for _, val := range data {
		robinhood.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			robinhood.Search(val)
		}
	}
}

func BenchmarkLinearProbingSearch(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	linearprobing := linearprobing.NewHashTable(simpleHash_)
	for _, val := range data {
		linearprobing.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			linearprobing.Search(val)
		}
	}
}

func BenchmarkSeparateChainingSearch(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	separatechaining := separatechaining.NewHashTable(simpleHash_)
	for _, val := range data {
		separatechaining.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			separatechaining.Search(val)
		}
	}
}

func BenchmarkRobinhoodDelete(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	robinhood := robinhood.NewHashTable(simpleHash_)
	for _, val := range data {
		robinhood.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			robinhood.Delete(val)
		}
	}
}

func BenchmarkLinearProbingDelete(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	linearprobing := linearprobing.NewHashTable(simpleHash_)
	for _, val := range data {
		linearprobing.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			linearprobing.Delete(val)
		}
	}
}

func BenchmarkSeparateChainingDelete(b *testing.B) {
	data := testdata.GenerateTestData(TEST_DATA_COUNT)
	separatechaining := separatechaining.NewHashTable(simpleHash_)
	for _, val := range data {
		separatechaining.Insert(val)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, val := range data {
			separatechaining.Delete(val)
		}
	}
}
