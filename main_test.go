package main

import (
	"fmt"
	"testing"

	"github.com/daesu/monster/world"
)

// consts
const (
	mapString         = "data/world.txt"
	readFileCityCount = 6763
	generateCityCount = 1000
	CityNameLength    = 10
)

// TestReadDefault simple test to read data from the default world file
// if it exists and parse it to city structs
func TestReadDefault(t *testing.T) {
	mapWorld, err := world.GetCitiesFromFile(mapString)
	if err != nil {
		t.Error("Expected slice of City generated")
	}

	if len(mapWorld) != readFileCityCount {
		t.Error(fmt.Sprintf("City count was %d. Expected count was %d. ", readFileCityCount, len(mapWorld)))
	}
}

// TestGenerateMap simple test to dynamically generate city structs
func TestGenerateMap(t *testing.T) {
	mapWorld, err := world.GenerateRandomCitiesMap(generateCityCount)
	if err != nil {
		t.Error("Expected slice of City generated")
	}

	if len(mapWorld) != generateCityCount {
		t.Error(fmt.Sprintf("City count was %d. Expected count was %d. ", generateCityCount, len(mapWorld)))
	}
}
