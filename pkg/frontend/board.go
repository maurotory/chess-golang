package frontend

import (
	"fmt"
	"io"
	"log"
	"time"

	//"time"

	"github.com/google/uuid"
	pb "github.com/maurotory/chess-golang/api/proto"
	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	pieces        []Piece
	whiteTurn     bool
	window        *sdl.Window
	renderer      *sdl.Renderer
	background    *sdl.Texture
	selectedPiece Piece
	PlayerID      uuid.UUID
	PlayerWhite   bool
}

func InitBoard() (*Board, error) {
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

	board := &Board{pieces: pieces, whiteTurn: true, window: w, renderer: r, background: bg, selectedPiece: nil}

	return board, nil
}

func (board *Board) RunBoard(stream pb.SendMoveRequest_MoveClient) (chan int, error) {
	done := make(chan int)
	go func() {
		fmt.Println("Listening to events...")
		go func() {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Fatalf("can not receive data: %v", err)
				}

				initPos := backend.Position{X: resp.XInitPos, Y: resp.YInitPos}
				finalPos := backend.Position{X: resp.XFinalPos, Y: resp.YFinalPos}
				fmt.Println(resp.WhiteTurn)
				board.whiteTurn = resp.WhiteTurn
				for _, piece := range board.pieces {
					if piece.getPosition() == initPos {
						for i, piece2 := range board.pieces {
							if piece2.getPosition() == finalPos {
								board.pieces = append(board.pieces[:i], board.pieces[i+1:]...)
							}
						}
						piece.move(finalPos)
						if board.IsCheck() {
							log.Println("It is check!")
							// if board.IsCheckMate() {
							// log.Println("GAME OVER, YOU LOSE")
							// os.Exit(0)
							// }
						}
						board.Render()
					}
				}
			}
		}()
		for {
			event := sdl.WaitEvent()
			err := board.handleEvents(event, done, stream)
			if err != nil {
				fmt.Printf("Could not handle events %v", err)
			}
		}
	}()
	return done, nil
}

func (board *Board) handleEvents(e sdl.Event, done chan int, stream pb.SendMoveRequest_MoveClient) error {
	switch e := e.(type) {
	case *sdl.QuitEvent:
		done <- 0
	case *sdl.MouseButtonEvent:
		if e.Type == 1025 && board.whiteTurn == board.PlayerWhite {
			pos := backend.GetCell(e.X, e.Y)
			if !board.PlayerWhite {
				pos = *pos.Simmetry()
			}

			if board.selectedPiece == nil {
				for _, piece := range board.pieces {
					if piece.getPosition() == pos && board.PlayerWhite == piece.isColourWhite() {
						_, ok := piece.(*King)
						if board.IsCheck() && !ok {
							fmt.Println("Can not select this piece, is check!")
						} else {
							log.Printf("The piece selected is: %v\n", piece)
							board.selectedPiece = piece
						}
					}
				}

			} else {
				if canMove(board.pieces, board.selectedPiece, pos) {
					// if board.selectedPiece.canMove(pos) {
					for _, piece := range board.pieces {
						if piece == board.selectedPiece {
							//piece.move(pos)
							playerID := board.PlayerID.String()
							req := pb.MoveRequest{
								XInitPos:  board.selectedPiece.getPosition().X,
								YInitPos:  board.selectedPiece.getPosition().Y,
								XFinalPos: pos.X,
								YFinalPos: pos.Y,
								Id:        playerID,
							}
							stream.Send(&req)
						}
					}
					board.selectedPiece = nil

				} else {
					log.Println("This is not a valid position")
					board.selectedPiece = nil
				}
			}
			time.Sleep(time.Millisecond * 100)
			board.Render()
		}
	}

	return nil
}

func (board *Board) FinishBoard() error {
	defer sdl.Quit()
	defer board.window.Destroy()
	defer board.background.Destroy()

	for _, p := range board.pieces {

		if err := p.destroy(); err != nil {
			return fmt.Errorf("could not destroy the piece: %v", err)
		}
	}
	return nil
}

func (board *Board) IsCheck() bool {
	for _, p := range board.pieces {
		if p.isColourWhite() == board.whiteTurn {
			if _, ok := p.(*King); ok {
				if isCheck(board.pieces, p, backend.Position{X: p.getPosition().X, Y: p.getPosition().Y}) {
					log.Println("It is check to the king")
					return true
				}
			}
		}
	}
	return false
}

func (board *Board) IsCheckMate() bool {
	copyPieces := make([]Piece, len(board.pieces))

	copy(copyPieces, board.pieces)
	for xx, piece := range copyPieces {
		fmt.Println(xx)
		if piece.isColourWhite() == board.whiteTurn {
			var x int32
			var y int32
			for x = 0; x < 8; x++ {
				for y = 0; y < 8; y++ {
					pos := backend.Position{X: x, Y: y}
					if canMove(copyPieces, piece, pos) {
						for i, p := range copyPieces {
							if p.getPosition() == pos {
								copyPieces = append(copyPieces[:i], copyPieces[i+1:]...)
							}
						}
						piece.move(pos)
						for _, p := range copyPieces {
							if p.isColourWhite() == board.whiteTurn {
								if _, ok := p.(*King); ok {
									if !isCheck(copyPieces, p, backend.Position{X: p.getPosition().X, Y: p.getPosition().Y}) {
										return false
									}
								}
							}
						}
						copyPieces = nil
						copy(copyPieces, board.pieces)
					}
				}
			}
		}
	}
	return true
}
