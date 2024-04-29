package internal

import (
	"fmt"
	"math/rand"
)

func Init_Game() {
	fmt.Println("Initilizing the game...")

	p1, p2 := GetNewPlayer("player1"), GetNewPlayer("player2")
	fmt.Println("P1's team is '@'")
	fmt.Println("P2's team is '*'")
	playerGoFirst := goFirst(p1, p2)
	grid := GetNewBoard()
	game := GetNewGame(p1, p2, playerGoFirst, grid)

	Play(game)
}

func goFirst(p1 *Player, p2 *Player) *Player {
	fmt.Println("Deciding who will go first...")
	num := rand.Intn(100)

	if num < 50 {
		fmt.Printf("Player1 %s will go first!\n", p1.PlayerName)
		return p1
	} else {
		fmt.Printf("%s will go first!\n", p2.PlayerName)
		return p2
	}
}

func throwDices() uint {
	var move uint = 0
	for i := 0; i < 4; i++  {
		if randomColor() == White {
			move += 1
		}
	}
	return move
}

func randomColor() DiceColor {
    // Generate a random number 0 or 1 and return corresponding color
    if rand.Intn(2) == 0 {
        return Black
    }
    return White
}

func Play(game *Game) {
	fmt.Println("Game started")
	var winner *Player = nil

	PrintBoard(game.Grid)
	for winner != nil {
		fmt.Printf("It's %s's turn...\n", game.CurrentPlayer.PlayerName)
		winner = GetWinner(game)
	}

}