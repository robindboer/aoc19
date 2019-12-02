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

	// calculateFuelRequirement
	var totalFuelRequirement int
	for _, i := range masses {
		totalFuelRequirement += calculateFuelRequirement(i)
	}

	fmt.Println(totalFuelRequirement)
}

func calculateFuelRequirement(mass int) int {
	return int((mass / 3) - 2)
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
