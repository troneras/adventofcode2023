package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FromToMap struct {
	source_start      int
	destination_start int
	delta             int
}

func main() {
	maps := make(map[string]map[string][]FromToMap)

	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	filescanner := bufio.NewScanner(file)

	filescanner.Split(bufio.ScanLines)

	seeds := []int{}

	from, to := "", ""

	for filescanner.Scan() {
		// does the line starts with "seeds:" ? if yes, print the line
		line := filescanner.Text()
		if len(line) > 6 && line[0:6] == "seeds:" {
			// split the line by space
			splited_line := strings.Split(line, " ")
			// convert the string to int
			for _, seed := range splited_line[1:] {
				number, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, number)
			}

			fmt.Println("seeds: ", seeds)
		}

		// we have the seeds, now we need to create the maps
		// there are lines starting with "somethig-to-something map:" that we need to parse
		if len(line) > 6 && line[len(line)-4:] == "map:" {
			from = line[0:strings.Index(line, "-")]
			to = line[strings.Index(line, "-")+4 : len(line)-5]
			// fmt.Println("from: ", from, " to: ", to)
			// create the map
			maps[from] = make(map[string][]FromToMap)
			// create the map for the destination
			maps[from][to] = []FromToMap{}
		}

		// if the line starts with a digit, we need to parse it
		if len(line) > 0 && line[0] >= '0' && line[0] <= '9' {
			// split the line by space
			splited_line := strings.Split(line, " ")
			if len(splited_line) != 3 {
				continue
			}
			source_start, err := strconv.Atoi(splited_line[1])
			if err != nil {
				log.Fatal(err)
			}

			destination_start, err := strconv.Atoi(splited_line[0])
			if err != nil {
				log.Fatal(err)
			}

			delta, err := strconv.Atoi(splited_line[2])
			if err != nil {
				log.Fatal(err)
			}

			maps[from][to] = append(maps[from][to], FromToMap{
				source_start:      source_start,
				destination_start: destination_start,
				delta:             delta,
			})
		}

	}

	seeds2 := []struct {
		start int
		end   int
	}{}
	for k, v := range seeds {
		if k%2 == 1 {
			continue
		}

		if k+1 < len(seeds) {
			seeds2 = append(seeds2, struct {
				start int
				end   int
			}{
				start: v,
				end:   v + seeds[k+1],
			})
		}
	}
	fmt.Print("seeds2: ", seeds2, "\n")

	to_resource := 0
	lowest := 9999999999
	for _, seed := range seeds2 {
		fmt.Print()
		for i := seed.start; i < seed.end; i++ {
			from = "seed"
			to = "soil"
			to_resource = i
			// while to != "location"
			for {
				//prev_to_resource = to_resource
				to_resource = getToResource(from, to, to_resource, maps)
				//fmt.Println("from: ", from, " to: ", to, " seed: ", prev_to_resource, " to_resource: ", to_resource)
				from = to
				if to == "location" {
					break
				}
				var firstKey string
				for key := range maps[from] {
					firstKey = key
					break
				}
				to = firstKey

			}

			if to_resource < lowest {
				lowest = to_resource
			}
		}
	}

	fmt.Println("lowest: ", lowest)
}

func getToResource(from string, to string, seed int, maps map[string]map[string][]FromToMap) int {
	// get the map
	m := maps[from][to]

	// fmt.Println("from: ", from, " to: ", to, " seed: ", seed, " map: ", m)
	// get the destination
	for _, v := range m {
		if seed >= v.source_start && seed <= (v.source_start+v.delta) {
			return v.destination_start + (seed - v.source_start)
		}
	}
	return seed
}
