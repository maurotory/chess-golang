package frontend

import (
	"fmt"

	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (board *Board) Render() error {
	board.renderer.Clear()

	bg, err := img.LoadTexture(board.renderer, "imgs/board.png")
	if err != nil {
		return fmt.Errorf("could not load background: %v", err)
	}
	defer bg.Destroy()

	if err := board.renderer.Copy(board.background, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	if board.selectedPiece != nil {
		piece := board.selectedPiece

		x, y := piece.coordinates(board.PlayerWhite)

		rect := &sdl.Rect{X: x - 3, Y: y - 3, W: 22 * 3, H: 22 * 3}

		sq, err := img.LoadTexture(board.renderer, "imgs/square_green.png")
		if err != nil {
			fmt.Printf("Could not load green square: %v", err)
		}
		defer sq.Destroy()

		if err := board.renderer.Copy(sq, nil, rect); err != nil {
			fmt.Printf("could not copy green square: %v", err)
		}

		for _, piece := range board.pieces {
			err := piece.draw(board.renderer, board.PlayerWhite)
			if err != nil {
				return fmt.Errorf("could not draw the piece: %v", err)
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if canMove(board.pieces, board.selectedPiece, backend.Position{int32(i), int32(j)}) {
					// if board.selectedPiece.canMove(backend.Position{int32(i), int32(j)}) {
					pos := backend.Position{int32(i), int32(j)}
					if !board.whiteTurn {
						pos = *pos.Simmetry()
					}
					x, y := pos.GetCordinates()
					rect := &sdl.Rect{X: x + 22, Y: y + 22, W: 5 * 3, H: 5 * 3}
					cl, err := img.LoadTexture(board.renderer, "imgs/circle_red.png")
					if err != nil {
						fmt.Printf("could not red circle: %v", err)
					}
					defer cl.Destroy()
					if err := board.renderer.Copy(cl, nil, rect); err != nil {
						fmt.Printf("could not copy red circle: %v", err)
					}
				}
			}
		}
	} else {
		for _, piece := range board.pieces {
			err := piece.draw(board.renderer, board.PlayerWhite)
			if err != nil {
				return fmt.Errorf("could not draw the piece: %v", err)
			}
		}
	}
	board.renderer.Present()

	return nil
}
