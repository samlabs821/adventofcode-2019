package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func checkWarn(e error) {
	if e != nil {
		log.Info().Msg(e.Error())
	}
}

func checkFatal(e error) {
	if e != nil {
		panic(e)
	}
}

func scanFile(input string) ([]int64, error) {
	var err error
	a := make([]int64, 0)
	b, err := ioutil.ReadFile(input)
	checkFatal(err)

	for _, d := range strings.Split(string(b), ",") {
		if len(d) > 0 {
			i, err := strconv.ParseInt(d, 10, 64)
			if err != nil {
				log.Info().Msg(err.Error())
			}
			a = append(a, i)
		}
	}

	return a, nil
}
