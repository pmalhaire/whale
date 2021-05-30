package whale

type Card int

const (
	UnknownCard Card = iota
	Water
	Wave
	DoubleWave
	Pirat
	Ghost
	Round
	Piranha
)

// String converts enum to string for printing
func (c Card) String() string {
	return [...]string{
		"UnknownCard",
		"||",
		"~",
		"~~",
		"Pirat",
		"Ghost",
		"Round",
		"Piranha",
	}[c]
}

var cardList = [...]Card{
	Water,
	Wave,
	DoubleWave,
	Pirat,
	Ghost,
	Round,
	Piranha,
}

// CardList reuturns the listof available cards
func CardList() []Card {
	return cardList[:]
}
