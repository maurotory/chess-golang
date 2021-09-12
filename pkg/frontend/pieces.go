package frontend

import (
	"fmt"

	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Bishop struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (b *Bishop) coordinates() (int32, int32) {
	return (3 + b.p.X*22) * 3, (3 + b.p.Y*22) * 3
}

func (b *Bishop) getPosition() backend.Position {
	return b.p
}

func (b *Bishop) destroy() error {
	err := b.t.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bishop) canMove(pos backend.Position) bool {

	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X == b.p.X && pos.Y == b.p.Y {
		return false
	}

	diff := math.Abs(float64((b.p.X - pos.X)))

	if diff == math.Abs(float64((b.p.Y - pos.Y))) {
		return true
	}
	return false
}

func (b *Bishop) move(pos backend.Position) {
	b.p = pos
}

func (b *Bishop) draw(r *sdl.Renderer) error {

	x, y := b.coordinates()

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(b.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil

}

func blockedPath(piece Piece, pos backend.Position, pieces []Piece) bool {

	return false
}

type Pawn struct {
	t     *sdl.Texture
	p     backend.Position
	white bool
}

func (b *Pawn) coordinates() (int32, int32) {
	return (3 + b.p.X*22) * 3, (3 + b.p.Y*22) * 3
}

func (p *Pawn) destroy() error {

	err := p.t.Destroy()
	if err != nil {
		return err
	}

	return nil
}

func (p *Pawn) getPosition() backend.Position {
	return p.p
}

func (p *Pawn) canMove(pos backend.Position) bool {

	if pos.Y == p.p.Y+2 || pos.Y == p.p.Y+1 {
		if pos.X == p.p.X {
			return true
		}
	}

	return false
}

func (p *Pawn) move(pos backend.Position) {
	p.p = pos
}

func (p *Pawn) draw(r *sdl.Renderer) error {

	x, y := p.coordinates()

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(p.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil

}

type Piece interface {
	coordinates() (int32, int32)
	destroy() error
	getPosition() backend.Position
	canMove(backend.Position) bool
	move(backend.Position)
	draw(*sdl.Renderer) error
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

	bPawn1 := &Pawn{t: bPawnt, p: backend.Position{2, 1}, white: false}

	x, y = bPawn1.coordinates()

	rect = &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(bPawnt, nil, rect); err != nil {
		return nil, fmt.Errorf("Could not copy Pawn: %v", err)
	}

	pieces := []Piece{bBishop1, bPawn1}

	return pieces, nil

}
