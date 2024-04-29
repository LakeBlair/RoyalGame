package internal


func (cp ChessPiece) DeepCopy() ChessPiece {
    return ChessPiece{
        PieceID:      cp.PieceID,
        GridPosition: cp.GridPosition,
        State:        cp.State,
        PieceType:    cp.PieceType,
    }
}


func (p *Player) DeepCopy() *Player {
    if p == nil {
        return nil
    }
    newPieces := make([]ChessPiece, len(p.Pieces))
    for i, piece := range p.Pieces {
        newPieces[i] = piece.DeepCopy()
    }
    return &Player{
        PlayerName: p.PlayerName,
        Pieces:     newPieces,
    }
}


func (b *Board) DeepCopy() *Board {
    if b == nil {
        return nil
    }
    newBoardState := make(map[ChessPiece]struct{})
    for k, v := range b.BoardState {
        newBoardState[k.DeepCopy()] = v
    }
    return &Board{
        BoardState: newBoardState,
    }
}


func (g *Game) DeepCopy() *Game {
    if g == nil {
        return nil
    }
    newGame := &Game{
        GameID: g.GameID,
        Turn:   g.Turn,
        Player1:        g.Player1.DeepCopy(),
        Player2:        g.Player2.DeepCopy(),
        CurrentPlayer:  nil,
        Winner:         nil,
        Grid:           g.Grid.DeepCopy(),
    }

    // Assign CurrentPlayer and Winner by reference
    if g.CurrentPlayer == g.Player1 {
        newGame.CurrentPlayer = newGame.Player1
    } else if g.CurrentPlayer == g.Player2 {
        newGame.CurrentPlayer = newGame.Player2
    }

    if g.Winner == g.Player1 {
        newGame.Winner = newGame.Player1
    } else if g.Winner == g.Player2 {
        newGame.Winner = newGame.Player2
    }

    return newGame
}
