package martianrobot

import (
	"martian-robots/pkg/domain/entity"
	"strconv"
	"strings"
)

const (
	maxCommandLength = 100
)

type Service interface {
	InitRobots(input string) ([]entity.Robot, error)
}

type DefaultService struct {
	Service
}

func NewService() Service {
	return DefaultService{}
}

func (s DefaultService) InitRobots(input string) ([]entity.Robot, error) {
	lines := strings.Split(input, "\n")

	if len(lines) < 3 {
		return nil, errInvalidInput()
	}

	planet, err := getPlanetSize(lines[0])
	if err != nil {
		return nil, err
	}

	robots, err := getRobots(lines[1:], planet)
	if err != nil {
		return nil, err
	}

	return robots, nil
}

func getRobots(lines []string, planet *entity.Planet) ([]entity.Robot, error) {
	var robots []entity.Robot

	for i := 0; i < len(lines)-1; i += 2 {
		initialPosition, err := getInitialPosition(lines[i])
		if err != nil {
			return nil, err
		}

		commands, err := getCommands(lines[i+1])
		if err != nil {
			return nil, err
		}

		robots = append(robots, entity.NewRobot(*initialPosition, commands, planet))
	}

	return robots, nil
}

func getCommands(line string) ([]entity.Command, error) {
	commands := strings.Split(line, "")
	if len(commands) > maxCommandLength {
		return nil, errInvalidCommandLength(line)
	}

	var result []entity.Command

	for _, c := range commands {
		command, err := entity.NewCommand(c)
		if err != nil {
			return nil, err
		}

		result = append(result, *command)
	}

	return result, nil
}

func getInitialPosition(line string) (*entity.Position, error) {
	coords := strings.Split(line, " ")

	if len(coords) != 3 {
		return nil, errInvalidInitialPosition(line)
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return nil, err
	}

	orientation, err := entity.NewOrientation(coords[2])
	if err != nil {
		return nil, err
	}

	return entity.NewPosition(x, y, *orientation)
}

func getPlanetSize(line string) (*entity.Planet, error) {
	size := strings.Split(line, " ")

	if len(size) != 2 {
		return nil, errInvalidPlanetSize(line)
	}

	width, err := strconv.Atoi(size[0])
	if err != nil {
		return nil, err
	}

	height, err := strconv.Atoi(size[1])
	if err != nil {
		return nil, err
	}

	return entity.NewPlanet(width, height)
}
