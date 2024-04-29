package internal

type Game struct {
	GameID uint
	Turn uint
	Player1 *Player
	Player2 *Player
	CurrentPlayer *Player
	Winner *Player
	Grid *Board
}

type Player struct {
	PlayerName string
	Pieces []ChessPiece
}

// -----------------      ---------
// |A04|A03|A02|A01|      |A14|A13|
// --------------------------------
// |P05|P06|P07|P08|P09|P10|P11|P12|
// --------------------------------
// |B04|B03|B02|B01|      |B14|B13|
// -----------------      ---------
type Board struct {
	BoardState map[string]ChessPiece
}

type ChessPiece struct {
	PieceID uint
	GridPosition string
	State ChessState
	PieceType byte
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