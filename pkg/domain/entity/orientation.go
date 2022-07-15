package entity

import (
	"fmt"
)

type Orientation string

const (
	North Orientation = "N"
	East  Orientation = "E"
	South Orientation = "S"
	West  Orientation = "W"
)

var availableOrientations = map[string]Orientation{
	"N": North,
	"E": East,
	"S": South,
	"W": West,
}

func NewOrientation(orientation string) (*Orientation, error) {
	if o, ok := availableOrientations[orientation]; ok {
		return &o, nil
	}

	return nil, fmt.Errorf("invalid orientation: %s", orientation)
}
