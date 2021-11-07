package frontend

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Bishop struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (b *Bishop) coordinates(playerWhite bool) (int32, int32) {
	pos := backend.Position{X: b.p.X, Y: b.p.Y}
	if playerWhite {
		return (3 + b.p.X*22) * 3, (3 + b.p.Y*22) * 3
	} else {
		pos = *pos.Simmetry()
		return (3 + pos.X*22) * 3, (3 + pos.Y*22) * 3
	}
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

	return diff == math.Abs(float64((b.p.Y - pos.Y)))
}

func (b *Bishop) move(pos backend.Position) {
	b.p = pos
}

func (b *Bishop) draw(r *sdl.Renderer, playerWhite bool) error {
	x, y := b.coordinates(playerWhite)

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(b.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil
}

func (b *Bishop) setColour(white bool) {
	b.isWhite = white
}

func (b *Bishop) isColourWhite() bool {
	return b.isWhite
}
