package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func readDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	deck := strings.Split(string(bs), ",")
	return deck, err
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) writeToFile(filename string) error {
	data := []byte(d.toString())
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
