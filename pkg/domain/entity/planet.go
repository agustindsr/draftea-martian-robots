package entity

type Planet struct {
	scents []scent
	width  int
	height int
}

func NewPlanet(width int, height int) *Planet {
	return &Planet{
		width:  width,
		height: height,
	}
}

func (p Planet) IsScented(Position Position, command Command) bool {
	for _, scent := range p.scents {
		if scent.Position.x == Position.x && scent.Position.y == Position.y && Position.orientation == scent.Position.orientation && scent.Command == command {
			return true
		}
	}

	return false
}
