package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	fiveOfAKind  = 7
	fourOfAKind  = 6
	fullHouse    = 5
	threeOfAKind = 4
	twoPairs     = 3
	onePair      = 2
	highCard     = 1
)

type Hand struct {
	cards []byte
	bid   int
	rank  int
	kind  int
}

type ByKind []Hand

func (a ByKind) Len() int      { return len(a) }
func (a ByKind) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByKind) Less(i, j int) bool {
	if a[i].kind == a[j].kind {

		for k := 0; k < 5; k++ {
			lt, eq := compareCard(a[i].cards[k], a[j].cards[k])
			if eq {
				continue
			}
			return lt
		}

	}

	return a[i].kind < a[j].kind

}

func main() {
	body, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))
	hands := []Hand{}

	lines := bytes.Split(body, []byte("\n"))
	for _, line := range lines {
		hand := getHand(line)

		hands = append(hands, hand)
	}

	sort.Sort(ByKind(hands))

	total := 0
	for i := 0; i < len(hands); i++ {
		hand := hands[i]
		fmt.Println("cartas:", string(hand.cards), "apuesta: ", hand.bid, "TIPO: ", hand.kind)
		total += hand.bid * (i + 1)
	}

	fmt.Println(total)

}

func getHand(line []byte) Hand {
	hand := Hand{}
	hand.cards = line[0:5]
	hand.bid, _ = strconv.Atoi(string(line[6:]))
	hand.rank = 0

	num_jokers := bytes.Count(hand.cards, []byte{'J'})
	if num_jokers == 5 {
		hand.kind = fiveOfAKind
		return hand
	}
	cards := bytes.ReplaceAll(hand.cards, []byte{'J'}, []byte{})

	_, card := getKind(cards)
	for i := 0; i < num_jokers; i++ {
		cards = append(cards, card)
	}

	hand.kind, _ = getKind(cards)

	return hand
}

func getKind(cards []byte) (int, byte) {

	for _, card := range cards {
		if bytes.Count(cards, []byte{card}) == 5 {
			return fiveOfAKind, card
		}
		if bytes.Count(cards, []byte{card}) == 4 {
			return fourOfAKind, card
		}
		if bytes.Count(cards, []byte{card}) == 3 {
			cardsCopy := cards
			cardsCopy = bytes.ReplaceAll(cardsCopy, []byte{card}, []byte{})
			if len(cardsCopy) == 0 {
				return threeOfAKind, card
			}

			card2 := cardsCopy[0]
			count := bytes.Count(cardsCopy, []byte{card2})
			if count == 2 {
				return fullHouse, card
			}
			return threeOfAKind, card
		}
		if bytes.Count(cards, []byte{card}) == 2 {
			cardsCopy := cards
			cardsCopy = bytes.ReplaceAll(cardsCopy, []byte{card}, []byte{})
			for len(cardsCopy) > 1 {
				card2 := cardsCopy[0]
				count := bytes.Count(cardsCopy, []byte{card2})
				if count == 2 {
					return twoPairs, card
				}
				if count == 3 {
					return fullHouse, card2
				}
				cardsCopy = cardsCopy[1:]
			}

			return onePair, card
		}
	}
	return highCard, cards[0]
}

func compareCard(card1 byte, card2 byte) (bool, bool) {
	cv1 := getCardValue(card1)
	cv2 := getCardValue(card2)

	if cv1 == cv2 {
		return false, true
	}

	return getCardValue(card1) < getCardValue(card2), false
}

func getCardValue(card byte) int {
	all := "J23456789TQKA"
	return strings.Index(all, string(card)) + 2
}
