package card

type CardValue string

const (
	Two    CardValue = "cv2"
	Three            = "cv3"
	Four             = "cv4"
	Five             = "cv5"
	Six              = "cv6"
	Seven            = "cv7"
	Eight            = "cv8"
	Nine             = "cv9"
	Ten              = "cv10"
	Jack             = "cvJ"
	Queen            = "cvQ"
	Knight           = "cvK"
	Ace              = "cvA"
)

// String returns the string representation of CardValue
func (value CardValue) String() string {
	return string(value[2:])
}

type SuitValue string

const (
	Spade   SuitValue = "cs♠"
	Heart             = "cs♥"
	Diamond           = "cs♦"
	Club              = "cs♣"
)

// String returns the string representation of SuitValue
func (s SuitValue) String() string {
	return string(s[2:])
}

// Card represents a playing card
type Card struct {
	Value CardValue
	Suit  SuitValue
}

// New creates a new card with the given value and suit
func NewCard(v CardValue, s SuitValue) *Card {
	return &Card{Value: v, Suit: s}
}

func NewCardFromString(card string) *Card {
	if card[3] == 'c' {
		return &Card{Value: card[3:], Suit: card[:3]}
	} else {
		return &Card{Value: card[4:], Suit: card[:4]}
	}
}

// GetSuit returns the suit of a card
func GetSuit(card Card) SuitValue {
	return card.Suit
}

// GetValue returns the value of a card
func GetValue(card Card) CardValue {
	return card.Value
}

// ToString returns the string representation of a card
func (card Card) ToString() string {
	return card.Value.String() + card.Suit.String()
}
