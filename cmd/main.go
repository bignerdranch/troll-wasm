package main

import (
	ebitentest "ebiten-test"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"os"
)

var _ ebiten.Game = (*Game)(nil)
type Game struct {
	outsideWidth float64
	outsideHeight float64

	troll *ebitentest.Troll
}

func (g *Game) Update() error {
	return g.troll.Update(g.outsideWidth, g.outsideHeight)
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	screen.Fill(color.NRGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xff})
	g.troll.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Cursor: (%d, %d)", x, y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.outsideWidth, g.outsideHeight = float64(outsideWidth), float64(outsideHeight)
	return outsideWidth, outsideHeight
}

func main() {
	if err := ebiten.RunGame(&Game{troll: ebitentest.NewTroll()}); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}
