package internal

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func Init_Game() {
	fmt.Println("Initilizing the game...")

	p1, p2 := GetNewPlayer("player1", true), GetNewPlayer("player2", false)
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
	fmt.Printf("Your total moves is %d\n", move)
	return move
}

func randomColor() DiceColor {
    // Generate a random number 0 or 1 and return corresponding color
    if rand.Intn(2) == 0 {
		fmt.Println("You got Black")
        return Black
    }
	fmt.Println("You got White")
    return White
}

func switchCurrentPlayer(game *Game) {
	if game.CurrentPlayer == game.Player1 {
		game.CurrentPlayer = game.Player2
	} else {
		game.CurrentPlayer = game.Player1
	}
}

func makeMove(game *Game, move uint) {
	var available_moves []Game = make([]Game, 0)

	for _, piece := range game.CurrentPlayer.Pieces {
		if game.CurrentPlayer == game.Player1 {
			if piece.State == NotInPlay {
				if _, ok := game.Grid.BoardState["A0" + strconv.Itoa(int(move))]; !ok {

				}
			}
		} else {
			 
		}
	}
}

func Play(game *Game) {
	fmt.Println("Game started")
	var winner *Player = nil

	PrintBoard(game.Grid)
	for winner != nil {
		fmt.Printf("It's %s's turn...\n", game.CurrentPlayer.PlayerName)
		fmt.Printf("%s, please throw your dices...\n", game.CurrentPlayer.PlayerName)

		reader := bufio.NewReader(os.Stdin)
		// The newline character triggered by pressing "Enter"
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while waiting for input:", err)
			return
		}

		move := throwDices() 
		if move == 0 {
			fmt.Println("Rolled 0 LOL. Your turn is skipped...")
			switchCurrentPlayer()
			continue
		}
		makeMove(game.CurrentPlayer, move)

		winner = GetWinner(game)
	}
	fmt.Printf("Game over, the winner is ")

}