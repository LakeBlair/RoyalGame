package internal

type Game struct {
	GameID uint
	Turn uint
	Player1 Player
	Player2 Player
	Winner Player
}

type Player struct {
	PlayerID uint
	PlayerName string
	Pieces []ChessPiece
}

type Board struct {
	BoardState map[string]interface{}
}

type ChessPiece struct {
	PieceID uint
	State ChessState
	GridPosition string
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