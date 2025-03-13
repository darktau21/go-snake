package main

import "github.com/darktau21/go-snake/game"

func main() {
	state := &game.State{
		Snake: game.Snake{
			Body:      []game.Position{{X: 5, Y: 5}},
			Direction: game.Right,
		},
		Food:        game.Position{X: 10, Y: 10},
		BoardWidth:  40,
		BoardHeight: 20,
		Score:       0,
		GameOver:    false,
	}

	view := game.NewView()

	controller := game.NewController(state, view)

	controller.Run()
}
