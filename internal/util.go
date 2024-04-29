package internal

import (
	"fmt"
	"strconv"
)

const (
	StartingChessPieces = 7
	GridLength = 8
)

func GetNewBoard() *Board {
	return &Board{make(map[ChessPiece]struct{})}
}

func GetNewPlayer(pname string) *Player {
	fmt.Printf("Generated player %s\n", pname)
	return &Player{pname, make([]ChessPiece, StartingChessPieces)}
}

func GetNewGame(p1 *Player, p2 *Player, currentPlayer *Player, grid *Board) *Game {
	return &Game{
		Player1: p1,
		Player2: p2,
		CurrentPlayer: currentPlayer,
		Grid: grid,
	}
}

func GetWinner(game *Game) *Player {
	if game.Winner == game.Player1 || game.Winner == game.Player2 {
		return game.Winner
	}
	return nil
}

func PrintUnit(unit string, board *Board) {
	for piece := range board.BoardState {
		if piece.GridPosition == unit {
			fmt.Print(piece.PieceType)
			return
		}
	}
	print(" ")
}

func PrintBoard(board *Board) {
	// Print First Row
	for row := range GridLength {
		if row == 4 || row == 5 {
			if row == 4 {
				fmt.Print("-  ")
				continue
			}
			fmt.Print("   ")
			continue
		}
		fmt.Print("---")
	}

	fmt.Println("-")

	// Handle printing pieces
	for _, col := range []int{4,3,2,1,-1,-2,14,13} {
		if col < 0 {
			if col == -1 {
				fmt.Print("|  ")
				continue
			}
			fmt.Print("   ")
			continue
		}
		fmt.Print("|")
		PrintUnit("A" + strconv.Itoa(col), board)
		fmt.Print(" ")
	}

	fmt.Println("|")

	// Second Row
	for range GridLength {
		fmt.Print("---")
	}
	fmt.Println("-")

	for _, col := range []int{1,2,3,4,5,6,7,8} {
		fmt.Print("|")
		PrintUnit("P" + strconv.Itoa(col), board)
		fmt.Print(" ")
	}
	fmt.Println("|")

	// Third Row
	for range GridLength {
		fmt.Print("---")
	}

	fmt.Println("-")

	for _, col := range []int{4,3,2,1,-1,-2,14,13} {
		if col < 0 {
			if col == -1 {
				fmt.Print("|  ")
				continue
			}
			fmt.Print("   ")
			continue
		}
		fmt.Print("|")
		PrintUnit("B" + strconv.Itoa(col), board)
		fmt.Print(" ")
	}

	fmt.Println("|")

	for row := range GridLength {
		if row == 4 || row == 5 {
			if row == 4 {
				fmt.Print("-  ")
				continue
			}
			fmt.Print("   ")
			continue
		}
		fmt.Print("---")
	}
	fmt.Println("-")
}

func (p *Game) DeepCopy() *Player {
	
}