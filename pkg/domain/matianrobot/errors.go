package matianrobot

import "fmt"

var (
	errInvalidCommand = func(command string) error {
		return fmt.Errorf("invalid command %s", command)
	}

	errInvalidCommandLength = func() error {
		return fmt.Errorf("the maximum number of allowed commands is 100")
	}

	errInvalidInitialPosition = func(initialPosition string) error {
		return fmt.Errorf("invalid initial position %s", initialPosition)
	}
)
