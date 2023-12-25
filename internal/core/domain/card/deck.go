package card

import "math/rand"

// Deck represents a deck of cards
type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	d := &Deck{}
	d.Reset()
	d.Shuffles()
	return d
}

func (d *Deck) Draw() *Card {
	if d.cards == nil || len(d.cards) == 0 {
		return nil
	}

	card := d.cards[0]
	d.cards = d.cards[1:]
	return &card
}

func (d *Deck) Shuffles() {
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Burn() {
	if d.cards == nil || len(d.cards) == 0 {
		return
	}

	d.cards = d.cards[1:]
}

func (d *Deck) Reset() {
	d.cards = make([]Card, 0)
	allSuits := []SuitValue{Spade, Heart, Diamond, Club}
	allValue := []CardValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, Knight, Ace}
	for _, suit := range allSuits {
		for _, value := range allValue {
			c := NewCard(value, suit)
			d.cards = append(d.cards, *c)
		}
	}
}
