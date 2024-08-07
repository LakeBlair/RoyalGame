package internal

import (
	"github.com/gorilla/websocket"
)

type GameSession struct {
    ID string
    Game *Game
    Connections []*websocket.Conn
	MoveDataChannel chan int
}

type GameMessage struct {
	CurrentPlayer int	`json:"player"`
	MessageRecevier int	`json:"receiver"`
    Msg_Type string     `json:"type"`
    Move int            `json:"move"`
	Content string		`json:"content"`
}


type Msg_Type uint
const (
    Start Msg_Type = iota
	Grid
	TurnStart
	Dice
	MakeMove
	Progress
	AnnounceWinner
)

type Game struct {
	GameID uint
	Turn uint
	Player1 *Player
	Player2 *Player
	CurrentPlayer *Player
	Winner *Player
	Grid *Board
	BonusRound bool
}

type Player struct {
	Party rune
	PlayerName string
	Pieces []*ChessPiece
}

// -------------     ---------
// |A4|A3|A2|A1|     |A14|A13|
// ---------------------------
// | 5| 6| 7| 8| 9|10|11|12|
// ---------------------------
// |B4|B3|B2|B1|     |B14|B13|
// -------------     ---------
type Board struct {
	BoardState map[string]*ChessPiece
}

type ChessPiece struct {
	PieceID uint
	GridPosition string
	PieceType rune
}

type ChessState uint
const (
	NotInPlay ChessState = iota
	InPlay
	Finished
)

type DiceColor uint
const (
	Black DiceColor = iota // Black is 0
	White // White indicates 1 move
)

type RoyalDice struct {
	Color DiceColor
}