package frontend

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/backend"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Queen struct {
	t       *sdl.Texture
	p       backend.Position
	isWhite bool
}

func (q *Queen) coordinates() (int32, int32) {
	return (3 + q.p.X*22) * 3, (3 + q.p.Y*22) * 3
}

func (q *Queen) getPosition() backend.Position {
	return q.p
}

func (q *Queen) destroy() error {
	err := q.t.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (q *Queen) canMove(pos backend.Position) bool {
	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}

	if pos.X == q.p.X && pos.Y == q.p.Y {
		return false
	}

	diff := math.Abs(float64((q.p.X - pos.X)))

	if diff == math.Abs(float64((q.p.Y - pos.Y))) {
		return true
	}

	if pos.X > 7 || pos.Y > 7 || pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X == q.p.X && pos.Y != q.p.Y {
		return true
	}
	if pos.X != q.p.X && pos.Y == q.p.Y {
		return true
	}
	return false
}

func (q *Queen) move(pos backend.Position) {
	q.p = pos
}

func (q *Queen) draw(r *sdl.Renderer) error {
	x, y := q.coordinates()

	rect := &sdl.Rect{X: x, Y: y, W: 20 * 3, H: 20 * 3}

	if err := r.Copy(q.t, nil, rect); err != nil {
		return fmt.Errorf("could not copy bishop: %v", err)
	}

	return nil
}
func (q *Queen) setColour(white bool) {
	q.isWhite = white
}

func (q *Queen) isColourWhite() bool {
	return q.isWhite
}
