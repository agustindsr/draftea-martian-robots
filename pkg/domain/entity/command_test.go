package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_NewCommand_Left_Success(t *testing.T) {
	command, err := NewCommand("L")

	assert.Nil(t, err)
	assert.Equal(t, TurnLeft, *command)
}

func TestCommand_NewCommand_Right_Success(t *testing.T) {
	command, err := NewCommand("R")

	assert.Nil(t, err)
	assert.Equal(t, TurnRight, *command)
}

func TestCommand_NewCommand_Forward_Success(t *testing.T) {
	command, err := NewCommand("F")

	assert.Nil(t, err)
	assert.Equal(t, MoveForward, *command)
}

func TestCommand_NewCommand_Error(t *testing.T) {
	command, err := NewCommand("A")

	assert.EqualError(t, err, "invalid command: A")
	assert.Nil(t, command)
}
