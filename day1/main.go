package main

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"math"
	"os"
	"strconv"
)

func checkFatal(e error) {
	if e != nil {
		panic(e)
	}
}

func checkWarn(e error) {
	if e != nil {
		log.Info().Msg(e.Error())
	}
}

func scanFile(input string) ([]float64, error) {
	var err error
	a := make([]float64, 0)
	file, err := os.Open(input)
	checkFatal(err)

	defer func() {
		err = file.Close()
		checkWarn(err)
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d := scanner.Text()
		if len(d) > 0 {
			i, err := strconv.ParseFloat(scanner.Text(), 10)
			if err != nil {
				log.Info().Msg(err.Error())
			}
			a = append(a, i)
		}
	}
	return a, nil
}

func main() {
	a, err := scanFile("input.txt")
	if err != nil {
		log.Info().Msg(err.Error())
	}
	var sum float64
	var sumWithFuel float64

	for _, v := range a {
		sum += getFuel(v)
	}

	for _, v := range a {
		sumWithFuel += calculateFuelForMass(v)
	}

	fmt.Printf("sum: %d -- sum with fuel: %d \n", int(sum), int(sumWithFuel))

}

func calculateFuelForMass(mass float64) float64 {
	var sum float64

	for {
		f := getFuel(mass)
		mass = f
		sum += f
		if f <= 0 {
			break
		}
	}
	return sum
}

func getFuel(mass float64) float64 {
	m := math.Floor(mass/3) - 2
	if m > 0 {
		return m
	}
	return 0
}
