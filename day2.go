package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var debug = flag.Bool("debug", false, "enable debug")
var input = flag.String("input", "input.txt", "input file name")
var partB = flag.Bool("partB", false, "whether to use part B logic")

const finalValue = 19690720

func main() {

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Error().Msg(err.Error())
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
	if !*partB {
		tapeCopy := make([]int, len(tape))
		copy(tapeCopy, tape)
		tapeCopy[1] = 12
		tapeCopy[2] = 2
		fmt.Println(process(tapeCopy))
	} else {
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				tapeCopy := make([]int, len(tape))
				copy(tapeCopy, tape)
				tapeCopy[1] = noun
				tapeCopy[2] = verb
				if process(tapeCopy) == finalValue {
					log.Info().Msgf("noun %d, verb %d, result %d", noun, verb, 100*noun+verb)
					return
				}
			}
		}
		info("failed")
	}
}

func process(tape []int) int {
	offset := 0
	for {
		if offset >= len(tape) {
			info("offset if bigger than tape length")
			return -1
		}

		if tape[offset] == 99 {
			return tape[0]
		}

		a := tape[offset+1]
		b := tape[offset+2]
		dstOffset := tape[offset+3]
		offsetSize := 0

		switch tape[offset] {
		case 1:
			tape[dstOffset] = tape[a] + tape[b]
			offsetSize = 4
		case 2:
			tape[dstOffset] = tape[a] * tape[b]
			offsetSize = 4
		default:
			info("strange opcode")
			return -1
		}
		offset += offsetSize
	}
}

func info(msg string) {
	log.Info().Msg(msg)
}
