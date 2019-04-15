package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/daesu/monster/world"
)

func main() {
	rand.Seed(time.Now().Unix())

	mapString := flag.String("f", "data/world.txt", "specify world map filename")
	monstersInt := flag.Int("m", 5000, "specify number of monsters")
	generateBool := flag.Bool("g", false, "generate map dynamically")
	flag.Parse()

	if *mapString == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var mapCities = []world.City{}
	if *generateBool {
		mapWorld, err := world.GenerateRandomCitiesMap(*monstersInt)
		if err != nil {
			log.Fatal("could not generate cities")
		}
		mapCities = mapWorld
	} else {
		fmt.Println(fmt.Sprintf("getting map from file %s", *mapString))
		mapWorld, err := world.GetCitiesFromFile(*mapString)
		if err != nil {
			log.Fatal("could not read cities from provided file")
		}
		mapCities = mapWorld
	}

	cities, monsters := world.Rampage(mapCities, *monstersInt)

	world.PrintStats(cities, monsters)

}
