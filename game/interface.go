package game

type GameState interface {
	UpdateState()
	ChangeDirection(newDirection Direction)
	GetSnake() Snake
	GetFood() Position
	GetScore() int
	IsGameOver() bool
	GetBoardWidth() int
	GetBoardHeight() int
}

type GameView interface {
	Render(state GameState)
	HandleEvents(controller *Controller)
	Quit()
}
