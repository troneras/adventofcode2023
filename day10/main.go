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
	body, _ := os.ReadFile("test.txt")
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
	var q Point
	pipe := []Point{}

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

					q = Point{X: x, Y: y, nextX: nextX, nextY: nextY}
					fmt.Printf("Point %v (%d,%d), next(%d,%d)\n", string(lines[q.Y][q.X]), q.Y, q.X, q.nextY, q.nextX)
					temp[q] = true
					visitedNodes[y][x] = true
					pipe = append(pipe, q)
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

	visitedNodes[q.nextY][q.nextX] = true

	fmt.Println(steps)
	c := ""
	visited := 0
	not_visited := 0
	for i := 0; i < square_height; i++ {
		for j := 0; j < square_width; j++ {
			if visitedNodes[i][j] {
				c = "X"
				visited++
			} else {
				c = string(lines[i][j])
				if c != "." {
					// c = "0"
				}
				not_visited++
			}
			fmt.Printf("%v", c)
		}
		fmt.Printf("\n")
	}

	internal := map[Point]bool{}
	dir := false
	for _, p := range pipe {
		// direction
		vX, vY := p.nextX-p.X, p.nextY-p.Y

		// see if has to the right any point that is not visited
		if vX == 0 {
			if dir {
				if vY > 0 && p.X-1 >= 0 {
					// arrow pointing down right is y, x - 1
					if visitedNodes[p.Y][p.X-1] == false {
						p := Point{X: p.X - 1, Y: p.Y}
						internal[p] = true
					}
				}
				if vY < 0 && p.X+1 < square_width {
					// arrow pointing up, right is y, x + 1
					if visitedNodes[p.Y][p.X+1] == false {
						p := Point{X: p.X + 1, Y: p.Y}
						internal[p] = true
					}
				}
			} else {

				if vY < 0 && p.X-1 >= 0 {
					// arrow pointing down right is y, x - 1
					if visitedNodes[p.Y][p.X-1] == false {
						p := Point{X: p.X - 1, Y: p.Y}
						internal[p] = true
					}
				}
				if vY > 0 && p.X+1 < square_width {
					// arrow pointing up, right is y, x + 1
					if visitedNodes[p.Y][p.X+1] == false {
						p := Point{X: p.X + 1, Y: p.Y}
						internal[p] = true
					}
				}
			}
		} else { // vY == 0
			if dir {

				if vX > 0 && p.Y+1 < square_height {
					// arrow pointing right, righ is y + 1, x
					if visitedNodes[p.Y+1][p.X] == false {
						p := Point{X: p.X, Y: p.Y + 1}
						internal[p] = true
					}
				}
				if vX < 0 && p.Y-1 >= 0 {
					// arrow pointing left, right is y - 1, x
					if visitedNodes[p.Y-1][p.X] == false {
						p := Point{X: p.X, Y: p.Y - 1}
						internal[p] = true
					}
				}
			} else {
				if vX < 0 && p.Y+1 < square_height {
					// arrow pointing right, righ is y + 1, x
					if visitedNodes[p.Y+1][p.X] == false {
						p := Point{X: p.X, Y: p.Y + 1}
						internal[p] = true
					}
				}
				if vX > 0 && p.Y-1 >= 0 {
					// arrow pointing left, right is y - 1, x
					if visitedNodes[p.Y-1][p.X] == false {
						p := Point{X: p.X, Y: p.Y - 1}
						internal[p] = true
					}
				}
			}
		}

	}

	for i, _ := range internal {
		// fmt.Println(i)
		fmt.Printf("Point %v (%d,%d)\n", string(lines[i.Y][i.X]), i.Y, i.X)
	}
	fmt.Println("total: ", visited+not_visited, "visited: ", visited, "internal: ", len(internal))

	// out := []Point{}
	// in := []Point{}

	// for not_visited > 0 {
	// 	// find one not visited
	// 	var p Point
	// 	for i := 0; i < square_height; i++ {
	// 		for j := 0; j < square_width; j++ {
	// 			if visitedNodes[i][j] == false {
	// 				p := Point{X: j, Y: i, nextX: 0, nextY: 0}
	// 				break
	// 			}
	// 		}
	// 		// check if it's an external point
	// 		if p.X == 0 || p.Y == 0 {
	// 			out = append(out, Point{X: p.X, Y: p.Y})

	// 			for y := p.Y - 1; y <= p.Y; y++ {
	// 				for x := p.X - 1; x <= p.X+1; x++ {
	// 					if x < 0 || y < 0 || x >= square_width || y >= square_height || visitedNodes[y][x] {
	// 						continue
	// 					}
	// 					out = append(out, Point{X: x, Y: y})
	// 				}
	// 			}
	// 		}
	// 	}
	// }

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
