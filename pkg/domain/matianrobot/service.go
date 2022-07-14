package matianrobot

import (
	"bufio"
	"errors"
	"log"
	"martian-robots/pkg/domain/entity"
	"os"
	"strconv"
	"strings"
)

type service struct {
	Robots []entity.Robot
}

func NewService(input *os.File) (*service, error) {
	var lines []string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(lines) < 2 {
		return nil, errors.New("invalid input")
	}

	planet := entity.NewPlanet(5, 3)

	robots, err := getRobots(lines[1:], planet)
	if err != nil {
		return nil, err
	}

	return &service{Robots: robots}, nil
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
	if len(commands) > 100 {
		return nil, errInvalidCommandLength()
	}

	var result []entity.Command

	for _, c := range commands {
		result = append(result, entity.NewCommand(c))
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
		return nil, errInvalidInitialPosition(line)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return nil, err
	}

	orientation := entity.NewOrientation(coords[2])

	position := entity.NewPosition(x, y, orientation)

	return &position, nil
}
