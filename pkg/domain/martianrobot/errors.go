package martianrobot

import "fmt"

var (
	errInvalidInput = func() error {
		return fmt.Errorf("invalid input")
	}

	errInvalidCommandLength = func(commands string) error {
		return fmt.Errorf("the maximum number of allowed commands is %d: %s", maxCommandLength, commands)
	}

	errInvalidInitialPosition = func(initialPosition string) error {
		return fmt.Errorf("invalid initial position: %s", initialPosition)
	}

	errInvalidPlanetSize = func(planetSize string) error {
		return fmt.Errorf("invalid planet size: %s", planetSize)
	}
)
