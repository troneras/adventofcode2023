package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const (
	DAMAGED     = '#'
	OPERATIONAL = '.'
	UNKNOWN     = '?'
)

func main() {
	body, _ := os.ReadFile("test.txt")
	fmt.Println(string(body))
	fmt.Println("----")

	lines := bytes.Split(body, []byte("\n"))

	n := 0
	for _, line := range lines {
		tmp_groups := []int{}
		parts := bytes.Split(line, []byte(" "))
		real_groups := []int{}
		partials := []string{}

		for _, c := range bytes.Split(parts[1], []byte(",")) {
			i, _ := strconv.Atoi(string(c))
			real_groups = append(real_groups, i)
		}

		start := 0
		for z, spring := range parts[0] {
			if spring == UNKNOWN || spring == DAMAGED {
				n++
			}
			if spring == OPERATIONAL {
				if n > 0 {
					tmp_groups = append(tmp_groups, n)
					partials = append(partials, string(parts[0][start:z]))
					start = z + 1
					n = 0
				}
			}
		}
		if n > 0 {
			tmp_groups = append(tmp_groups, n)
			partials = append(partials, string(parts[0][start:len(parts[0])]))
		}

		possible_positions := []int{}
		if len(tmp_groups) != len(real_groups) {
			// if some of the groups are the same, check if the other are ok

		} else {
			for i, v := range tmp_groups {
				possible_positions = append(possible_positions, getPossiblePositions(real_groups[i], v, partials[i]))
			}
		}
		fmt.Println(string(line), tmp_groups, possible_positions, partials)

	}

}

// partials contains UNKNOWN and DAMAGED.
// we know that damaged are in groups, and each group is separated by at least one OPERATIONAL
// ex:
//
//	damaged = 2, space = 3, partial = "???" => 2   (##. or .##)
//	damaged = 2, space = 3, partial = "??#" => 1   (because we know the last one is damaged)
//	damaged = 2, space = 4, partial = "??##" => 1 (the real would be ..##)
//	damaged = 2, space = 4, partial = "??#?" => 2 (the real would be .##. or ?.##)
func getPossiblePositions(damaged, space int, partial string) int {
	if damaged == space {
		return 1
	}

	return abs(damaged-space) + 1

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
