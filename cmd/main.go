package main

import (
	"fmt"
	"log"
	"martian-robots/pkg/domain/matianrobot"
	"os"
	"path/filepath"
)

const inputPath = "./data/input.txt"

func main() {
	path, err := filepath.Abs(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mr, err := matianrobot.NewService(f)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range mr.Robots {
		r.ExecuteCommands()

		fp := r.Position()

		msg := fmt.Sprintf("%d %d %s", fp.X(), fp.Y(), fp.Orientation())

		if r.IsLost() {
			log.Printf("%s LOST \n", msg)
		} else {
			log.Printf("%s \n", msg)
		}
	}
}
