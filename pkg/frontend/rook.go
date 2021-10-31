package frontend

import (
	"fmt"

	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
)

type Rook struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (r *Rook) coordinates(playerWhite bool) (int32, int32) {
	pos := backend.Position{r.p.X, r.p.Y}
	if playerWhite {
		return (3 + r.p.X*22) * 3, (3 + r.p.Y*22) * 3
	} else {
		pos = *pos.Simmetry()
		return (3 + pos.X*22) * 3, (3 + pos.Y*22) * 3
	}
}

func (b *Rook) getPosition() backend.Position {
	return b.p
}

func (b *Rook) destroy() error {
	err := b.t.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (b *Rook) canMove(pos backend.Position) bool {
	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X == b.p.X && pos.Y != b.p.Y {
		return true
	}
	if pos.X != b.p.X && pos.Y == b.p.Y {
		return true
	}

	return false
}

func (b *Rook) move(pos backend.Position) {
	b.p = pos
}

func (b *Rook) draw(r *sdl.Renderer, playerWhite bool) error {
	x, y := b.coordinates(playerWhite)

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(b.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy the tour: %v", err)
	}

	return nil
}

func (r *Rook) setColour(white bool) {
	r.isWhite = white
}

func (r *Rook) isColourWhite() bool {
	return r.isWhite
}
