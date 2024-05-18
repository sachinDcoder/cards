package main

import "fmt"

func main() {
	cards := deck{"Six of Diamonds", newCards()}
	cards = append(cards, "Five of Ace")
	//fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCards() string {
	return "Five of Diamonds"
}
