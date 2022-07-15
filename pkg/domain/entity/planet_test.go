package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanet_NewPlanet_Success(t *testing.T) {
	planet, err := NewPlanet(3, 3)

	assert.Nil(t, err)
	assert.Equal(t, 3, planet.x)
	assert.Equal(t, 3, planet.y)
}

func TestPlanet_NewPlanet_Width_Max_Error(t *testing.T) {
	planet, err := NewPlanet(51, 3)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, planet)
}

func TestPlanet_NewPlanet_Height_Max_Error(t *testing.T) {
	planet, err := NewPlanet(2, 53)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, planet)
}
