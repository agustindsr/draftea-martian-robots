package entity

import "fmt"

type Robot struct {
	isLost   bool
	position Position
	commands []Command
	planet   *Planet
}

func NewRobot(position Position, commands []Command, planet *Planet) Robot {
	return Robot{
		commands: commands,
		position: position,
		planet:   planet,
	}
}

func (r Robot) IsLost() bool {
	return r.isLost
}

func (r *Robot) ExecuteCommands() {
	for _, c := range r.commands {
		if r.planet.isScented(r.position, c) {
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
		}

		if !r.planet.isInPlanet(np) {
			r.planet.scents = append(r.planet.scents, scent{r.position, c})
			r.isLost = true
			break
		}

		r.position = np
	}
}

func (r Robot) Position() Position {
	return r.position
}

func (r Robot) moveForward() Position {
	switch r.position.orientation {
	case North:
		r.position.y++
	case East:
		r.position.x++
	case South:
		r.position.y--
	case West:
		r.position.x--
	}

	return r.position
}

func (r Robot) turnLeft() Position {
	switch r.position.orientation {
	case North:
		r.position.orientation = West
	case East:
		r.position.orientation = North
	case South:
		r.position.orientation = East
	case West:
		r.position.orientation = South
	}

	return r.position
}

func (r *Robot) turnRight() Position {
	switch r.position.orientation {
	case North:
		r.position.orientation = East
	case East:
		r.position.orientation = South
	case South:
		r.position.orientation = West
	case West:
		r.position.orientation = North
	}

	return r.position
}

func (r Robot) String() string {
	if r.IsLost() {
		return fmt.Sprintf("%s LOST \n", r.position.String())
	} else {
		return fmt.Sprintf("%s \n", r.position.String())
	}
}
