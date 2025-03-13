package game

import (
	"log"
	"time"
)

type Controller struct {
	state         GameState
	view          GameView
	directionChan chan Direction
}

func NewController(state GameState, view GameView) *Controller {
	return &Controller{
		state:         state,
		view:          view,
		directionChan: make(chan Direction),
	}
}

func (c *Controller) ChangeDirection(newDirection Direction) {
	c.directionChan <- newDirection
}

func (c *Controller) Run() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			c.view.HandleEvents(c)
		}
	}()

	for {
		select {
		case <-ticker.C:
			c.state.UpdateState()
			c.view.Render(c.state)
			if c.state.IsGameOver() {
				log.Println("Game Over!")
				return
			}
		case newDirection := <-c.directionChan:
			c.state.ChangeDirection(newDirection)
		}
	}
}
