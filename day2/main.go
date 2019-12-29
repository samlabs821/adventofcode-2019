package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var debug = flag.Bool("debug", false, "enable debug")

func main() {

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return
	}

	contents := string(bytes[:len(bytes)-2])
	split := strings.Split(contents, ",")
	tape := make([]int, len(split))

	for i, s := range split {
		n, _ := strconv.Atoi(s)
		tape[i] = n
	}

	for _, n := range tape {
		log.Debug().Msgf("%d ", n)
	}

	// init
	tape[1] = 12
	tape[2] = 2

	offset := 0
	for {
		if offset >= len(tape) {
			info("offset if bigger than tape length")
			return
		}

		if tape[offset] == 99 {
			log.Info().Msgf("0 index: %d. exited normally", tape[0])
			return
		}

		a := tape[offset+1]
		b := tape[offset+2]
		dstOffset := tape[offset+3]

		switch tape[offset] {
		case 1:
			tape[dstOffset] = tape[a] + tape[b]
		case 2:
			tape[dstOffset] = tape[a] * tape[b]
		default:
			info("strange opcode")
		}
		offset += 4
	}
}

func info(msg string) {
	log.Info().Msg(msg)
}
