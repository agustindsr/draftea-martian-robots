package martianrobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_NewService_Success(t *testing.T) {
	service := NewService()

	assert.NotNil(t, service)
}

func TestService_InitRobots_InvalidInput_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("5 3\n1 1 E")

	assert.EqualError(t, err, "invalid input")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidPlanetSize_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("5\n1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "invalid planet size: 5")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidPlanetWidth_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("5 f\n1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "strconv.Atoi: parsing \"f\": invalid syntax")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidPlanetHeight_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("s 5\n1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "strconv.Atoi: parsing \"s\": invalid syntax")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidInitialPosition_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\n1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "invalid initial position: 1 E")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidRobotXCoords_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\na E E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "strconv.Atoi: parsing \"a\": invalid syntax")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidRobotYCoords_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\n1 E E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "strconv.Atoi: parsing \"E\": invalid syntax")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidCommandLength_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\n1 2 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "the maximum number of allowed commands is 100: FRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLLFRRFLLFFRRFLL")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidCommand_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\n1 2 E\nYFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "invalid command: Y")
	assert.Nil(t, robots)
}

func TestService_InitRobots_InvalidOrientation_Error(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("3 5\n1 2 J\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.EqualError(t, err, "invalid orientation: J")
	assert.Nil(t, robots)
}

func TestService_InitRobots_Success(t *testing.T) {
	service := NewService()

	robots, err := service.InitRobots("5 3\n1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n0 3 W\nLLFFFLFLFL")

	assert.Nil(t, err)
	assert.Equal(t, 3, len(robots))
}
