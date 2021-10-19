package backend

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

const (
	FPS        int32 = 60
	SCALE      int32 = 3
	PIECE_SIZE int32 = 20
)

type Position struct {
	X, Y int32
}

func GetCell(X int32, Y int32) Position {
	x := (X - 2*SCALE) / (22 * SCALE)
	y := (Y - 2*SCALE) / (22 * SCALE)
	fmt.Println(x)
	fmt.Println(y)

	return Position{x, y}
}

func (p *Position) GetCordinates() (int32, int32) {
	return (3 + p.X*22) * 3, (3 + p.Y*22) * 3
}

type Game struct {
	WhitePlayer *Player
	BlackPlayer *Player
	Mu          sync.RWMutex
	WhiteTurn   bool
}

type Player struct {
	Name          string
	ChangeChannel chan Move
	Id            uuid.UUID
	IsWhite       bool
}

type Move struct {
	InitialX int
	InitialY int
	FinalX   int
	FinalY   int
}

func (pos *Position) Simmetry() *Position {

	return nil

}
