package world

import (
	"fmt"
	"strings"

	log "github.com/daesu/monster/logging"
)

// Const variables
const (
	maxMoves         = 10000 // max moves per monster
	monsterNameChars = 8     // number of chars in monster name
	dashLen          = 80    // number of dashes to print for formating
)

// Rampage ...
func Rampage(cities []City, numMons int) ([]City, map[string]Monster) {

	// Prevent user adding more monsters than cities
	if numMons > len(cities) {
		err := fmt.Sprintf("Number of monsters %d cannot be greater than number of cities %d", numMons, len(cities))
		fmt.Println(err)
		log.Fatal(err)
	}

	monsters := GetMonsters(numMons, cities)

	// For a maximum of maxMoves per monster
	for move := 0; move <= maxMoves; move++ {

		msgs := []Monster{}
		for _, monster := range monsters {

			// Ignore dead or trapped monsters
			if !monster.Dead && !monster.Trapped {
				monsters[monster.ID] = Roam(monster)
				msgs = append(msgs, monsters[monster.ID])
			}
		}

		// For every destroyed city, kill the monsters
		destroyedCities := GetDestroyedCities(msgs)
		for _, city := range destroyedCities {

			for _, monster := range city.Monsters {
				log.Debug(fmt.Sprintf("killing monster %s", monster.ID))

				mons := monsters[monster.ID]
				mons.Dead = true
				monsters[monster.ID] = mons
			}

			fmt.Println(fmt.Sprintf("%s", strings.Repeat("-", dashLen)))
			fmt.Println(fmt.Sprintf("%s has been destroyed by monster %s and monster %s!", city.City.Name, city.Monsters[0].Name, city.Monsters[1].Name))

			// Presume that any other monsters in the city
			// will also be killed when the city has been destroyed.
			if len(city.Monsters) > 2 {
				monsters := []string{}
				for i := 2; i < len(city.Monsters); i++ {
					monsters = append(monsters, fmt.Sprintf("%s", city.Monsters[i].Name))
				}

				monstersStr := strings.Join(monsters[:], ", ")
				fmt.Println(fmt.Sprintf("Also %s died in the collateral damage", monstersStr))
			}

			city.City.Destroyed = true
		}
	}

	return cities, monsters
}

// PrintStats ...
func PrintStats(cities []City, monsters map[string]Monster) {
	fmt.Println(fmt.Sprintf("%s", strings.Repeat("-", dashLen)))
	fmt.Println("World map with remaining cities")

	cityValidRouteCount := 0
	cityBlockedCount := 0
	destroyedCityCount := 0
	for _, city := range cities {
		if city.Destroyed == false {
			roads := GetValidRoutesFormatted(city)

			fmt.Println(fmt.Sprintf("%s %s", city.Name, roads))

			if roads == "" {
				cityBlockedCount = cityBlockedCount + 1
			} else {
				cityValidRouteCount = cityValidRouteCount + 1
			}
		} else {
			destroyedCityCount = destroyedCityCount + 1
		}
	}

	activeMonsters := GetMonsterCounts(monsters)

	fmt.Println(fmt.Sprintf("%s", strings.Repeat("-", dashLen)))
	fmt.Println(fmt.Sprintf("active monsters: %d", activeMonsters[`active`]))
	fmt.Println(fmt.Sprintf("dead monsters: %d", activeMonsters[`dead`]))
	fmt.Println(fmt.Sprintf("trapped monsters: %d", activeMonsters[`trapped`]))

	fmt.Println(fmt.Sprintf("%s", strings.Repeat("-", dashLen)))
	fmt.Println(fmt.Sprintf("remaining cities with valid routes: %d", cityValidRouteCount))
	fmt.Println(fmt.Sprintf("remaining cities cut off from routes: %d", cityBlockedCount))
	fmt.Println(fmt.Sprintf("destroyed cities: %d", destroyedCityCount))
}

// GetMonsters ...
func GetMonsters(numMons int, cities []City) map[string]Monster {

	// Copy of cities slice data to be used to randomly
	// insert monsters into cities without duplicates
	freeCities := make([]City, len(cities))
	copy(freeCities, cities)

	// create monsters
	monsters := make(map[string]Monster)
	for i := 0; i < numMons; i++ {

		// Pick a random city not already occupied
		// by a monster to insert a new monster into
		randInd := Random(0, len(freeCities))
		index, err := FindCityIndexByName(cities, freeCities[randInd].Name)
		if err != nil {
			log.Debug(err.Error())
			//TODO
		}

		city := &cities[index]

		// Remove city from slice as it has being occupied by a monster
		freeCities = append(freeCities[:randInd], freeCities[randInd+1:]...)

		// Generate monster with random name/ID
		monster := Monster{ID: GoodEnoughUUID(), Name: RandomMonsterName(monsterNameChars), City: city}
		monsters[monster.ID] = monster
	}

	return monsters
}

// Roam travels the City map if valid routes exist and return
// the current City name and Monster name
func Roam(monster Monster) Monster {

	randDirection, err := RandomDirection(monster.City)
	if err != nil {
		log.Debug(fmt.Sprintf("No valid direction detected. monster %s is trapped", monster.Name))
		monster.Trapped = true
		return monster
	}

	monster.Moves++
	log.Debug(fmt.Sprintf("Monster %s has moved %d times", monster.Name, monster.Moves))
	log.Debug(fmt.Sprintf("Monster %s at city %s and roaming %s", monster.Name, monster.City.Name, randDirection))

	switch randDirection {
	case North:
		monster.City = monster.City.North
	case South:
		monster.City = monster.City.South
	case East:
		monster.City = monster.City.East
	case West:
		monster.City = monster.City.West
	}

	return monster
}
