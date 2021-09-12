package frontend

import (
	"fmt"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func render(g *Game) error {
	g.renderer.Clear()

	/*

		w, r, err := sdl.CreateWindowAndRenderer(180*3, 180*3, sdl.WINDOW_SHOWN)
		if err != nil {
			return fmt.Errorf("Could not create window: %v", err)
		}
		defer w.Destroy()

	*/

	bg, err := img.LoadTexture(g.renderer, "imgs/board.png")
	if err != nil {
		return fmt.Errorf("could not load background: %v", err)
	}
	defer bg.Destroy()

	if err := g.renderer.Copy(bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	if g.selectedPiece != nil {
		piece := g.selectedPiece

		x, y := piece.coordinates()

		rect := &sdl.Rect{X: x - 3, Y: y - 3, W: 22 * 3, H: 22 * 3}

		sq, err := img.LoadTexture(g.renderer, "imgs/square_green.png")
		if err != nil {
			fmt.Printf("could not load background: %v", err)
		}
		defer sq.Destroy()

		if err := g.renderer.Copy(sq, nil, rect); err != nil {
			fmt.Printf("could not copy background: %v", err)
		}
	}

	for _, piece := range g.pieces {
		err := piece.draw(g.renderer)
		if err != nil {
			return fmt.Errorf("could not draw the piece: %v", err)
		}
	}

	g.renderer.Present()

	return nil
}
