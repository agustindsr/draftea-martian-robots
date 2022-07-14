package entity

type Position struct {
	x           int
	y           int
	orientation Orientation
}

func NewPosition(x int, y int, orientation Orientation) Position {
	return Position{
		x:           x,
		y:           y,
		orientation: orientation,
	}
}

func (p Position) Orientation() Orientation {
	return p.orientation
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
}
