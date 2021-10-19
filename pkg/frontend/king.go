package frontend

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
	// "math"
)

type King struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (k *King) coordinates() (int32, int32) {
	return (3 + k.p.X*22) * 3, (3 + k.p.Y*22) * 3
}

func (k *King) getPosition() backend.Position {
	return k.p
}

func (k *King) destroy() error {
	err := k.t.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (k *King) canMove(pos backend.Position) bool {
	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X-k.p.X == 1 || pos.X-k.p.X == -1 || pos.X-k.p.X == 0 {
		if pos.Y-k.p.Y == 1 || pos.Y-k.p.Y == -1 || pos.Y-k.p.Y == 0 {
			return true
		}
	}

	// if pos.X == k.p.X && pos.Y == k.p.Y {
	// 	return false
	// }

	// diff := math.Abs(float64((k.p.X - pos.X)))

	// if diff == math.Abs(float64((k.p.Y - pos.Y))) {
	// 	return true
	// }

	// if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
	// 	return false
	// }
	// if pos.X == k.p.X && pos.Y != k.p.Y {
	// 	return true
	// }
	// if pos.X != k.p.X && pos.Y == k.p.Y {
	// 	return true
	// }
	return false
}

func (k *King) move(pos backend.Position) {
	k.p = pos
}

func (k *King) draw(r *sdl.Renderer) error {
	x, y := k.coordinates()

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(k.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil
}

func (k *King) setColour(white bool) {
	k.isWhite = white
}

func (k *King) isColourWhite() bool {
	return k.isWhite
}
