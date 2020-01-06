package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var debug = flag.Bool("debug", false, "enable debug")
var partB = flag.Bool("partB", false, "is this part B")
var input = flag.String("input", "input.txt", "select input file")

type Coord struct {
	X, Y int
}

type Tracker [2]float64

func main() {
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	hit := make(map[Coord]Tracker)

	for i, l := range strings.Split(string(bytes), "\n") {
		if l == "" {
			continue
		}

		loc := Coord{0, 0}
		traversed := float64(0)

		for _, step := range strings.Split(l, ",") {
			distance, err := strconv.Atoi(step[1:])
			if err != nil {
				log.Warn().Msg(err.Error())
			}

			switch step[0] {
			case 'R':
				for n := 0; n < distance; n++ {
					loc.X++
					hits := hit[loc]
					if hits[i] != float64(0) {
						hits[i] = math.Min(hits[i], float64(n)+traversed)
					} else {
						hits[i] = float64(n) + traversed
					}
					hit[loc] = hits
				}
			case 'L':
				for n := 0; n < distance; n++ {
					loc.X--
					hits := hit[loc]
					if hits[i] != float64(0) {
						hits[i] = math.Min(hits[i], float64(n)+traversed)
					} else {
						hits[i] = float64(n) + traversed
					}
					hit[loc] = hits
				}
			case 'U':
				for n := 0; n < distance; n++ {
					loc.Y++
					hits := hit[loc]
					if hits[i] != float64(0) {
						hits[i] = math.Min(hits[i], float64(n)+traversed)
					} else {
						hits[i] = float64(n) + traversed
					}
					hit[loc] = hits
				}
			case 'D':
				for n := 0; n < distance; n++ {
					loc.Y--
					hits := hit[loc]
					if hits[i] != float64(0) {
						hits[i] = math.Min(hits[i], float64(n)+traversed)
					} else {
						hits[i] = float64(n) + traversed
					}
					hit[loc] = hits
				}
			default:
				log.Error().Msg("not recognized")
			}
			traversed += float64(distance)
		}
	}

	shortest := math.MaxFloat64
	least := math.MaxFloat64
	for loc, hits := range hit {
		if hits[0] != float64(0) && hits[1] != float64(0) {
			distance := math.Abs(float64(loc.X)) + math.Abs(float64(loc.Y))
			if distance < shortest {
				shortest = distance
			}
			sum := hits[0] + hits[1]
			if sum < least {
				least = sum
			}
		}
	}

	if *partB {
		fmt.Println(least)
	} else {
		fmt.Println(shortest)
	}
}
