package main

import (
	"fmt"
	"github.com/maurotory/chess-golang/pkg/frontend"
	"os"
	"time"
)


func main() {
	fmt.Println("This is a chess game")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	g, err := frontend.InitGame()
	if err != nil {
		return fmt.Errorf("Could not create Game: %v", err)
	}

	err = g.RunGame()
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 30)

	if err := g.FinishGame(); err != nil {
		return fmt.Errorf("Could not finish the Game correctly: %v", err)
	}

	return nil
}

