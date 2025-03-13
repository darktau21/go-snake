package game

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type View struct {
	screen tcell.Screen
	quit   chan struct{}
}

func NewView() *View {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}

	return &View{
		screen: screen,
	}
}

func (v *View) Render(state GameState) {
	v.screen.Clear()

	screenWidth, screenHeight := v.screen.Size()
	boardWidth := state.GetBoardWidth()
	boardHeight := state.GetBoardHeight()

	offsetX := (screenWidth - boardWidth) / 2
	offsetY := (screenHeight - boardHeight) / 2

	for x := range boardWidth {
		v.screen.SetContent(offsetX+x, offsetY, tcell.RuneHLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		v.screen.SetContent(offsetX+x, offsetY+boardHeight-1, tcell.RuneHLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	for y := range boardHeight {
		v.screen.SetContent(offsetX, offsetY+y, tcell.RuneVLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		v.screen.SetContent(offsetX+boardWidth-1, offsetY+y, tcell.RuneVLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	v.screen.SetContent(offsetX, offsetY, tcell.RuneULCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	v.screen.SetContent(offsetX+boardWidth-1, offsetY, tcell.RuneURCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	v.screen.SetContent(offsetX, offsetY+boardHeight-1, tcell.RuneLLCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	v.screen.SetContent(offsetX+boardWidth-1, offsetY+boardHeight-1, tcell.RuneLRCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))

	for _, pos := range state.GetSnake().Body {
		v.screen.SetContent(offsetX+pos.X, offsetY+pos.Y, tcell.RuneBlock, nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	}

	v.screen.SetContent(offsetX+state.GetFood().X, offsetY+state.GetFood().Y, tcell.RuneDiamond, nil, tcell.StyleDefault.Foreground(tcell.ColorRed))

	scoreStr := fmt.Sprintf("Score: %d", state.GetScore())
	for i, r := range scoreStr {
		v.screen.SetContent(i, 0, r, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	v.screen.Show()
}

func (v *View) HandleEvents(controller *Controller) {
	for {
		ev := v.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				v.screen.Fini()
				os.Exit(0)
				return
			case tcell.KeyUp:
				controller.ChangeDirection(Up)
			case tcell.KeyDown:
				controller.ChangeDirection(Down)
			case tcell.KeyLeft:
				controller.ChangeDirection(Left)
			case tcell.KeyRight:
				controller.ChangeDirection(Right)
			}
		case *tcell.EventResize:
			v.screen.Sync()
		}
	}

}

func (v *View) Quit() {
	v.quit <- struct{}{}
}
