package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	body, _ := os.ReadFile("input.txt")

	input := bytes.Split(body, []byte("\n"))

	total := 0

	for _, line := range input {
		sequence := []int{}
		numbs := bytes.Split(line, []byte(" "))
		for _, numb := range numbs {
			n, _ := strconv.Atoi(string(numb))
			sequence = append(sequence, n)
		}

		difs := [][]int{}
		difs = append(difs, sequence)

		for {
			dif := difs[len(difs)-1]
			dif = getDif(dif)
			difs = append(difs, dif)
			if allZeroes(dif) {
				break
			}
		}

		next := 0
		for i := len(difs) - 1; i >= 0; i-- {
			dif := difs[i]
			if i != len(difs)-1 {
				next = dif[len(dif)-1] + next
			}
			dif = append(dif, next)
			difs[i] = dif
		}

		total += difs[0][len(difs[0])-1]
	}
	fmt.Println(total)
}

func getDif(sequence []int) []int {
	dif := []int{}
	for i := 0; i < len(sequence)-1; i++ {
		dif = append(dif, sequence[i+1]-sequence[i])
	}
	return dif
}

func sumSequence(sequence []int) int {
	sum := 0
	for _, numb := range sequence {
		sum += numb
	}
	return sum
}

func allZeroes(sequence []int) bool {
	for _, numb := range sequence {
		if numb != 0 {
			return false
		}
	}
	return true
}
