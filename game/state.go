package game

import (
	"math/rand"
	"slices"
	"sync"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X int
	Y int
}

type Snake struct {
	Body      []Position
	Direction Direction
}

type State struct {
	Snake       Snake
	Food        Position
	BoardWidth  int
	BoardHeight int
	Score       int
	GameOver    bool
	mu          sync.Mutex
}

func (s *State) generateNewFoodPosition() Position {
	for {
		newFood := Position{
			X: rand.Intn(s.BoardWidth),
			Y: rand.Intn(s.BoardHeight),
		}
		if !slices.Contains(s.Snake.Body, newFood) {
			return newFood
		}
	}
}

func (s *State) UpdateState() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.GameOver {
		return
	}

	head := s.Snake.Body[0]
	newHead := head

	switch s.Snake.Direction {
	case Up:
		newHead.Y--
	case Down:
		newHead.Y++
	case Left:
		newHead.X--
	case Right:
		newHead.X++
	}

	if newHead.X < 0 {
		newHead.X = s.BoardWidth - 1
	} else if newHead.X >= s.BoardWidth {
		newHead.X = 0
	}

	if newHead.Y < 0 {
		newHead.Y = s.BoardHeight - 1
	} else if newHead.Y >= s.BoardHeight {
		newHead.Y = 0
	}

	if slices.Contains(s.Snake.Body, newHead) {
		s.GameOver = true
		return
	}

	s.Snake.Body = append([]Position{newHead}, s.Snake.Body...)

	if newHead == s.Food {
		s.Score++
		s.Food = s.generateNewFoodPosition()
	} else {
		s.Snake.Body = s.Snake.Body[:len(s.Snake.Body)-1]
	}
}

func (s *State) ChangeDirection(newDirection Direction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (s.Snake.Direction == Up && newDirection == Down) ||
		(s.Snake.Direction == Down && newDirection == Up) ||
		(s.Snake.Direction == Left && newDirection == Right) ||
		(s.Snake.Direction == Right && newDirection == Left) {
		return
	}

	s.Snake.Direction = newDirection
}

func (s *State) GetSnake() Snake {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Snake
}

func (s *State) GetFood() Position {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Food
}

func (s *State) GetScore() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Score
}

func (s *State) IsGameOver() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.GameOver
}

func (s *State) GetBoardWidth() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.BoardWidth
}

func (s *State) GetBoardHeight() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.BoardHeight
}
