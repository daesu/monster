package world

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	log "github.com/daesu/monster/logging"
)

// consts
const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CityNameLength = 10
)

// GetMonsterCounts returns a map containing the current state of
// monsters. "active", "dead", or "trapped".
func GetMonsterCounts(monsters map[string]Monster) map[string]int {

	mCount := make(map[string]int)
	for _, v := range monsters {
		if v.Dead {
			mCount[`dead`]++
		} else if v.Trapped {
			mCount[`trapped`]++
		} else {
			mCount[`active`]++
		}
	}

	return mCount
}

// RandString returns a random string contain n runes
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

// Random returns a random number within the range min to max
func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomDirection accepts a City and returns a valid
// route to another City if it exists
func RandomDirection(city *City) (string, error) {
	validRoads := []string{}

	if city.North != nil && city.North.Destroyed == false {
		validRoads = append(validRoads, North)
	}

	if city.South != nil && city.South.Destroyed == false {
		validRoads = append(validRoads, South)
	}

	if city.East != nil && city.East.Destroyed == false {
		validRoads = append(validRoads, East)
	}

	if city.West != nil && city.West.Destroyed == false {
		validRoads = append(validRoads, West)
	}

	log.Debug(fmt.Sprintf("Valid roads from %s are: %+v\n", city.Name, validRoads))

	if len(validRoads) < 1 {
		err := fmt.Errorf("no valid routes available from city %s", city.Name)
		return "", err
	}

	randDirection := validRoads[Random(0, len(validRoads))]

	return randDirection, nil
}

// GetDuplicatesFromSlice returns slice of duplicate Cities
// from a slice of ChannelMessage objects.
func GetDuplicatesFromSlice(s []Monster) map[string]int {

	duplicateFrequency := make(map[string]int)
	for _, item := range s {
		_, exist := duplicateFrequency[item.City.Name]

		if exist {
			duplicateFrequency[item.City.Name]++

		} else {
			duplicateFrequency[item.City.Name] = 1
		}
	}

	return duplicateFrequency
}

// GetDestroyedCities returns slice of []DestroyedCity
// from a slice of ChannelMessage objects.
func GetDestroyedCities(mons []Monster) []DestroyedCity {

	// Key map of dups with occurences
	duplicates := GetDuplicatesFromSlice(mons)

	// Loop through map:city:duplicates of cities
	destroyedCities := []DestroyedCity{}
	for k, v := range duplicates {

		// Only 1 minster is in the city
		if v < 2 {
			break
		}

		destroyedCity := DestroyedCity{}
		for _, item := range mons {
			// for every Monster
			if item.City.Name == k {
				if destroyedCity.City == nil {
					destroyedCity.City = item.City
				}

				if destroyedCity.Monsters == nil {
					monsters := []Monster{}
					destroyedCity.Monsters = monsters
				}

				destroyedCity.Monsters = append(destroyedCity.Monsters, item)
			}
		}

		destroyedCities = append(destroyedCities, destroyedCity)
	}

	return destroyedCities
}

// FindCityIndexByName Scans a slice of City objects for a city name
// and returns the index if found.
func FindCityIndexByName(cities []City, name string) (int, error) {
	for i, city := range cities {
		if city.Name == name {
			return i, nil
		}
	}

	err := fmt.Errorf("no city exists with name %s", name)
	return 0, err
}

// RandomCityName returns a randomly generated
// string of length 3 to 10
func RandomCityName() string {
	randInt := Random(3, CityNameLength)
	randStr := RandString(randInt)

	return string(randStr)
}

// RandomMonsterName returns a randomly generated string int combination
// of n+1 length
func RandomMonsterName(n int) string {
	randStr := RandString(n)
	randInt := Random(0, n)

	return string(randStr) + strconv.Itoa(randInt)
}

// GetRoute accepts a string in the format `direction=city`
// and returns the direction and city as a []string of length 2
func GetRoute(route, delim string) ([]string, error) {
	splitStr := strings.Split(route, delim)

	if len(splitStr) != 2 {
		err := fmt.Errorf("invalid format of direction detected for %s", route)
		return nil, err
	}

	return splitStr, nil
}

// StringInSlice returns bool indicating if a string exists in a list
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetRandomDirections returns set of
// non-duplicate random compass directions
func GetRandomDirections() []string {
	randInt := Random(0, 2)

	mockCompass := make([]string, len(Compass))
	copy(mockCompass, Compass)

	for i := 0; i <= randInt; i++ {
		randInd := Random(0, len(mockCompass))
		mockCompass = append(mockCompass[:randInd], mockCompass[randInd+1:]...)
	}

	return mockCompass
}

// GetValidRoutesFormatted takes a City object and generates
// formatted string representing valid routes from the city
func GetValidRoutesFormatted(city City) string {
	roads := ""
	if city.North != nil && !city.North.Destroyed {
		northRoad := fmt.Sprintf("%s=%s", North, city.North.Name)
		if roads == "" {
			roads = northRoad
		} else {
			roads = fmt.Sprintf("%s %s", roads, northRoad)
		}
	}
	if city.South != nil && !city.South.Destroyed {
		southRoad := fmt.Sprintf("%s=%s", South, city.South.Name)
		if roads == "" {
			roads = southRoad
		} else {
			roads = fmt.Sprintf("%s %s", roads, southRoad)
		}
	}
	if city.East != nil && !city.East.Destroyed {
		eastRoad := fmt.Sprintf("%s=%s", East, city.East.Name)
		if roads == "" {
			roads = eastRoad
		} else {
			roads = fmt.Sprintf("%s %s", roads, eastRoad)
		}
	}
	if city.West != nil && !city.West.Destroyed {
		westRoad := fmt.Sprintf("%s=%s", West, city.West.Name)
		if roads == "" {
			roads = westRoad
		} else {
			roads = fmt.Sprintf("%s %s", roads, westRoad)
		}
	}

	return roads
}

// GoodEnoughUUID generates a simple UUID
func GoodEnoughUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
