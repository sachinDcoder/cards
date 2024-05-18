package main

func main() {
	cards := deck{"Six of Diamonds", newCards()}
	cards = append(cards, "Five of Ace")

	cards.print()
}

func newCards() string {
	return "Five of Diamonds"
}
