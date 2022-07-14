package entity

import "fmt"

type Command string

const (
	MoveForward Command = "F"
	TurnLeft    Command = "L"
	TurnRight   Command = "R"
)

var availableCommands = map[string]Command{
	"F": MoveForward,
	"L": TurnLeft,
	"R": TurnRight,
}

func NewCommand(c string) Command {
	if command, ok := availableCommands[c]; ok {
		return command
	}

	panic(fmt.Sprintf("invalid Command %s", c))
}
