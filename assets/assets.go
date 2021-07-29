package assets

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

func cacheFunc(name string, data []byte) func() *ebiten.Image {
	var cache *ebiten.Image

	return func() *ebiten.Image {
		if cache == nil {
			pngImg, err := png.Decode(bytes.NewReader(data))
			if err != nil {
				panic(fmt.Errorf("assets: unable to load '%s': %w", name, err))
			}
			cache = ebiten.NewImageFromImage(pngImg)
		}
		return cache
	}
}

//go:embed troll.png
var trollBytes []byte
var Troll = cacheFunc("troll.png", trollBytes)
