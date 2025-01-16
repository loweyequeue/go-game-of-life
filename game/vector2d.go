package game

// A simple X,Y vector2D with built-in constraints
type Vector2D struct {
	X int8
	Y int8
}

// Normalize ensures that the vector is within the bounds of the board
func (v *Vector2D) Normalize() *Vector2D {
	if v.Y > BOARD_SIZE-1 {
		v.Y = 0
	}
	if v.Y < 0 {
		v.Y = BOARD_SIZE - 1
	}
	if v.X > BOARD_SIZE-1 {
		v.X = 0
	}
	if v.X < 0 {
		v.X = BOARD_SIZE - 1
	}
	return v
}
