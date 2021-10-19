package frontend

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
)

type Pawn struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (p *Pawn) coordinates() (int32, int32) {
	return (3 + p.p.X*22) * 3, (3 + p.p.Y*22) * 3
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
	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}

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

func (p *Pawn) setColour(white bool) {
	p.isWhite = white
}

func (p *Pawn) isColourWhite() bool {
	return p.isWhite
}
