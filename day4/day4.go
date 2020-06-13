package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var input = flag.String("input", "input.txt", "input file")
var debug = flag.Bool("debug", false, "debug usage")

var MinMaxValues [2]int

func main() {
	flag.Parse()

	if !*debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Warn().Msg(err.Error())
		return
	}

	for i, in := range strings.Split(string(bytes), "-") {
		MinMaxValues[i], err = strconv.Atoi(in)
		if err != nil {
			log.Warn().Msg(err.Error())
		}
	}
	log.Debug().Msgf("%d", MinMaxValues)
	possibleValues := 0

	for i := MinMaxValues[0]; i <= MinMaxValues[1]; i++ {
		if !validateDoubleInt(strconv.Itoa(i)) {
			continue
		}
		if !validateIncrease(strconv.Itoa(i)) {
			continue
		}
		possibleValues++
	}

	log.Info().Msgf("possible values: %d", possibleValues)

}

func validateDoubleInt(s string) bool {
	var prev rune = 0
	runs := 1
	var exactDouble bool
	for _, j := range s {
		if prev == j {
			runs++
		} else {
			if runs == 2 {
				exactDouble = true
			}
			runs = 1
		}
		prev = j
	}

	if runs == 2 {
		exactDouble = true
	}

	return exactDouble
}

func validateIncrease(s string) bool {
	prevInt := 0

	for i := 0; i < len(s); i++ {

		d, err := strconv.Atoi(string(s[i]))
		if prevInt > d {
			return false
		}
		prevInt = d
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}

	return true
}
