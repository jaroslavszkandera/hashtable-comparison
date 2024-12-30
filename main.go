package main

import (
	"fmt"
	"hashtablecomparison/linearprobing"
	"hashtablecomparison/robinhood"
	"hashtablecomparison/separatechaining"
	"hashtablecomparison/testdata"
	"time"
)

func simpleHash(val string) uint {
	var hash uint
	for i, char := range val {
		hash += uint(char) * uint(i*i*i*i)
	}
	return hash
}

func main() {
	numValues := 10000000
	data := testdata.GenerateTestData(numValues)

	robinhood := robinhood.NewHashTable(simpleHash)
	linearprobing := linearprobing.NewHashTable(simpleHash)
	separatechaining := separatechaining.NewHashTable(simpleHash)

	insertData("Robinhood Hashing", robinhood.Insert, data)
	searchData("Robinhood Hashing", robinhood.Search, data)
	deleteData("Robinhood Hashing", robinhood.Delete, data)

	insertData("Linear Probing", linearprobing.Insert, data)
	searchData("Linear Probing", linearprobing.Search, data)
	deleteData("Linear Probing", linearprobing.Delete, data)

	insertData("Separate Chaining", separatechaining.Insert, data)
	searchData("Separate Chaining", separatechaining.Search, data)
	deleteData("Separate Chaining", separatechaining.Delete, data)
}

func insertData(name string, insertFn func(string) bool, data []string) {
	start := time.Now()
	inserted := 0
	for _, v := range data {
		if insertFn(v) {
			inserted++
		}
	}
	duration := time.Since(start)
	avgDuration := duration / time.Duration(len(data))
	fmt.Printf("%s: Inserted %d items in %v, average time per insert: %v\n", name, inserted, duration, avgDuration)
}

func searchData(name string, searchFn func(string) bool, data []string) {
	start := time.Now()
	found := 0
	for _, v := range data {
		if searchFn(v) {
			found++
		}
	}
	duration := time.Since(start)
	avgDuration := duration / time.Duration(len(data))
	fmt.Printf("%s: Found %d items in %v, average time per search: %v\n", name, found, duration, avgDuration)
}

func deleteData(name string, deleteFn func(string) bool, data []string) {
	start := time.Now()
	deleted := 0
	for _, v := range data {
		if deleteFn(v) {
			deleted++
		}
	}
	duration := time.Since(start)
	avgDuration := duration / time.Duration(len(data))
	fmt.Printf("%s: Deleted %d items in %v, average time per delete: %v\n", name, deleted, duration, avgDuration)
}
