package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Range struct {
	a int
	b int
	c int
}

type SeedRange struct {
	start int
	end   int
}

func main() {
	body, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	before, after, _ := bytes.Cut(body, []byte("\n\n"))

	parts := bytes.Split(after, []byte("\n\n"))

	// split seeds by ":" and take the first value
	before = bytes.Split(before, []byte(": "))[1]

	// split seeds by " " as int array
	inputs, _ := byteToIntArray(before)
	seeds := []SeedRange{}
	for i := 0; i < len(inputs); i += 2 {
		seeds = append(seeds, SeedRange{start: inputs[i], end: inputs[i] + inputs[i+1]})
	}

	// fmt.Println("seeds: ", seeds)

	found := false

	for _, part := range parts {
		ranges := []Range{}
		for _, line := range bytes.Split(part, []byte("\n"))[1:] {
			r, _ := byteToIntArray(line)
			ranges = append(ranges, Range{a: r[0], b: r[1], c: r[2]})
		}

		new := []SeedRange{}

		for len(seeds) > 0 {
			// pop first element
			seed := seeds[0]
			seeds = seeds[1:]
			// fmt.Println("seed: ", seed)
			// fmt.Println("ranges: ", ranges)

			found = false

			for _, r := range ranges {
				// fmt.Println("range: ", r)
				var overlap_start int
				if seed.start < r.b {
					overlap_start = r.b
				} else {
					overlap_start = seed.start
				}
				var overlap_end int
				if seed.end < r.b+r.c {
					overlap_end = seed.end
				} else {
					overlap_end = r.b + r.c
				}

				// fmt.Println("overlap_start: ", overlap_start)
				// fmt.Println("overlap_end: ", overlap_end)

				if overlap_start < overlap_end {
					new = append(new, SeedRange{start: overlap_start - r.b + r.a, end: overlap_end - r.b + r.a})
					found = true
					if overlap_start > seed.start {
						seeds = append(seeds, SeedRange{start: seed.start, end: overlap_start})
					}
					if seed.end > overlap_end {
						seeds = append(seeds, SeedRange{start: overlap_end, end: seed.end})
					}
					break
				}
			}

			if !found {
				new = append(new, seed)
			}

		}
		seeds = new
	}

	if len(seeds) == 0 {
		fmt.Println("No elements in the slice")
		return
	}

	// fmt.Println("seeds: ", seeds)

	s := seeds[0]
	min := s.start
	for _, value := range seeds[1:] {
		if value.start < min {
			min = value.start
		}
	}

	fmt.Println(min)
}

func byteToIntArray(input []byte) ([]int, error) {
	parts := bytes.Split(input, []byte(" "))
	result := []int{}
	for _, part := range parts {
		seed, err := strconv.Atoi(string(part))
		if err != nil {
			return nil, err
		}
		result = append(result, seed)
	}
	return result, nil
}
