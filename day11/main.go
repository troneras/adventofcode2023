package main

import (
	"bytes"
	"fmt"
	"os"
)

type Galaxy struct {
	n int
	x int
	y int
}

type GalaxyPair struct {
	Galaxy1, Galaxy2 Galaxy
}

func main() {
	body, _ := os.ReadFile("input.txt")
	expansion := 999999

	lines := bytes.Split(body, []byte{'\n'})

	emptyColumns := make(map[int]int, len(lines[0]))
	empty := 0

	for i := 0; i < len(lines[0]); i++ {
		found := false

		emptyColumns[i] = empty

		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				found = true
				break
			}
		}
		if !found {
			empty = empty + expansion
		}
	}

	galaxies := []Galaxy{}
	distanceMatrix := make(map[GalaxyPair]int)

	i := 0
	empty = 0
	emptyLines := make(map[int]int)
	for {
		emptyLines[i] = empty
		// is empty
		if bytes.IndexByte(lines[i], '#') < 0 {
			empty = empty + expansion
			// fmt.Println(string(linesCopy[j]))
		} else {
			// find the position of #
			for pos, b := range lines[i] {
				if b == '#' {
					galaxies = append(galaxies, Galaxy{x: pos + emptyColumns[pos], y: i + emptyLines[i], n: len(galaxies) + 1})
				}
			}

		}
		// fmt.Println(string(linesCopy[j]))
		i++

		if i == len(lines) {
			break
		}
	}

	gl := len(galaxies)
	sum := 0
	for i := 0; i < gl; i++ {
		for j := i + 1; j < gl; j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			d := getDistance(g1, g2)
			sum += d

			// fmt.Printf("distance %v(%v,%v) - %v(%v,%v) : %d\n", g1.n, g1.y, g1.x, g2.n, g2.y, g2.x, d)
			distanceMatrix[GalaxyPair{g1, g2}] = d
		}
	}

	fmt.Println("total: ", sum)
}

// Manhattan Distance for each pair
func getDistance(g1 Galaxy, g2 Galaxy) int {
	return abs(g1.x-g2.x) + abs(g1.y-g2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
