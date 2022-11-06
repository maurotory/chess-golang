package frontend

import (
	"fmt"
	"log"

	"github.com/maurotory/chess-golang/pkg/backend"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Piece interface {
	coordinates(bool) (int32, int32)
	destroy() error
	getPosition() backend.Position
	canMove(backend.Position) bool
	move(backend.Position)
	draw(*sdl.Renderer, bool) error
	setColour(bool)
	isColourWhite() bool
}

func createPieces(r *sdl.Renderer) ([]Piece, error) {
	var pieces []Piece

	// Load black pieces
	bPawnt, err := img.LoadTexture(r, "imgs/black_pawn.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Pawn: %v", err)
	}
	for i := 0; i < 8; i++ {
		pieces = append(pieces, &Pawn{t: bPawnt, p: backend.Position{X: int32(i), Y: 1}, isWhite: false})
	}

	bBishopt, err := img.LoadTexture(r, "imgs/black_bishop.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Bishop: %v", err)
	}
	pieces = append(pieces, &Bishop{t: bBishopt, p: backend.Position{X: 2, Y: 0}, isWhite: false})
	pieces = append(pieces, &Bishop{t: bBishopt, p: backend.Position{X: 5, Y: 0}, isWhite: false})

	bknightt, err := img.LoadTexture(r, "imgs/black_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	pieces = append(pieces, &Knight{t: bknightt, p: backend.Position{X: 1, Y: 0}, isWhite: false})
	pieces = append(pieces, &Knight{t: bknightt, p: backend.Position{X: 6, Y: 0}, isWhite: false})

	brookt, err := img.LoadTexture(r, "imgs/black_rook.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Rook: %v", err)
	}
	pieces = append(pieces, &Rook{t: brookt, p: backend.Position{X: 0, Y: 0}, isWhite: false})
	pieces = append(pieces, &Rook{t: brookt, p: backend.Position{X: 7, Y: 0}, isWhite: false})

	bqueent, err := img.LoadTexture(r, "imgs/black_queen.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Queen: %v", err)
	}
	pieces = append(pieces, &Queen{t: bqueent, p: backend.Position{X: 3, Y: 0}, isWhite: false})

	bkingt, err := img.LoadTexture(r, "imgs/black_king.png")
	if err != nil {
		return nil, fmt.Errorf("could not load King: %v", err)
	}
	pieces = append(pieces, &King{t: bkingt, p: backend.Position{X: 4, Y: 0}, isWhite: false})

	// Load white pieces
	wPawnt, err := img.LoadTexture(r, "imgs/white_pawn.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Pawn: %v", err)
	}
	for i := 0; i < 8; i++ {
		pieces = append(pieces, &Pawn{t: wPawnt, p: backend.Position{X: int32(i), Y: 6}, isWhite: true})
	}

	wBishopt, err := img.LoadTexture(r, "imgs/white_bishop.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Bishop: %v", err)
	}
	pieces = append(pieces, &Bishop{t: wBishopt, p: backend.Position{X: 2, Y: 7}, isWhite: true})
	pieces = append(pieces, &Bishop{t: wBishopt, p: backend.Position{X: 5, Y: 7}, isWhite: true})

	wknightt, err := img.LoadTexture(r, "imgs/white_knight.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Knight: %v", err)
	}
	pieces = append(pieces, &Knight{t: wknightt, p: backend.Position{X: 6, Y: 7}, isWhite: true})
	pieces = append(pieces, &Knight{t: wknightt, p: backend.Position{X: 1, Y: 7}, isWhite: true})

	wrookt, err := img.LoadTexture(r, "imgs/white_rook.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Rook: %v", err)
	}
	pieces = append(pieces, &Rook{t: wrookt, p: backend.Position{X: 0, Y: 7}, isWhite: true})
	pieces = append(pieces, &Rook{t: wrookt, p: backend.Position{X: 7, Y: 7}, isWhite: true})

	wqueent, err := img.LoadTexture(r, "imgs/white_queen.png")
	if err != nil {
		return nil, fmt.Errorf("could not load Queen: %v", err)
	}
	pieces = append(pieces, &Queen{t: wqueent, p: backend.Position{X: 3, Y: 7}, isWhite: true})

	wkingt, err := img.LoadTexture(r, "imgs/white_king.png")
	if err != nil {
		return nil, fmt.Errorf("could not load King: %v", err)
	}
	pieces = append(pieces, &King{t: wkingt, p: backend.Position{X: 4, Y: 7}, isWhite: true})

	// white := createWhitePieces(pieces)
	// pieces = append(pieces, white...)

	return pieces, nil
}

func canMove(pieces []Piece, piece Piece, pos backend.Position) bool {
	if pos == piece.getPosition() {
		return false
	}
	if !piece.canMove(pos) {
		return false
	}
	switch piece.(type) {
	case *Knight:
		fmt.Println("ccc")
		if piece.canMove(pos) {
			fmt.Println("bbb")
			for _, p := range pieces {
				if pos == p.getPosition() {
					fmt.Println("aa")
					if piece.isColourWhite() == !p.isColourWhite() {
						fmt.Println("debuging")
						return true

					} else {
						fmt.Println("debuging    2")
						return false
					}
				}
			}
			return true
		}
	case *Pawn:

		if piece.canMove(pos) {
			if piece.getPosition().X != pos.X {
				for _, p := range pieces {
					if pos == p.getPosition() && p.isColourWhite() != piece.isColourWhite() {
						return true
					}
				}
				return false
			} else {
				return checkStraight(pieces, piece, pos)
			}
		}
	// case *King:
	// 	move := checkStraight(pieces, piece, pos)
	// 	// if move /*&& isCheck(pieces, piece, pos)*/ {
	// 	// return false
	// 	// }
	// 	return move
	default:
		return checkStraight(pieces, piece, pos)
	}

	return false
}

func isCheck(pieces []Piece, piece Piece, pos backend.Position) bool {
	for _, p := range pieces {
		if p.isColourWhite() != piece.isColourWhite() && canMove(pieces, p, pos) {
			return true
		}
	}
	return false
}

func checkStraight(pieces []Piece, piece Piece, pos backend.Position) bool {
	if pos.X == piece.getPosition().X && pos.Y != piece.getPosition().Y {
		dist := pos.Y - piece.getPosition().Y
		var i int32
		if dist < 0 {
			for i = -1; i >= dist; i-- {
				for _, p := range pieces {
					if p.getPosition().Y == piece.getPosition().Y+i && p.getPosition().X == piece.getPosition().X {
						if pos.Y == piece.getPosition().Y+i && pos.X == piece.getPosition().X && piece.isColourWhite() == !p.isColourWhite() {
							return true
						}
						return false
					}
				}
			}
		} else {
			for i = 1; i <= dist; i++ {
				for _, p := range pieces {
					if p.getPosition().Y == piece.getPosition().Y+i && p.getPosition().X == piece.getPosition().X {
						if pos.Y == piece.getPosition().Y+i && pos.X == piece.getPosition().X && piece.isColourWhite() == !p.isColourWhite() {
							return true
						}
						return false
					}
				}
			}
		}
	} else if pos.Y == piece.getPosition().Y && pos.X != piece.getPosition().X {
		dist := pos.X - piece.getPosition().X
		var i int32
		if dist < 0 {
			for i = -1; i >= dist; i-- {
				for _, p := range pieces {
					if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y {
						if pos.Y == piece.getPosition().Y && pos.X == piece.getPosition().X+i && piece.isColourWhite() == !p.isColourWhite() {
							return true
						}
						return false
					}
				}
			}
		} else {
			for i = 1; i <= dist; i++ {
				for _, p := range pieces {
					if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y {
						if pos.Y == piece.getPosition().Y && pos.X == piece.getPosition().X+i && piece.isColourWhite() == !p.isColourWhite() {
							return true
						}
						return false
					}
				}
			}
		}
	} else {
		return checkDiagonal(pieces, piece, pos)
	}

	return true
}

func checkDiagonal(pieces []Piece, piece Piece, pos backend.Position) bool {
	dist := pos.X - piece.getPosition().X
	if dist < 0 {
		dist = -dist
	}
	var i int32
	if pos.Y > piece.getPosition().Y && pos.X > piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y+i {
					if pos.Y == piece.getPosition().Y+i && pos.X == piece.getPosition().X+i && piece.isColourWhite() == !p.isColourWhite() {
						return true
					}
					return false
				}
			}
		}
	} else if pos.Y < piece.getPosition().Y && pos.X > piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X+i && p.getPosition().Y == piece.getPosition().Y-i {
					if pos.Y == piece.getPosition().Y-i && pos.X == piece.getPosition().X+i && piece.isColourWhite() == !p.isColourWhite() {
						return true
					}
					return false
				}
			}
		}
	} else if pos.Y > piece.getPosition().Y && pos.X < piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X-i && p.getPosition().Y == piece.getPosition().Y+i {
					if pos.Y == piece.getPosition().Y+i && pos.X == piece.getPosition().X-i && piece.isColourWhite() == !p.isColourWhite() {
						return true
					}
					return false
				}
			}
		}
	} else if pos.Y < piece.getPosition().Y && pos.X < piece.getPosition().X {
		for i = 1; i <= dist; i++ {
			for _, p := range pieces {
				if p.getPosition().X == piece.getPosition().X-i && p.getPosition().Y == piece.getPosition().Y-i {
					if pos.Y == piece.getPosition().Y-i && pos.X == piece.getPosition().X-i && piece.isColourWhite() == !p.isColourWhite() {
						return true
					}
					return false
				}
			}
		}
	}

	return true
}

func checkAfterMove(pieces []Piece, piece Piece, pos backend.Position) bool {
	copyPieces := make([]Piece, len(pieces))
	copy(copyPieces, pieces)
	if isCheck(pieces, piece, backend.Position{X: piece.getPosition().X, Y: piece.getPosition().Y}) {
		log.Println("It is check to the king")
		return true
	}

	return false
}
