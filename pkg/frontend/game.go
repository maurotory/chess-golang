package frontend

import (
	"fmt"
	"log"
	"time"

	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	pieces        []Piece
	whiteTurn     bool
	window        *sdl.Window
	renderer      *sdl.Renderer
	background    *sdl.Texture
	selectedPiece Piece
}

func InitGame() (*Game, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, fmt.Errorf("could not initialize SDL: %v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(180*3, 180*3, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("could not create window: %v", err)
	}

	bg, err := img.LoadTexture(r, "imgs/board.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background: %v", err)
	}

	if err := r.Copy(bg, nil, nil); err != nil {
		return nil, fmt.Errorf("could not copy background: %v", err)
	}

	pieces, err := createPieces(r)
	if err != nil {
		return nil, fmt.Errorf("could not create pieces: %v", err)
	}

	r.Present()

	game := &Game{pieces: pieces, whiteTurn: true, window: w, renderer: r, background: bg, selectedPiece: nil}

	return game, nil
}

func (g *Game) RunGame() error {
	go func() {
		fmt.Println("Listening to events...")
		for {
			event := sdl.WaitEvent()
			err := handleEvents(event, g)
			if err != nil {
				fmt.Printf("Could not handle events %v", err)
			}
		}
	}()
	return nil
}

func handleEvents(e sdl.Event, g *Game) error {
	switch e := e.(type) {

	case *sdl.QuitEvent:
		fmt.Printf("event: %v", e)

	case *sdl.MouseButtonEvent:

		if e.Type == 1025 {
			fmt.Printf("Button pressed is: %b\n", e.Type)

			pos := backend.GetCell(e.X, e.Y)

			if g.selectedPiece == nil {
				for _, piece := range g.pieces {

					if piece.getPosition() == pos {
						fmt.Printf("The piece selected is: %v\n", piece)
						g.selectedPiece = piece
					}
				}

			} else {
				if g.selectedPiece.canMove(pos) {
					for _, piece := range g.pieces {
						if piece == g.selectedPiece {
							piece.move(pos)
						}
					}
					g.selectedPiece = nil

				} else {
					log.Println("This is not a valid position")
					g.selectedPiece = nil
				}
			}
			time.Sleep(time.Millisecond * 100)
			render(g)
		}
	}

	return nil
}

func (g *Game) FinishGame() error {
	defer sdl.Quit()
	defer g.window.Destroy()
	defer g.background.Destroy()

	for _, p := range g.pieces {

		if err := p.destroy(); err != nil {
			return fmt.Errorf("could not destroy the piece: %v", err)
		}
	}

	return nil
}
