package entity

import "fmt"

const (
	maxCoordsValue = 50
)

type point struct {
	x int
	y int
}

func NewPoint(x int, y int) (*point, error) {
	if x > maxCoordsValue || y > maxCoordsValue {
		return nil, fmt.Errorf("the maximum coordinate value is %d", maxCoordsValue)
	}

	return &point{x, y}, nil
}
