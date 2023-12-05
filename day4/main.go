package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type card struct {
	num      int
	winning  []int
	existing []int
	hits     map[int]int
	value    int
}

func main() {
	file, err := os.Open("input.txt")
	cards := make(map[int]struct {
		hits        int
		card_copies []card
	})

	if err != nil {
		panic(err)
	}

	defer file.Close()

	filescanner := bufio.NewScanner(file)

	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		line := filescanner.Text()
		fmt.Println(line)

		// find the card number
		current := 0
		card_num := 0
		for i := 0; i < len(line); i++ {
			if line[i] == ':' {
				current = i + 1
				break
			}
			if unicode.IsDigit(rune(line[i])) {
				card_num = card_num*10 + int(line[i]-'0')
			}
		}
		//fmt.Println("Card number: ", card_num)

		// find the winning numbers, numbers are separated by spaces
		winning := []int{}
		for i := current; i < len(line); i++ {
			if line[i] == '|' {
				current = i + 1
				break
			}
			if unicode.IsDigit(rune(line[i])) {
				partial := string(line[i])
				i++
				for unicode.IsDigit(rune(line[i])) {
					partial += string(line[i])
					i++
				}
				number, err := strconv.Atoi(partial)
				if err != nil {
					panic(err)
				}
				winning = append(winning, number)
			}
		}

		//fmt.Println("Winning numbers: ", winning)

		// find the existing numbers, numbers are separated by spaces
		existing := []int{}
		for i := current; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				partial := string(line[i])
				i++
				for i < len(line) && unicode.IsDigit(rune(line[i])) {
					partial += string(line[i])
					i++
				}
				number, err := strconv.Atoi(partial)
				if err != nil {
					panic(err)
				}
				existing = append(existing, number)
			}
			current = i
		}

		// if empty, add the card to the map
		if _, ok := cards[card_num]; !ok {
			cards[card_num] = struct {
				hits        int
				card_copies []card
			}{
				hits: 0,
				card_copies: []card{
					{num: card_num, winning: winning, existing: existing},
				},
			}

		} else {
			// otherwise, append the card to the existing card
			copies := cards[card_num].card_copies
			copies = append(copies, card{num: card_num, winning: winning, existing: existing})
			element := cards[card_num]
			element.card_copies = copies
			cards[card_num] = element
		}

	}

	// print how many cards there are
	//fmt.Println("There are ", len(cards), " cards")

	// calculate the number of hits for each card
	for i := 1; i <= len(cards); i++ {
		hits := make(map[int]int)
		for j := 0; len(cards[i].card_copies) > 0 && j < len(cards[i].card_copies[0].winning); j++ {
			for k := 0; k < len(cards[i].card_copies[0].existing); k++ {
				if cards[i].card_copies[0].winning[j] == cards[i].card_copies[0].existing[k] {
					hits[cards[i].card_copies[0].winning[j]]++
				}
			}
		}
		card := cards[i]
		card.hits = len(hits)
		cards[i] = card

		//fmt.Printf("Card %d has %d hits\n", cards[i].card_copies[0].num, cards[i].hits)
	}

	total := 0

	for i := 1; i <= len(cards); i++ {
		hits := cards[i].hits
		fmt.Printf("Card %d has %d hits\n", cards[i].card_copies[0].num, hits)
		if hits == 0 {
			continue
		}
		for k := 0; k < len(cards[i].card_copies); k++ {
			hits = cards[i].hits
			//fmt.Printf("Card %d has %d copies\n", cards[i].card_copies[k].num, len(cards[i].card_copies))
			for j := i + 1; j <= len(cards); j++ {
				copies := cards[j].card_copies
				// append another card to the card copies that is a copy of the first card
				copies = append(copies, cards[i].card_copies[0])
				card := cards[j]
				card.card_copies = copies
				cards[j] = card
				//fmt.Printf("Won a copy of card %d\n", cards[j].card_copies[0].num)
				hits--
				if hits == 0 {
					break
				}
			}
		}
		fmt.Printf("There are %d copies of card %d which has %d hits\n", len(cards[i].card_copies), cards[i].card_copies[0].num, cards[i].hits)

	}

	for i := 1; i <= len(cards); i++ {
		fmt.Printf("Card %d has %d copies\n", cards[i].card_copies[0].num, len(cards[i].card_copies))
		total += len(cards[i].card_copies)
	}

	fmt.Println("Total number of cards: ", total)

}
