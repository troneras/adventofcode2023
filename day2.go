package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
)

const (
	total_red   = 12
	total_blue  = 14
	total_green = 13
)

type Game struct {
	max_red   int
	max_blue  int
	max_green int
}

func main3() {
	part1()

}

func part2() {

	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// var total int = 0

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	num := 0
	total := 0

	for filescanner.Scan() {
		line := filescanner.Text()
		num++
		line = strings.Split(line, ":")[1]

		sets := strings.Split(line, ";")

		max_red := 0
		max_blue := 0
		max_green := 0

		for _, set := range sets {
			re := regexp.MustCompile(`(\d+) (red|blue|green)`)

			matches := re.FindAllStringSubmatch(set, -1)
			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]

				red := 0
				blue := 0
				green := 0

				switch color {
				case "red":
					red += count
				case "blue":
					blue += count
				case "green":
					green += count
				}

				if red > max_red {
					max_red = red
				}

				if blue > max_blue {
					max_blue = blue
				}

				if green > max_green {
					max_green = green
				}

			}
		}

		total += max_red*max_blue + max_green

	}

	if err := filescanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}
}

func part1() {

	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// var total int = 0

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	num := 0
	total := 0

	for filescanner.Scan() {
		line := filescanner.Text()
		num++
		line = strings.Split(line, ":")[1]

		sets := strings.Split(line, ";")

		all_valid := true
		for _, set := range sets {
			if !isValidSet(set) {
				all_valid = false
				break
			}
		}

		if all_valid {
			total += num
		}

	}

	if err := filescanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}

	fmt.Println(total)

}

func isValidSet(set string) bool {
	re := regexp.MustCompile(`(\d+) (red|blue|green)`)

	red := 0
	blue := 0
	green := 0

	matches := re.FindAllStringSubmatch(set, -1)
	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]

		switch color {
		case "red":
			red += count
		case "blue":
			blue += count
		case "green":
			green += count
		}
	}

	if red > total_red || blue > total_blue || green > total_green {
		return false
	}

	return true

}
