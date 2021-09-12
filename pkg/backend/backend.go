
package backend

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

	return Position{x, y}
}