package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func main() {
	body, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := bytes.Split(body, []byte("\n"))

	times := string(bytes.Split(lines[0], []byte(":"))[1])
	// remove spaces from times
	times = strings.ReplaceAll(times, " ", "")
	timeInt, err := strconv.Atoi(times)
	if err != nil {
		panic(err)
	}
	distances := string(bytes.Split(lines[1], []byte(":"))[1])
	distances = strings.ReplaceAll(distances, " ", "")
	distancesInt, err := strconv.Atoi(distances)

	race := Race{time: timeInt, distance: distancesInt}

	fmt.Println("race", race)
	total := 0
	for time := race.time - 1; time > 0; time-- {
		speed := race.time - time
		remaining := race.time - speed
		distance := speed * remaining

		// fmt.Println("speed", speed, "remaining", remaining, "distance", distance, "race-distance", race.distance)

		if distance > race.distance {
			total++
		}

	}

	fmt.Println(total)

}
