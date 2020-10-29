package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingDeck := deal(cards, 4)
	hand.print()
	remainingDeck.print()
	cards.writeToFile("all_cards.txt")

	cards, err := readDeckFromFile("all_cards.txt")
	if err == nil {
		cards.toString()
	} else {
		fmt.Println("ERROR READING FILE")
	}
	cards.shuffle()
	fmt.Println(cards.toString())
	cards.writeToFile("shuffled.txt")

	cards, err = readDeckFromFile("all_cards2.txt")
	if err == nil {
		cards.toString()
	} else {
		fmt.Println("ERROR READING FILE 2")
	}
	fmt.Println("End")
}
