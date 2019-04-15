# monster
Simple command line go program to parse a file into nodes and navigate node routes. 

### Quick Start:
 - Clone this repo to Go path e.g. `~/src/github.com/daesu/monster`

 - Run `make run`.

   Make will install dependencies and run the app with the default world map provided.

### Usage:
The app accepts three arguments;

   - -f string

    	specify world map filename (default "data/world.txt")

   - -g	generate map dynamically 

   Generates N cities and routes between them based on the number of monsters. (default 5000)

   - -m int
    	specify number of monsters (default 5000)

### Tests:
Simple tests are included to verify City struct parsing.

 - `make test` 