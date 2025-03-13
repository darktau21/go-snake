package game

import (
	"log"
	"time"
)

type Controller struct {
	state         GameState
	view          GameView
	directionChan chan Direction
	quitChan      chan struct{}
}

func NewController(state GameState, view GameView) *Controller {
	return &Controller{
		state:         state,
		view:          view,
		directionChan: make(chan Direction),
		quitChan:      make(chan struct{}),
	}
}

func (c *Controller) ChangeDirection(newDirection Direction) {
	c.directionChan <- newDirection
}

func (c *Controller) Exit() {
	c.quitChan <- struct{}{}
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
				c.Exit()
				return
			}
		case newDirection := <-c.directionChan:
			c.state.ChangeDirection(newDirection)
		case <-c.quitChan:
			c.view.Quit()
			return
		}
	}
}
