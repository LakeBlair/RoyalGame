package objects

type Game struct {
	GameID uint
	Player1 Player
	Player2 Player
}

type Player struct {
	PlayerID uint
	PlayerName string
	Pieces []ChessPiece
}

type Board struct {
	BoardID uint
	BoardState [][]byte
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
	 