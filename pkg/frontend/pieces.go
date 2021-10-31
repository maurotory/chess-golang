package frontend

import (
	"fmt"

	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func blockedPath(piece Piece, pos backend.Position, pieces []Piece) bool {

	return false
}

type Piece interface {
	coordinates(bool) (int32, int32)
	destroy() error
	getPosition() backend.Position
	canMove(backend.Position) bool
	move(backend.Position)
	draw(*sdl.Renderer, bool) error
	setColour(bool)
	isColourWhite() bool
}

func createPieces(r *sdl.Renderer) ([]Piece, error) {
	bBishopt, err := img.LoadTexture(r, "imgs/black_bishop.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Bishop: %v", err)
	}
	bBishop1 := &Bishop{t: bBishopt, p: backend.Position{3, 0}, isWhite: false}

	bPawnt, err := img.LoadTexture(r, "imgs/black_pawn.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Pawn: %v", err)
	}
	bPawn1 := &Pawn{t: bPawnt, p: backend.Position{2, 1}, isWhite: false}

	brookt, err := img.LoadTexture(r, "imgs/black_rook.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Rook: %v", err)
	}
	bRook1 := &Rook{t: brookt, p: backend.Position{0, 0}, isWhite: false}

	bqueent, err := img.LoadTexture(r, "imgs/black_queen.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Queen: %v", err)
	}
	bQueen := &Queen{t: bqueent, p: backend.Position{5, 0}, isWhite: false}

	bkingt, err := img.LoadTexture(r, "imgs/black_king.png")
	if err != nil {
		return nil, fmt.Errorf("could not load King: %v", err)
	}
	bKing := &King{t: bkingt, p: backend.Position{6, 0}, isWhite: false}

	bknightt, err := img.LoadTexture(r, "imgs/black_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	bKnight1 := &Knight{t: bknightt, p: backend.Position{6, 1}, isWhite: false}

	wBishopt, err := img.LoadTexture(r, "imgs/white_bishop.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Bishop: %v", err)
	}
	wBishop1 := &Bishop{t: wBishopt, p: backend.Position{3, 6}, isWhite: true}

	wPawnt, err := img.LoadTexture(r, "imgs/white_pawn.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Pawn: %v", err)
	}
	wPawn1 := &Pawn{t: wPawnt, p: backend.Position{2, 7}, isWhite: true}

	wrookt, err := img.LoadTexture(r, "imgs/white_rook.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Rook: %v", err)
	}
	wRook1 := &Rook{t: wrookt, p: backend.Position{0, 5}, isWhite: true}

	wqueent, err := img.LoadTexture(r, "imgs/white_queen.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Queen: %v", err)
	}
	wQueen := &Queen{t: wqueent, p: backend.Position{5, 5}, isWhite: true}

	wkingt, err := img.LoadTexture(r, "imgs/white_king.png")
	if err != nil {
		return nil, fmt.Errorf("could not load King: %v", err)
	}
	wKing := &King{t: wkingt, p: backend.Position{5, 7}, isWhite: true}

	wknightt, err := img.LoadTexture(r, "imgs/white_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	wKnight1 := &Knight{t: wknightt, p: backend.Position{6, 7}, isWhite: true}

	pieces := []Piece{bBishop1, bPawn1, bRook1, bQueen, bKing, bKnight1, wKnight1, wBishop1, wPawn1, wRook1, wQueen, wKing}

	// white := createWhitePieces(pieces)
	// pieces = append(pieces, white...)

	return pieces, nil
}

func makeSimmetry(pieces []Piece) []Piece {
	var simmetryPieces []Piece
	for _, piece := range pieces {
		var newPiece Piece
		newPiece = piece
		pos := newPiece.getPosition()
		if pos.X < 3 {
			pos.X = 8 - pos.X
		} else {
			pos.X = 3 - (pos.X - 3)
		}
		if pos.Y < 3 {
			pos.Y = 8 - pos.Y
		} else {
			pos.Y = 3 - (pos.Y - 3)
		}

		newPiece.move(pos)
		simmetryPieces = append(simmetryPieces, newPiece)
	}

	return simmetryPieces
}

func createWhitePieces(pieces []Piece) []Piece {
	var whitePieces []Piece

	simmetryPieces := makeSimmetry(pieces)

	for _, piece := range simmetryPieces {
		piece.setColour(true)

		whitePieces = append(whitePieces, piece)
	}

	return whitePieces
}

func canMove(pieces []Piece, piece Piece, pos backend.Position) bool {
	if pos == piece.getPosition() {
		return false
	}
	if !piece.canMove(pos) {
		return false
	}
	switch piece.(type) {
	case *Knight:
		if piece.canMove(pos) {
			return true
		}
	default:
		return checkStraight(pieces, piece, pos)
	}

	return false
}

func checkStraight(pieces []Piece, piece Piece, pos backend.Position) bool {
	if pos.X == piece.getPosition().X && pos.Y != piece.getPosition().Y {
		dist := pos.Y - piece.getPosition().Y
		var i int32
		if dist < 0 {
			for i = -1; i >= dist; i-- {
				for _, p := range pieces {
					if p.getPosition().Y == piece.getPosition().Y+i && p.getPosition().X == piece.getPosition().X {
						return false
					}
				}
			}
		} else {
			for i = 1; i <= dist; i++ {
				for _, p := range pieces {
					if p.getPosition().Y == piece.getPosition().Y+i && p.getPosition().X == piece.getPosition().X {
						return false
					}
				}
			}
		}
	} else if pos.Y == piece.getPosition().Y && pos.X != piece.getPosition().X {
		dist := pos.X - piece.getPosition().X
		var i int32
		if dist < 0 {
			for i = -1; i >= dist; i-- {
				for _, p := range pieces {
					if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y {
						return false
					}
				}
			}
		} else {
			for i = 1; i <= dist; i++ {
				for _, p := range pieces {
					if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y {
						return false
					}
				}
			}
		}
	} else {
		return checkDiagonal(pieces, piece, pos)
	}

	return true
}

func checkDiagonal(pieces []Piece, piece Piece, pos backend.Position) bool {
	dist := pos.X - piece.getPosition().X
	if dist < 0 {
		dist = -dist
	}
	var i int32
	if pos.Y > piece.getPosition().Y && pos.X > piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y+i {
					return false
				}
			}
		}
	} else if pos.Y < piece.getPosition().Y && pos.X > piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y-i {
					return false
				}
			}
		}
	} else if pos.Y > piece.getPosition().Y && pos.X < piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X-i && p.getPosition().Y == piece.getPosition().Y+i {
					return false
				}
			}
		}
	} else if pos.Y < piece.getPosition().Y && pos.X < piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X-i && p.getPosition().Y == piece.getPosition().Y-i {
					return false
				}
			}
		}
	}

	return true
}
