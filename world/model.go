package world

// City represents a Node
type City struct {
	Name      string
	North     *City
	South     *City
	East      *City
	West      *City
	Destroyed bool
}

// Monster represents an individual monster
type Monster struct {
	ID      string
	Name    string
	City    *City
	Moves   int
	Dead    bool
	Trapped bool
}

// DestroyedCity ...
type DestroyedCity struct {
	City     *City
	Monsters []Monster
}
