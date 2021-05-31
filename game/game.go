package game

import "math/rand"

// Game holds data related to a game
type Game struct {
	// deck of cards
	Deck *Deck
	// slice holding players
	Players []Player
	// number of current round
	Round int
	// index of playing player
	playerIndex int
}

// State data available to all players
type State struct {
	// info available to all players
	PlayersInfo []PlayerInfo
	// number of current round
	Round int
	// index of playing player
	PlayerIndex int
}

// number of cards in the initial hand
const intialCardCount = 3

// Table creates nbPlayers players with random bonuses
func Table(nbPlayers int) []Player {
	if nbPlayers < 2 {
		panic("invalid player count must be more than 2")
	}
	if nbPlayers > 4 {
		panic("max player is 4")
	}
	bonuses := BonusList()
	rand.Shuffle(len(bonuses), func(i, j int) {
		bonuses[i], bonuses[j] = bonuses[j], bonuses[i]
	})
	players := make([]Player, nbPlayers)
	for i := 0; i < nbPlayers; i++ {
		players[i] = *NewPlayer(bonuses[i])
	}
	return players
}

// NewGame creates a new game with nbPlayers players.
func NewGame(nbPlayers int) *Game {
	deck := NewDeck()
	deck.Shuffle()
	players := Table(nbPlayers)

	// draw cards
	for i := 0; i < nbPlayers*intialCardCount; i++ {
		players[i%nbPlayers].AddCard(deck.Pick())
	}
	return &Game{
		Deck:        deck,
		Players:     players,
		Round:       0,
		playerIndex: 0,
	}
}

// PlayerCount is the number of players in the game
func (g *Game) PlayerCount() int {
	return len(g.Players)
}

// CurentPlayer is player to play
func (g *Game) CurentPlayer() *Player {
	return &g.Players[g.playerIndex]
}

// CurentPlayerIndex index is the index of player to play
func (g *Game) CurentPlayerIndex() int {
	return g.playerIndex
}

// NextPlayer returns the next player to play
func (g *Game) NextPlayer() *Player {
	// end game
	if g.Players[g.playerIndex].IsWinner() {
		return nil
	}

	g.playerIndex++

	if g.playerIndex == len(g.Players) {
		g.Round++
		g.playerIndex = 0
	}
	return &g.Players[g.playerIndex]
}

// State return information about the game available to all players
func (g *Game) State() State {
	playersInfo := []PlayerInfo{}
	for _, p := range g.Players {
		playersInfo = append(playersInfo, PlayerInfo{
			CardCount:   len(p.Cards),
			Water:       p.Water,
			BonusType:   p.BonusType,
			BonusPlayed: p.BonusPlayed,
		})
	}
	return State{
		PlayersInfo: playersInfo,
		Round:       g.Round,
		PlayerIndex: g.playerIndex,
	}
}