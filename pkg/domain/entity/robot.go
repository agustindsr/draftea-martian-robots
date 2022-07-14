package entity

type Robot interface {
	ExecuteCommands()
	Position() Position
	IsLost() bool
}

type robot struct {
	isLost   bool
	position Position
	commands []Command
	planet   *Planet
}

func NewRobot(position Position, commands []Command, planet *Planet) Robot {
	return &robot{
		commands: commands,
		position: position,
		planet:   planet,
	}
}

func (r robot) IsLost() bool {
	return r.isLost
}

func (r *robot) ExecuteCommands() {
	for _, c := range r.commands {
		if r.planet.IsScented(r.position, c) {
			continue
		}

		var np Position

		switch c {
		case MoveForward:
			np = r.moveForward()
		case TurnLeft:
			np = r.turnLeft()
		case TurnRight:
			np = r.turnRight()

		default:
			panic("unknown Command")
		}

		if !r.isInPlanet(np) {
			r.planet.scents = append(r.planet.scents, scent{r.position, c})
			r.isLost = true
			break
		}

		r.position = np

	}
}

func (r robot) Position() Position {
	return r.position
}

func (r robot) moveForward() Position {
	switch r.position.orientation {
	case North:
		r.position.y++
	case East:
		r.position.x++
	case South:
		r.position.y--
	case West:
		r.position.x--
	default:
		panic("unknown orientation")
	}

	return r.position
}

func (r robot) turnLeft() Position {
	switch r.position.orientation {
	case North:
		r.position.orientation = West
	case East:
		r.position.orientation = North
	case South:
		r.position.orientation = East
	case West:
		r.position.orientation = South
	default:
		panic("unknown orientation")
	}

	return r.position
}

func (r *robot) turnRight() Position {
	switch r.position.orientation {
	case North:
		r.position.orientation = East
	case East:
		r.position.orientation = South
	case South:
		r.position.orientation = West
	case West:
		r.position.orientation = North
	default:
		panic("unknown orientation")
	}

	return r.position
}

func (r *robot) isInPlanet(Position Position) bool {
	return Position.x >= 0 && Position.x <= r.planet.width && r.position.y >= 0 && Position.y <= r.planet.height
}
