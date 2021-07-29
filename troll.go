package ebitentest

import (
	"ebiten-test/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
	"math/rand"
)

type Troll struct {
	*Rectangle

	img  *ebiten.Image
}

func NewTroll() *Troll {
	return &Troll{
		Rectangle: &Rectangle{
			Width: 64,
			Height: 64,
		},
		img: assets.Troll(),
	}
}

func (t *Troll) Move(x, y float64) {
	log.Printf("Moving troll to (%f, %f)\n", x, y)
	t.PosX, t.PosY = x, y
}

func (t *Troll) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := t.img.Size()
	scaleX, scaleY := t.Rectangle.Width / float64(w), t.Rectangle.Height / float64(h)
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(t.PosX, t.PosY)

	screen.DrawImage(t.img, op)

	// Reference rect
	//ebitenutil.DrawRect(screen, t.Rectangle.PosX, t.Rectangle.PosY, t.Rectangle.Width, t.Rectangle.Height, color.NRGBA{R: 0xff, A: 0xff})
}

func (t *Troll) Update(outsideWidth, outsideHeight float64) error {
	curX, curY := ebiten.CursorPosition()
	fcurX, fcurY := float64(curX), float64(curY)
	if !t.PointCollides(fcurX, fcurY) {
		return nil
	}
	maxX, maxY := outsideWidth - t.Width, outsideHeight - t.Height
	newX, newY := rand.Float64() * outsideWidth, rand.Float64() * outsideHeight
	newX = math.Min(newX, maxX)
	newY = math.Min(newY, maxY)
	t.Move(newX, newY)
	//alert("Click the troll")
	return nil
}
