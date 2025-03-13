package main

import "github.com/darktau21/go-snake/game"

func main() {
	// Initialize the game state
	state := &game.State{
		Snake: game.Snake{
			Body:      []game.Position{{X: 5, Y: 5}},
			Direction: game.Right,
		},
		Food:        game.Position{X: 10, Y: 10},
		BoardWidth:  20,
		BoardHeight: 20,
		Score:       0,
		GameOver:    false,
	}

	// Initialize the game view
	view := game.NewView()

	// Initialize the controller
	controller := game.NewController(state, view)

	// Run the game
	controller.Run()
}
