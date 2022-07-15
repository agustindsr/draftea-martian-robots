package main

import (
	"fmt"
	"martian-robots/pkg/domain/martianrobot"
	"os"
	"path/filepath"
)

const inputPath = "./data/input.txt"

func main() {
	f, err := readInput(inputPath)
	if err != nil {
		panic(err)
	}

	mr := martianrobot.NewService()

	robots, err := mr.InitRobots(f)
	if err != nil {
		panic(err)
	}

	for _, r := range robots {
		r.ExecuteCommands()

		fmt.Print(r.String())
	}
}

func readInput(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return "", err
	}

	filesize := fileInfo.Size()
	buffer := make([]byte, filesize)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
