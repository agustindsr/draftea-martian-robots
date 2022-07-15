package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint_NewPoint_Success(t *testing.T) {
	point, err := NewPoint(3, 3)

	assert.Nil(t, err)
	assert.Equal(t, 3, point.x)
	assert.Equal(t, 3, point.y)
}

func TestPoint_NewPoint_X_Max_Error(t *testing.T) {
	point, err := NewPoint(51, 3)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, point)
}

func TestPoint_NewPoint_Y_Max_Error(t *testing.T) {
	point, err := NewPoint(2, 53)

	assert.EqualError(t, err, "the maximum coordinate value is 50")
	assert.Nil(t, point)
}
