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