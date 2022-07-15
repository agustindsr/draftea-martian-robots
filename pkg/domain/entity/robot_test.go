package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobot_NewRobot(t *testing.T) {
	position := Position{
		point: point{
			x: 1,
			y: 2,
		},
		orientation: North,
	}

	commands := []Command{MoveForward, TurnLeft, TurnRight}
	planet := Planet{
		point: point{
			x: 10,
			y: 8,
		},
	}

	robot := NewRobot(position, commands, &planet)

	assert.Equal(t, 1, robot.position.x)
	assert.Equal(t, 2, robot.position.y)
	assert.Equal(t, North, robot.position.orientation)
	assert.Equal(t, []Command{MoveForward, TurnLeft, TurnRight}, robot.commands)
	assert.Equal(t, []Command{MoveForward, TurnLeft, TurnRight}, robot.commands)
	assert.Equal(t, 10, robot.planet.x)
	assert.Equal(t, 8, robot.planet.y)
}

func TestRobot_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		initial  Orientation
		expected Orientation
	}{
		{
			name:     "when facing north and turning left the facing should be west",
			initial:  North,
			expected: West,
		},
		{
			name:     "when facing east and turning left the facing should be north",
			initial:  East,
			expected: North,
		},
		{
			name:     "when facing south and turning left the facing should be east",
			initial:  South,
			expected: East,
		},
		{
			name:     "when facing west and turning left the facing should be south",
			initial:  West,
			expected: South,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			robot := Robot{
				position: Position{
					orientation: test.initial,
				},
			}
			position := robot.turnLeft()
			assert.Equal(t, test.expected, position.orientation)
		})
	}
}

func TestRobot_TurnRight(t *testing.T) {
	tests := []struct {
		name     string
		initial  Orientation
		expected Orientation
	}{
		{
			name:     "when facing north and turning right the facing should be east",
			initial:  North,
			expected: East,
		},
		{
			name:     "when facing east and turning right the facing should be south",
			initial:  East,
			expected: South,
		},
		{
			name:     "when facing south and turning right the facing should be west",
			initial:  South,
			expected: West,
		},
		{
			name:     "when facing west and turning right the facing should be north",
			initial:  West,
			expected: North,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			robot := Robot{
				position: Position{
					orientation: test.initial,
				},
			}
			position := robot.turnRight()
			assert.Equal(t, test.expected, position.orientation)
		})
	}
}

func TestRobot_MoveForward(t *testing.T) {
	tests := []struct {
		name     string
		initial  Position
		expected Position
	}{
		{
			name:     "when the orientation is north and i move forward, the y coordinate increases",
			initial:  Position{point: point{x: 0, y: 0}, orientation: North},
			expected: Position{point: point{x: 0, y: 1}, orientation: North},
		},
		{
			name:     "when the orientation is east and i move forward, the x coordinate increases",
			initial:  Position{point: point{x: 0, y: 0}, orientation: East},
			expected: Position{point: point{x: 1, y: 0}, orientation: East},
		},
		{
			name:     "when the orientation is south and i move forward, the y coordinate decrease",
			initial:  Position{point: point{x: 0, y: 1}, orientation: South},
			expected: Position{point: point{x: 0, y: 0}, orientation: South},
		},
		{
			name:     "when the orientation is west and i move forward, the x coordinate increases",
			initial:  Position{point: point{x: 1, y: 0}, orientation: West},
			expected: Position{point: point{x: 0, y: 0}, orientation: West},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			robot := Robot{
				position: test.initial,
			}

			position := robot.moveForward()
			assert.Equal(t, test.expected, position)
		})
	}
}

func TestRobot_Position(t *testing.T) {
	robot := Robot{
		position: Position{
			point: point{
				x: 1,
				y: 2,
			},
			orientation: North,
		},
	}

	expectedPosition := Position{
		point: point{
			x: 1,
			y: 2,
		},
		orientation: North,
	}

	expected := robot.Position()
	assert.Equal(t, expected, expectedPosition)
}

func TestRobot_IsLost(t *testing.T) {
	robot := Robot{
		isLost: true,
	}

	expected := robot.IsLost()

	assert.Equal(t, expected, robot.isLost)
}

func TestRobot_ExecuteCommands(t *testing.T) {
	tests := []struct {
		commands         []Command
		initialPosition  Position
		scents           []scent
		expectedPosition Position
		expectedLost     bool
		expectedScents   []scent
	}{
		{
			initialPosition: Position{point: point{1, 1}, orientation: East},
			commands:        []Command{TurnRight, MoveForward, TurnRight, MoveForward, TurnRight, MoveForward, TurnRight, MoveForward},
			expectedPosition: Position{
				point: point{
					x: 1,
					y: 1,
				},
				orientation: East,
			},
			expectedLost:   false,
			expectedScents: nil,
		},
		{
			initialPosition: Position{point: point{3, 2}, orientation: North},
			commands:        []Command{MoveForward, TurnRight, TurnRight, MoveForward, TurnLeft, TurnLeft, MoveForward, MoveForward, TurnRight, TurnRight, MoveForward, TurnLeft, TurnLeft},
			expectedPosition: Position{
				point: point{
					x: 3,
					y: 3,
				},
				orientation: North,
			},
			expectedScents: []scent{
				{Position{point: point{x: 3, y: 3}, orientation: North}, MoveForward},
			},
			expectedLost: true,
		},
		{
			initialPosition: Position{point: point{0, 3}, orientation: West},
			commands:        []Command{TurnLeft, TurnLeft, MoveForward, MoveForward, MoveForward, TurnLeft, MoveForward, TurnLeft, MoveForward, TurnLeft},
			expectedPosition: Position{
				point: point{
					x: 2,
					y: 3,
				},
				orientation: South,
			},
			scents: []scent{
				{Position{point: point{x: 3, y: 3}, orientation: North}, MoveForward},
			},
			expectedLost: false,
			expectedScents: []scent{
				{Position{point: point{x: 3, y: 3}, orientation: North}, MoveForward},
			},
		},
	}

	for _, test := range tests {
		t.Run("Execute command tests", func(t *testing.T) {
			robot := Robot{
				planet:   &Planet{point: point{x: 5, y: 3}, scents: test.scents},
				position: test.initialPosition,
				commands: test.commands,
			}

			robot.ExecuteCommands()

			assert.Equal(t, test.expectedPosition, robot.position)
			assert.Equal(t, test.expectedLost, robot.isLost)
			assert.Equal(t, test.expectedScents, robot.planet.scents)
		})
	}
}

func TestRobot_String(t *testing.T) {
	robot := Robot{
		position: Position{
			point: point{
				x: 1,
				y: 2,
			},
			orientation: North,
		},
		isLost: false,
	}

	expected := "1 2 N \n"

	assert.Equal(t, expected, robot.String())
}

func TestRobot_RobotLost_String(t *testing.T) {
	robot := Robot{
		position: Position{
			point: point{
				x: 1,
				y: 2,
			},
			orientation: North,
		},
		isLost: true,
	}

	expected := "1 2 N LOST \n"

	assert.Equal(t, expected, robot.String())
}
