package frontend

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
)

type Knight struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (k *Knight) coordinates(playerWhite bool) (int32, int32) {
	pos := backend.Position{k.p.X, k.p.Y}
	if playerWhite {
		return (3 + k.p.X*22) * 3, (3 + k.p.Y*22) * 3
	} else {
		pos = *pos.Simmetry()
		return (3 + pos.X*22) * 3, (3 + pos.Y*22) * 3
	}
}

func (k *Knight) destroy() error {
	err := k.t.Destroy()
	if err != nil {
		return err
	}

	return nil
}

func (k *Knight) getPosition() backend.Position {
	return k.p
}

func (k *Knight) canMove(pos backend.Position) bool {
	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}

	if pos.Y == k.p.Y+2 && pos.X == k.p.X+1 {
		return true
	}
	if pos.Y == k.p.Y-2 && pos.X == k.p.X-1 {
		return true
	}
	if pos.Y == k.p.Y+2 && pos.X == k.p.X-1 {
		return true
	}
	if pos.Y == k.p.Y-2 && pos.X == k.p.X+1 {
		return true
	}

	if pos.Y == k.p.Y+1 && pos.X == k.p.X+2 {
		return true
	}
	if pos.Y == k.p.Y-1 && pos.X == k.p.X-2 {
		return true
	}
	if pos.Y == k.p.Y+1 && pos.X == k.p.X-2 {
		return true
	}
	if pos.Y == k.p.Y-1 && pos.X == k.p.X+2 {
		return true
	}

	return false
}

func (k *Knight) move(pos backend.Position) {
	k.p = pos
}

func (k *Knight) draw(r *sdl.Renderer, playerWhite bool) error {
	x, y := k.coordinates(playerWhite)

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(k.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil
}

func (k *Knight) setColour(white bool) {
	k.isWhite = white
}

func (k Knight) isColourWhite() bool {
	return k.isWhite
}
