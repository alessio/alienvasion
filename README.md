# alienvasion

## Getting Started

Just go get-and-install and run it:

```
$ go get github.com/alessio/alienvasion/...
$ go install github.com/alessio/alienvasion/cmd/alienvasion
$ alienvasion -help
Mad aliens are about to invade the earth and this
program aims to simulate the invasion.

Other than the flags described below, this program reads a map from
STDIN. A map contains the names of cities in the non-existent world of
X, one city each line. The city name comes first, followed by 1-4
directions (north, south, east, or west). Each one represents a
road to another city that lies in that direction.
The city and each of the pairs are separated by a single space, and
the directions are separated from their respective cities with an
equals (=) sign.

N aliens are created and given random name, where N is specified as
an optional command-line argument. These aliens start out at random
places on the map, and wander around randomly, following links. Each
iteration, the aliens can travel in any of the directions leading out
of a city.

When multiple aliens end up in the same place, they fight, and in the
process kill each other and destroy the city. When a city is destroyed,
it is removed from the map, and so are any roads that lead into or out
of it. Once a city is destroyed, aliens can no longer travel to or
through it. This may lead to aliens getting "trapped".

  -file string
    	read map from file instead of STDIN
  -maxmoves int
    	max number of moves each alien can make (default 10000)
  -naliens int
    	generate N aliens (default 8)

```
