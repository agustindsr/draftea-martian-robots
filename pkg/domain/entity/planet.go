package entity

type Planet struct {
	point
	scents []scent
}

func NewPlanet(x int, y int) (*Planet, error) {
	point, err := NewPoint(x, y)
	if err != nil {
		return nil, err
	}

	return &Planet{
		point: *point,
	}, nil
}

func (p Planet) isScented(position Position, command Command) bool {
	for _, scent := range p.scents {
		if scent.Position.x == position.x && scent.Position.y == position.y && position.orientation == scent.Position.orientation && scent.Command == command {
			return true
		}
	}

	return false
}

func (p Planet) isInPlanet(position Position) bool {
	return position.x >= 0 && position.x <= p.point.x && position.y >= 0 && position.y <= p.point.y
}
