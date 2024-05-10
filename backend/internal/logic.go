package internal

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var BonusTile = map[string]struct{}{
	"A4": {},
	"B4": {},
	"8": {},
	"A14": {},
	"B14": {},
}

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
	if game.BonusRound {
		fmt.Printf("%s activates an extra turn!!\n", game.CurrentPlayer.PlayerName)
		game.BonusRound = false
		return
	}

	if game.CurrentPlayer == game.Player1 {
		game.CurrentPlayer = game.Player2
	} else {
		game.CurrentPlayer = game.Player1
	}
}

func isBonusTile(move string) bool {
	_, ok := BonusTile[move]
	return ok
}

func findNewMove(input string, move uint) (uint) {
    var numericPart string

    // Extract the numeric part of the input string
    if len(input) > 0 && unicode.IsLetter(rune(input[0])) {
        numericPart = input[1:]
    } else {
        numericPart = input
    }

    gridPos, _ := strconv.Atoi(numericPart)
    newMove := move + uint(gridPos)

    return newMove
}

func unitOccupied(game *Game, potential_move string) bool {
	_, occupied := game.Grid.BoardState[potential_move]
    return occupied
}

func isEnemyPiece(game *Game, newMove string) bool {
	return game.CurrentPlayer.Pieces[0].PieceType != game.Grid.BoardState[newMove].PieceType
}

func findChessPiece(pieces []*ChessPiece, piece *ChessPiece) *ChessPiece {
    for _, p := range pieces {
        if areChessPiecesEqual(p, piece) {
            return p
        }
    }
    return nil
}

func areChessPiecesEqual(piece1, piece2 *ChessPiece) bool {
    if piece1 == nil && piece2 == nil {
        return true
    }
    if piece1 == nil || piece2 == nil {
        return false
    }
    return piece1.PieceID == piece2.PieceID &&
        piece1.GridPosition == piece2.GridPosition &&
        piece1.PieceType == piece2.PieceType
}

func addMove(moves *[]*Game, messages *[]string, potential_game *Game, message *string, count *uint) {
	for _, str := range *messages {
		if strings.Contains(str, (*message)[3:]) {
			return
		}
	}
	*moves = append(*moves, potential_game)
	*messages = append(*messages, *message)
	*count += 1
}

func getOpponentPlayer(game *Game) *Player {
	if game.CurrentPlayer == game.Player1 {
		return game.Player2
	} else {
		return game.Player1
	}
}

func handleCapturing(game *Game, move string) {
	opponent := getOpponentPlayer(game)
	for _, piece := range opponent.Pieces {
		if piece.GridPosition == move {
			piece.GridPosition = "0"
			return
		}
	}
}

func handleNewMove(game *Game, piece *ChessPiece, move string, message *string) {
	if isBonusTile(move) {
		game.BonusRound = true
		*message += " (Bonus!)"
	}

	delete(game.Grid.BoardState, piece.GridPosition)
	piece.GridPosition = move
	game.Grid.BoardState[move] = piece
}

func findMoves(game *Game, move uint) ([]*Game, []string) {
	var moves []*Game = make([]*Game, 0)
	var messages []string = make([]string, 0)
	var move_count uint = 0

	for _, piece := range game.CurrentPlayer.Pieces {
		var potential_game *Game = game.DeepCopy()
		var potential_piece *ChessPiece = findChessPiece(potential_game.CurrentPlayer.Pieces, piece)
		var potential_move string
		var message string = strconv.Itoa(int(move_count)) + ". "

		newMove := findNewMove(potential_piece.GridPosition, move)
		if newMove <= 4 || newMove >= 13 {
			if (newMove > 15) { // No available move
				continue
			}

			potential_move = string(potential_game.CurrentPlayer.Party) + strconv.Itoa(int(newMove))
			if unitOccupied(potential_game, potential_move) { // Cannot do anything to friendly pieces
				continue
			}

			if (newMove == 15) { // Finish this piece
				message += "Ascended a piece from tile " + potential_piece.GridPosition + " (Finish!)"
			} else { // Move to an empty unit
				message += "Move a piece from tile " + potential_piece.GridPosition + " to tile " + potential_move
			}

			handleNewMove(potential_game, potential_piece, potential_move, &message)
			addMove(&moves, &messages, potential_game, &message, &move_count)
		} else { // The new move is somewhere between 5-12
			potential_move = strconv.Itoa(int(newMove))
			jump := false

			if unitOccupied(potential_game, potential_move) {
				if (isEnemyPiece(potential_game, potential_move)) { // Potentially eat enemy piece
					if (potential_move == "8") {
						if !unitOccupied(potential_game, "9") { // If enemy is on 8, potentially jump to 9 if 9 is empty
							potential_move = "9"
							message += "Jump the piece from tile " + potential_piece.GridPosition + " to tile " + potential_move
							jump = true
						} else {
							continue // It cannot jump, therefore no move
						}
					}
					if !jump {
						message += "Move a piece from tile " + potential_piece.GridPosition + " to tile " + potential_move + " (Capture!)"
						handleCapturing(potential_game, potential_move)
					}
					handleNewMove(potential_game, potential_piece, potential_move, &message)
					addMove(&moves, &messages, potential_game, &message, &move_count)
				}
			} else { // Potentially move to an empty unit
				message += "Move a piece from tile " + potential_piece.GridPosition + " to tile " + potential_move
				handleNewMove(potential_game, potential_piece, potential_move, &message)
				addMove(&moves, &messages, potential_game, &message, &move_count)
			}
		}
	}

	return moves, messages
}

func Play(game *Game) {
	fmt.Println("Game started")
	var winner *Player = nil
	var player_move uint

	PrintBoard(game.Grid)
	for winner == nil {
		fmt.Printf("It's %s's turn...\n", game.CurrentPlayer.PlayerName)
		fmt.Printf("%s, please throw your dices...\n", game.CurrentPlayer.PlayerName)

		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadByte()
		dices_res := throwDices() 

		if dices_res == 0 {
			fmt.Println(game.CurrentPlayer.PlayerName + " rolled 0 LOL. Your turn is skipped...")
			switchCurrentPlayer(game)
			continue
		}

		moves, messages := findMoves(game, dices_res)
		if len(moves) == 0 {
			fmt.Println("You don't have any moves available, switching players...")
			continue
		}

		fmt.Printf("%s, please choose your move\n\n", game.CurrentPlayer.PlayerName)
		for _, message := range messages {
			fmt.Println(message)
		}

		player_move = uint(ParseMove(len(moves)))
		game = moves[player_move]
		switchCurrentPlayer(game)
		PrintMap(game.Grid.BoardState)
		PrintBoard(game.Grid)
		PrintPlayerProgress(game.Player1)
		PrintPlayerPieces(game.Player1)
		PrintPlayerProgress(game.Player2)
		PrintPlayerPieces(game.Player2)
		winner = GetWinner(game)
	}

	fmt.Printf("Game over, the winner is %s\n", game.Winner.PlayerName)
}