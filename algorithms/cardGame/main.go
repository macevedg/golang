package main

import (
	"fmt"
	"math/rand"
)

// input
// un deck de cartas de 52 cartas
// limite de jugadores: 2-10 jugadores
// 1. cada jugador debe tener su set de cartas
// 2. siempre se reparten la misma candidad de cartas
// 3. condicion: repartir hasta donde sea posible
// 4. habria una candidad de comparaciones igual al numero de cartas por jugador
// 4. repartir todas las cartas de modo que todos tengan el mismo numero de cartas
// 4.b print el score de todos los jugadores
// 5 salida: el ganador, el score

var deckSize int

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	Game(10, 52)
}

type Player struct {
	ID     string
	Cards  []int
	Points int
}

func Game(playerAmount int, actualDeckSize int) {

	deckSize = actualDeckSize
	//1. crear jugadores
	deck := initializeDeck()
	//2. asignar cartas
	playerList, numberOfCards := initializePlayers(deck, playerAmount)

	// perform game
	playerList = handleGame(playerList, numberOfCards)
	// print and return

	for _, player := range playerList {
		fmt.Printf("Name: %s, points: %d\n", player.ID, player.Points)
	}
}

func handleGame(playerList []Player, numberOfCards int) []Player {

	for i := 0; i < numberOfCards; i++ {
		playerList = findRoundWinner(playerList, i)
	}

	return playerList
}

func findRoundWinner(playerList []Player, cardIndex int) []Player {

	winnerIndex := -1
	greaterCard := -1
	for index, p := range playerList {
		card := p.Cards[cardIndex]
		fmt.Printf("CARD %d\n", card)

		if greaterCard < card {
			greaterCard = card
			winnerIndex = index
		}
	}

	playerList[winnerIndex].Points += 1

	return playerList
}

// deck initialization
func initializeDeck() map[int]int {
	newDeck := make(map[int]int, deckSize)
	for i := 0; i < deckSize; i++ {
		newDeck[i] = i + 1
	}

	return newDeck
}

// players initialization
func initializePlayers(deck map[int]int, playerAmount int) ([]Player, int) {

	playerList := make([]Player, 0)

	//initialize players
	for i := 0; i < playerAmount; i++ {
		playerList = append(playerList, Player{
			ID:     fmt.Sprintf("PLayer %d", (i + 1)),
			Cards:  make([]int, 0),
			Points: 0,
		})
	}

	numberOfCards := 5
	deckCounter := 0

	fmt.Printf("number of cards %d, deck size: %d\n", numberOfCards, deckSize)

	for index, p := range playerList {

		for i := 0; i < numberOfCards; i++ {
			card := getCard(deck)
			fmt.Printf("Card is: %d\n", card)
			p.Cards = append(p.Cards, card)
			playerList[index] = p
			deckCounter++
		}
	}

	fmt.Printf("PLAYERS: %v\n", playerList)

	return playerList, numberOfCards
}

func getCard(deck map[int]int) int {

	gotCard := false

	var card int
	for !gotCard {
		index := rand.Intn(52)

		value, _ := deck[index]

		if value != -1 {
			deck[index] = -1
			gotCard = true
			card = index
			break
		}

	}

	return card
}
