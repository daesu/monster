package world

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/daesu/monster/logging"
)

// const values
const (
	North = "north"
	South = "south"
	East  = "east"
	West  = "west"
)

// Compass ...
var Compass = []string{
	North, South, East, West,
}

// GetCitiesFromFile reads the provided world map
// and parses the data into City structs
func GetCitiesFromFile(filename string) ([]City, error) {

	// Open the file
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	r.Comma = ' '
	r.FieldsPerRecord = -1

	// Iterate through the records
	cities := []City{}
	records := [][]string{}
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// Save records to iterate over and generate routes
		records = append(records, record)

		// Create City Node without routes
		city := City{}

		// Presume City names are unique
		city.Name = record[0]
		cities = append(cities, city)
	}

	// generate City routes
	for _, record := range records {
		for c, city := range cities {
			if city.Name == record[0] {

				// Set Up routes for city
				for i := 1; i < len(record); i++ {

					direction := strings.Split(record[i], "=")

					cityIndex, _ := FindCityIndexByName(cities, direction[1])
					if err != nil {
						log.Debug(fmt.Sprintf("No city found with name %s", direction[1]))
						continue
					}

					var nCity = &cities[cityIndex]

					if direction[0] == North {
						cities[c].North = nCity
					} else if direction[0] == South {
						cities[c].South = nCity
					} else if direction[0] == East {
						cities[c].East = nCity
					} else if direction[0] == West {
						cities[c].West = nCity
					}
				}

				break
			}
		}
	}

	return cities, nil
}

// GenerateRandomCitiesMap ...
func GenerateRandomCitiesMap(numCities int) ([]City, error) {

	randomCityNames := []string{}
	for i := 0; i < numCities; i++ {
		for {
			name := RandomCityName()

			if !StringInSlice(name, randomCityNames) {
				randomCityNames = append(randomCityNames, name)
				break
			}
		}
	}

	// Create City Node without routes
	cities := []City{}
	for i := 0; i < numCities; i++ {
		city := City{}
		city.Name = randomCityNames[i]
		cities = append(cities, city)
	}

	// generate City routes
	for c := range cities {

		// Copy of cities slice data to be used to randomly
		// grab a city without repeating ourselves
		freeCities := make([]City, len(cities))
		copy(freeCities, cities)

		freeCities = append(freeCities[:c], freeCities[c+1:]...)

		// Set Up routes for city
		directions := GetRandomDirections()
		for _, direction := range directions {

			// Pick a random city
			randInd := Random(0, len(freeCities))
			index, err := FindCityIndexByName(freeCities, freeCities[randInd].Name)
			if err != nil {
				log.Debug(err.Error())
				//TODO
			}

			var nCity = &cities[index]

			// Remove current city from slice
			freeCities = append(freeCities[:index], freeCities[index+1:]...)

			switch direction {
			case North:
				cities[c].North = nCity
			case South:
				cities[c].South = nCity
			case East:
				cities[c].East = nCity
			case West:
				cities[c].West = nCity
			}
		}
	}

	return cities, nil
}
