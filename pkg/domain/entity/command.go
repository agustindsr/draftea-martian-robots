package entity

import (
	"fmt"
)

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

func NewCommand(command string) (*Command, error) {
	if c, ok := availableCommands[command]; ok {
		return &c, nil
	}

	return nil, fmt.Errorf("invalid command: %s", command)
}
