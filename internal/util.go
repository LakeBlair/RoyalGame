package internal

import (
	"fmt"
)

const (
	BoardLength uint = 17
	BoardWidth uint = 7
)

func GetNewBoard() *Board {
	return &Board{make(map[string]interface{})}
}

// func PrintBoard(grid *Board) {
//     for _, row := range grid.BoardState {
//         for _, col := range row {
//             fmt.Printf("%c ", col)
//         }
//         fmt.Println()
//     }
// }