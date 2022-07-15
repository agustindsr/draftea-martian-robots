package entity

import (
	"fmt"
)

type Position struct {
	point
	orientation Orientation
}

func NewPosition(x int, y int, orientation Orientation) (*Position, error) {
	point, err := NewPoint(x, y)
	if err != nil {
		return nil, err
	}

	return &Position{
		point:       *point,
		orientation: orientation,
	}, nil
}

func (p Position) String() string {
	return fmt.Sprintf("%d %d %s", p.x, p.y, p.orientation)
}
