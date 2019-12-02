package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := flag.String("i", "", "input file")

	flag.Parse()

	// read the mass values from file
	masses := readMassValuesFromFile(inputPath)

	var totalFuelRequirement int
	var moduleFuelRequirement int
	for _, i := range masses {
		totalFuelRequirement += totalFuel(i)
		moduleFuelRequirement += massToFuel(i)
	}

	fmt.Println("total:", totalFuelRequirement)
	fmt.Println("module fuel:", moduleFuelRequirement)
}

func massToFuel(mass int) int {
	return int((mass / 3) - 2)
}

func totalFuel(mass int) int {
	total := massToFuel(mass)

	for f := massToFuel(total); f > 0; f = massToFuel(f) {
		total += f
	}

	return total
}

func readMassValuesFromFile(fp *string) []int {
	dat, err := ioutil.ReadFile(*fp)

	if err != nil {
		fmt.Println("could not read:", err)
		os.Exit(1)
	}

	str := string(dat)

	var masses []int

	for _, i := range strings.Split(str, "\n") {
		mass, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		masses = append(masses, mass)
	}

	return masses
}
