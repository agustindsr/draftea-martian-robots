package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrientation_NewOrientation_West_Success(t *testing.T) {
	orientation, err := NewOrientation("W")

	assert.Nil(t, err)
	assert.Equal(t, West, *orientation)
}

func TestOrientation_NewOrientation_East_Success(t *testing.T) {
	orientation, err := NewOrientation("E")

	assert.Nil(t, err)
	assert.Equal(t, East, *orientation)
}

func TestOrientation_NewOrientation_North_Success(t *testing.T) {
	orientation, err := NewOrientation("N")

	assert.Nil(t, err)
	assert.Equal(t, North, *orientation)
}

func TestOrientation_NewOrientation_South_Success(t *testing.T) {
	orientation, err := NewOrientation("S")

	assert.Nil(t, err)
	assert.Equal(t, South, *orientation)
}

func TestOrientation_NewOrientation_Error(t *testing.T) {
	orientation, err := NewOrientation("A")

	assert.EqualError(t, err, "invalid orientation: A")
	assert.Nil(t, orientation)
}
