package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X     int
	Y     int
	nextX int
	nextY int
}

var visitedNodes [][]bool

func main() {
	body, _ := os.ReadFile("input.txt")
	fmt.Println(string(body))

	// find character "S" position in body
	start_position := strings.Index(string(body), "S")

	lines := strings.Split(string(body), "\n")

	square_width, square_height := len(lines[0]), len(lines)

	startX := (start_position % square_width) - 1
	startY := start_position / square_width
	visitedNodes = make([][]bool, square_height)
	for i := 0; i < square_height; i++ {
		for j := 0; j < square_width; j++ {
			if lines[i][j] == 'S' {
				startX = j
				startY = i
			}
		}
		visitedNodes[i] = make([]bool, square_width)
	}

	fmt.Println("width:", square_width, "height: ", square_height, "start:", start_position)

	fmt.Println("startX", startX, "startY", startY)

	visitedNodes[startY][startX] = true

	points := make(map[Point]bool)

	for y := startY - 1; y <= startY+1; y++ {
		x := startX
		if isValid, nextY, nextX := check(lines, square_width, square_height, x, y, startX, startY); isValid {
			p := Point{X: x, Y: y, nextX: nextX, nextY: nextY}
			points[p] = true
			visitedNodes[y][x] = true
			fmt.Printf("Point %v (%d,%d), next(%d,%d)\n", string(lines[p.Y][p.X]), p.Y, p.X, p.nextY, p.nextX)
		}
	}
	for x := startX - 1; x <= startX+1; x++ {
		y := startY
		if isValid, nextY, nextX := check(lines, square_width, square_height, x, y, startX, startY); isValid {
			p := Point{X: x, Y: y, nextX: nextX, nextY: nextY}
			points[p] = true
			visitedNodes[y][x] = true
			fmt.Printf("Point %v (%d,%d), next(%d,%d)\n", string(lines[p.Y][p.X]), p.Y, p.X, p.nextY, p.nextX)
		}
	}
	var temp map[Point]bool
	steps := 1
	for len(points) > 1 {
		fmt.Println("step ", steps)
		temp = make(map[Point]bool)
		for p, _ := range points {
			if p.nextX >= square_width || p.nextY >= square_height || p.nextX < 0 || p.nextY < 0 {
				continue
			}

			startX := p.X
			startY := p.Y
			x := p.nextX
			y := p.nextY
			if isValid, nextY, nextX := check(lines, square_width, square_height, x, y, startX, startY); isValid {
				if !visitedNodes[nextY][nextX] {

					q := Point{X: x, Y: y, nextX: nextX, nextY: nextY}
					fmt.Printf("Point %v (%d,%d), next(%d,%d)\n", string(lines[q.Y][q.X]), q.Y, q.X, q.nextY, q.nextX)
					temp[q] = true
					visitedNodes[y][x] = true
				}
			} else {
				fmt.Println("not valid ", string(lines[y][x]))
			}

			// temp[p] = true
		}
		points = temp
		fmt.Printf("\n NEXT \n")
		steps++
	}

	fmt.Println(steps)
}

func check(lines []string, square_width, square_height, x, y, startX, startY int) (bool, int, int) {

	if y < 0 || y >= square_height || x < 0 || x >= square_width {
		return false, 0, 0
	}
	point := lines[y][x]
	// fmt.Println("Point: ", y, x, string(lines[y][x]))
	if point == '.' || point == 'S' {
		return false, 0, 0
	}
	switch lines[y][x] {
	case '-':
		if startX != x {
			return true, y, x + (x - startX)
		}
	case '|':
		if startY != y {
			return true, y + (y - startY), x
		}
	case 'L':
		if y-startY == 1 {
			return true, y, x + 1
		}
		if x-startX == -1 {
			return true, y - 1, x
		}
	case 'J':
		if startY-y == -1 {
			return true, y, x - 1
		}
		if x-startX == 1 {
			return true, y - 1, x
		}
	case '7':
		if x-startX == 1 {
			return true, y + 1, x
		}
		if y-startY == -1 {
			return true, y, x - 1
		}

	case 'F':
		if startY-y == 1 {
			return true, y, x + 1
		}
		if x-startX == -1 {
			return true, y + 1, x
		}
	}
	return false, 0, 0
}
