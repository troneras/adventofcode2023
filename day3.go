package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sol2()
}

func sol2() {
	file, err := os.Open("input3.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	schematic := []string{}
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		schematic = append(schematic, filescanner.Text())
	}

	total := 0

	for y, row := range schematic {
		//fmt.Println("y: ", y, "row: ", row)

		for x := 0; x < len(row); x++ {
			char := row[x]
			if char == '*' {
				type coord struct{ y, x int }
				coords := make(map[coord]struct{})
				//				fmt.Println("Found a * at: ", x, y)
				for j := y - 1; j <= y+1; j++ {
					for i := x - 1; i <= x+1; i++ {
						if j < 0 || i < 0 || j >= len(schematic) || i >= len(schematic[j]) {
							continue
						}
						//fmt.Printf("Checking: %d, %d  =  %c \n", i, j, schematic[j][i])
						if unicode.IsDigit(rune(schematic[j][i])) {
							startx := i
							for startx >= 0 && unicode.IsDigit(rune(schematic[j][startx])) {
								startx--
							}
							startx++

							//fmt.Printf("Found adjacent number: %c\n", schematic[j][startx])
							coords[coord{j, startx}] = struct{}{}
						}
					}
				}

				if len(coords) < 2 {
					continue
				}

				//fmt.Println("coords: ", coords)

				partial := 1
				for c := range coords {
					cx := c.x
					cy := c.y
					//fmt.Println("cx: ", cx, "cy: ", cy)
					endX := cx
					for endX < len(schematic[cy]) && unicode.IsDigit(rune(schematic[cy][endX])) {
						endX++
					}

					//fmt.Println("cx: ", cx, "cy: ", cy, "endX: ", endX)

					number, _ := strconv.Atoi(schematic[cy][cx:endX])
					//fmt.Println("number: ", number)
					partial *= number
				}

				//fmt.Println("partial: ", partial)

				total += partial
			}
		}
	}

	fmt.Println("Total: ", total)

}

func sumPartNumbers2(schematic []string) int {
	sume := 0
	for y, row := range schematic {
		for x := 0; x < len(row); x++ {
			char := row[x]
			if char == '*' {
				partial := findOperand(x, y, schematic)
				if partial != -1 {
					sume += partial
				}
			}
		}
	}
	return sume

}

func findOperand(x, y int, schematic []string) int {
	// Find the first number
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	adjacent_numbers := []int{}
	coords := make(map[int]int)

	for _, dir := range dirs {
		// get the adjacent numbers
		startX := x + dir.dx
		startY := y + dir.dy
		if startX >= 0 && startY >= 0 && startY < len(schematic) && startX < len(schematic[startY]) {
			if unicode.IsDigit(rune(schematic[startY][startX])) {
				println("found adjacent number: ", rune(schematic[startY][startX]))
				startX2 := startX

				for startX2 >= 0 && unicode.IsDigit(rune(schematic[startY][startX2])) {
					startX2--
				}
				startX2++

				coords[startY] = startX2
			}
		}
	}

	if len(coords) < 2 {
		return -1
	}

	for i, v := range coords {
		println("i: ", i, "v: ", v, "schematic[v][i]: ", schematic[i][v])
		startX := i
		for startX < len(schematic[v]) && unicode.IsDigit(rune(schematic[v][i])) {
			startX++
		}
		startX--

		number, _ := strconv.Atoi(schematic[v][i:startX])
		adjacent_numbers = append(adjacent_numbers, number)
	}

	// multiply the adjacent numbers
	total := 1
	for _, num := range adjacent_numbers {
		total *= num
	}
	return total
}

func sol1() {
	// schematic := []string{
	// 	"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	"......755.",
	// 	"...$.*....",
	// 	".664.598..",
	// }
	file, err := os.Open("input3.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// var total int = 0

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	schematic := []string{}
	for filescanner.Scan() {
		schematic = append(schematic, filescanner.Text())
	}

	sum := sumPartNumbers(schematic)
	fmt.Printf("Sum of part numbers: %d\n", sum)
}

func sumPartNumbers(schematic []string) int {
	sum := 0
	for y, row := range schematic {
		for x := 0; x < len(row); x++ {
			char := row[x]
			if unicode.IsDigit(rune(char)) {
				endX := x
				for endX < len(row) && unicode.IsDigit(rune(row[endX])) {
					endX++
				}

				number, _ := strconv.Atoi(row[x:endX])
				if isAdjacentToSymbol(x, y, endX, schematic) {
					sum += number
				}

				x = endX - 1 // Skip past the processed number
			}
		}
	}
	return sum
}

func isAdjacentToSymbol(startX, y, endX int, schematic []string) bool {
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, dir := range dirs {
		for x := startX; x < endX; x++ {
			newX, newY := x+dir.dx, y+dir.dy
			if newX >= 0 && newY >= 0 && newY < len(schematic) && newX < len(schematic[newY]) {
				adjacentChar := rune(schematic[newY][newX])
				if !unicode.IsDigit(adjacentChar) && adjacentChar != '.' {
					return true
				}
			}
		}
	}
	return false
}
