package internal

import (
	"fmt"
	"math/rand"
	"log"
	"strconv"
	"strings"
	"unicode"

	// "github.com/gorilla/websocket"
)

var BonusTile = map[string]struct{}{
	"A4": {},
	"B4": {},
	"8": {},
	"A14": {},
	"B14": {},
}

func Init_Game() (game *Game) {
	fmt.Println("Initilizing the game...")

	p1, p2 := GetNewPlayer("player1", true), GetNewPlayer("player2", false)
	fmt.Println("P1's team is '@'")
	fmt.Println("P2's team is '*'")
	playerGoFirst := goFirst(p1, p2)
	grid := GetNewBoard()
	game = GetNewGame(p1, p2, playerGoFirst, grid)

	return game
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

func throwDices(s *GameSession) uint {
	var move uint = 0
	for i := 0; i < 4; i++  {
		if randomColor() == White {
			move += 1
		}
	}
	SendToAll(s, Dice, fmt.Sprintf("%s's total moves is %d\n", s.Game.CurrentPlayer.PlayerName, move), 0)
	return move
}

func randomColor() DiceColor {
    // Generate a random number 0 or 1 and return corresponding color
    if rand.Intn(2) == 0 {
		log.Println("You got Black")
        return Black
    }
	log.Println("You got White")
    return White
}

func switchCurrentPlayer(game *Game) {
	if game.BonusRound {
		log.Printf("%s activates an extra turn!!\n", game.CurrentPlayer.PlayerName)
		game.BonusRound = false
		return
	}

	if game.CurrentPlayer == game.Player1 {
		game.CurrentPlayer = game.Player2
		log.Printf("Switching to %s\n", game.Player2.PlayerName)
	} else {
		game.CurrentPlayer = game.Player1
		log.Printf("Switching to %s\n", game.Player1.PlayerName)
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
	if potential_move[1:] == "15" {
		return false
	}
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


func Play(s *GameSession) {
	fmt.Println("Game started")
	s.Game = Init_Game()
	var winner *Player = nil

	SendToAll(s, Start, "", 0)
	SendToAll(s, Grid, PrintBoard(s.Game.Grid), 0) 
	for winner == nil {
		SendToAll(s, TurnStart, fmt.Sprintf("It's %s's turn...\n", s.Game.CurrentPlayer.PlayerName), 0)

		// fmt.Printf("%s, please throw your dices...\n", s.Game.CurrentPlayer.PlayerName)

		// reader := bufio.NewReader(os.Stdin)
		// _, _ = reader.ReadByte()
		dices_res := throwDices(s) 
		
		if dices_res == 0 {
			// SendToAll(s, Dice, fmt.Sprintf(s.Game.CurrentPlayer.PlayerName + " rolled 0 LOL. Their turn is skipped..."))
			log.Printf(s.Game.CurrentPlayer.PlayerName + " rolled 0 LOL. Their turn is skipped...\n")
			switchCurrentPlayer(s.Game)
			continue
		}
		SendToAll(s, Dice, fmt.Sprintf("%s rolled %d\n", s.Game.CurrentPlayer.PlayerName, dices_res), 0)

		moves, messages := findMoves(s.Game, dices_res)
		if len(moves) == 0 {
			SendToAll(s, MakeMove, fmt.Sprintf("%s don't have any moves available, switching players...", s.Game.CurrentPlayer.PlayerName), 0)
			switchCurrentPlayer(s.Game)
			continue
		}

		str_moves := ""
		for _, m := range messages {
			str_moves += m + "\n"
		}
		SendToAll(s, MakeMove, str_moves, 0)

		for _, message := range messages {
			fmt.Println(message)
		}

		player_move := <-s.MoveDataChannel
		s.Game = moves[player_move]
		switchCurrentPlayer(s.Game)
		SendToAll(s, Grid, PrintBoard(s.Game.Grid), 0) 

		SendToAll(s, Progress, PrintPlayerPieces(s.Game.Player1), 1)
		SendToAll(s, Progress, PrintPlayerPieces(s.Game.Player2), 2)
		// PrintPlayerProgress(s.Game.Player2)
		winner = GetWinner(s.Game)
	}

	SendToAll(s, AnnounceWinner, fmt.Sprintf("Game over, the winner is %s\n", s.Game.Winner.PlayerName), 0)
}

func getCurrentPlayerID(session *GameSession) int {
	if session.Game.CurrentPlayer == session.Game.Player1 {
		return 1
	}
	return 2
}

func SendToAll(session *GameSession, Msg_type Msg_Type, message string, receiver_id int) {
	var msg GameMessage

	SetGameParam(session, &msg, Msg_type, message, receiver_id)
	for _, player := range session.Connections {
		err := player.WriteJSON(msg)
		if err != nil {
			log.Println("Sending error: ", err)
		}
	}
}

func SetGameParam(session *GameSession, Msg *GameMessage, Msg_type Msg_Type, message string, receiver_id int) {
	Msg.Content = message
	Msg.CurrentPlayer = getCurrentPlayerID(session)

	if Msg_type == Start {
		Msg.Msg_Type = "Start_ACK"
	} else if Msg_type == MakeMove {
		Msg.Msg_Type = "Move"
		Msg.Move = strings.Count(message, "\n")
	} else if Msg_type == Grid {
		Msg.Msg_Type = "Grid"
	} else if Msg_type == TurnStart {
		Msg.Msg_Type = "TurnStart"
	} else if Msg_type == Dice {
		Msg.Msg_Type = "Dice"
	} else if Msg_type == Progress {
		Msg.Msg_Type = "Progress"
		Msg.MessageRecevier = receiver_id
	} else if Msg_type == AnnounceWinner {
		Msg.Msg_Type = "Winner"
	}
}