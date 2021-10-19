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
	coordinates() (int32, int32)
	destroy() error
	getPosition() backend.Position
	canMove(backend.Position) bool
	move(backend.Position)
	draw(*sdl.Renderer) error
	setColour(bool)
	isColourWhite() bool
}

func createPieces(r *sdl.Renderer) ([]Piece, error) {
	bBishopt, err := img.LoadTexture(r, "imgs/black_bishop.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Bishop: %v", err)
	}

	bBishop1 := &Bishop{t: bBishopt, p: backend.Position{3, 0}, isWhite: false}

	x, y := bBishop1.coordinates()

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bBishopt, nil, rect); err != nil {
		return nil, fmt.Errorf("could not copy bishop: %v", err)
	}

	bPawnt, err := img.LoadTexture(r, "imgs/black_pawn.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Pawn: %v", err)
	}

	bPawn1 := &Pawn{t: bPawnt, p: backend.Position{2, 1}, isWhite: false}

	x, y = bPawn1.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bPawnt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Pawn: %v", err)
	}

	brookt, err := img.LoadTexture(r, "imgs/black_rook.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Rook: %v", err)
	}

	bRook1 := &Rook{t: brookt, p: backend.Position{0, 0}, isWhite: false}

	x, y = bRook1.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(brookt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Rook: %v", err)
	}

	bqueent, err := img.LoadTexture(r, "imgs/black_queen.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Queen: %v", err)
	}

	bQueen := &Queen{t: bqueent, p: backend.Position{5, 0}, isWhite: false}

	x, y = bQueen.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bqueent, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Queen: %v", err)
	}
	bkingt, err := img.LoadTexture(r, "imgs/black_king.png")
	if err != nil {
		return nil, fmt.Errorf("could not load King: %v", err)
	}

	bKing := &King{t: bkingt, p: backend.Position{6, 0}, isWhite: false}

	x, y = bKing.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bkingt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy King: %v", err)
	}

	bknightt, err := img.LoadTexture(r, "imgs/black_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	bKnight1 := &Knight{t: bknightt, p: backend.Position{6, 1}, isWhite: false}

	x, y = bKnight1.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bknightt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Knight: %v", err)
	}

	wknightt, err := img.LoadTexture(r, "imgs/white_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	wKnight1 := &Knight{t: wknightt, p: backend.Position{6, 7}, isWhite: true}

	x, y = wKnight1.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(wknightt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Knight: %v", err)
	}

	pieces := []Piece{bBishop1, bPawn1, bRook1, bQueen, bKing, bKnight1, wKnight1}

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
		fmt.Println(pos)
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
