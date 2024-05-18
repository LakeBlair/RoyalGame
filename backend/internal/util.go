package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	StartingChessPieces = 7
	GridLength = 8
)

func GetNewBoard() *Board {
	return &Board{make(map[string]*ChessPiece)}
}

func GetNewPlayer(pname string, isPlayer1 bool) *Player {
	fmt.Printf("Generated player %s\n", pname)
	pieces := make([]*ChessPiece, StartingChessPieces)

	pieceType := rune('*') // Default to Player2
	party := rune('B')

    if isPlayer1 {
        pieceType = rune('@')
		party = rune('A')
    }

	for i := range pieces {
        pieces[i] = &ChessPiece{
            PieceID:      uint(i),
            GridPosition: "0", // Default position
            PieceType:    pieceType,
        }
    }
	return &Player{party, pname, pieces}
}

func GetNewGame(p1 *Player, p2 *Player, currentPlayer *Player, grid *Board) *Game {
	return &Game{
		Player1: p1,
		Player2: p2,
		CurrentPlayer: currentPlayer,
		Grid: grid,
	}
}

func allPiecesFinished(player *Player) bool {
	for _, pieces := range player.Pieces {
		if grid, err := strconv.Atoi(pieces.GridPosition[1:]); err != nil || grid != 15 {
			return false
		}
	}
	return true
}

func GetWinner(game *Game) *Player {
	if allPiecesFinished(game.Player1) {
		game.Winner = game.Player1
		return game.Player1
	} else if allPiecesFinished(game.Player2) {
		game.Winner = game.Player2
		return game.Player2
	} else {
		return nil
	}
}

func Remove[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}


func PrintPlayerProgress(player *Player) {
	count := 0

	for _, piece := range player.Pieces {
		if piece.GridPosition[1:] == "15" {
			count += 1
		}
	}

	fmt.Printf("%s's progress %d/%d\n", player.PlayerName, count, len(player.Pieces))
}

func PrintMap(BoardState map[string]*ChessPiece) {
    fmt.Printf("len: %d, map: {", len(BoardState))
    for k, v := range BoardState {
        fmt.Printf("%s: %c, ", k, v.PieceType)
    }
    fmt.Println("}")
}

func PrintPlayerPieces(player *Player) string {
    s := ""
	for _, piece := range player.Pieces {
		if piece.GridPosition != "0" {
			s += fmt.Sprintf("ID: %d, GridPos: %s\n", piece.PieceID, piece.GridPosition)
		}
	}
    return s
}

func ParseMove(num_moves int) int {
	var input string
	var err error
	var num int

	// Keep asking for input until a valid digit is entered
	for {
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		num, err = strconv.Atoi(input)
		if err != nil || num < 0 || num >= int(num_moves) {
			fmt.Printf("Invalid input. Please enter a digit from 0 to %d.\n", num_moves - 1)
			continue
		}

		break
	}

	return num
}

func PrintBoard(board *Board) string {
    var sb strings.Builder

    // Print First Row
    for row := range GridLength {
        if row == 4 || row == 5 {
            if row == 4 {
                sb.WriteString("-  ")
                continue
            }
            sb.WriteString("   ")
            continue
        }
        sb.WriteString("---")
    }

    sb.WriteString("-\n")

    // Handle printing pieces
    for _, col := range []int{4, 3, 2, 1, -1, -2, 14, 13} {
        if col < 0 {
            if col == -1 {
                sb.WriteString("|  ")
                continue
            }
            sb.WriteString("   ")
            continue
        }
        sb.WriteString("|")
        sb.WriteString(PrintUnit("A"+strconv.Itoa(col), board))
        sb.WriteString(" ")
    }

    sb.WriteString("|\n")

    // Second Row
    for range GridLength {
        sb.WriteString("---")
    }
    sb.WriteString("-\n")

    for _, col := range []int{5, 6, 7, 8, 9, 10, 11, 12} {
        sb.WriteString("|")
        sb.WriteString(PrintUnit(strconv.Itoa(col), board))
        sb.WriteString(" ")
    }
    sb.WriteString("|\n")

    // Third Row
    for range GridLength {
        sb.WriteString("---")
    }

    sb.WriteString("-\n")

    for _, col := range []int{4, 3, 2, 1, -1, -2, 14, 13} {
        if col < 0 {
            if col == -1 {
                sb.WriteString("|  ")
                continue
            }
            sb.WriteString("   ")
            continue
        }
        sb.WriteString("|")
        sb.WriteString(PrintUnit("B"+strconv.Itoa(col), board))
        sb.WriteString(" ")
    }

    sb.WriteString("|\n")

    for row := range GridLength {
        if row == 4 || row == 5 {
            if row == 4 {
                sb.WriteString("-  ")
                continue
            }
            sb.WriteString("   ")
            continue
        }
        sb.WriteString("---")
    }
    sb.WriteString("-\n")
    
    return sb.String()
}

func PrintUnit(unit string, board *Board) string {
    if piece, ok := board.BoardState[unit]; ok {
        return string(piece.PieceType)
    }
    return " "
}