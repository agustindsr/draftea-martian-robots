package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition_NewPosition_Success(t *testing.T) {
	position, err := NewPosition(3, 3, North)

	assert.Nil(t, err)
	assert.Equal(t, 3, position.x)
	assert.Equal(t, 3, position.y)
	assert.Equal(t, North, position.orientation)
}

func TestPosition_NewPosition_X_Max_Error(t *testing.T) {
	position, err := NewPosition(51, 3, North)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, position)
}

func TestPosition_NewPosition_Y_Max_Error(t *testing.T) {
	position, err := NewPosition(2, 53, North)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, position)
}

func TestPosition_String(t *testing.T) {
	position := Position{
		point:       point{x: 1, y: 2},
		orientation: North,
	}

	expected := "1 2 N"

	assert.Equal(t, expected, position.String())
}
