package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/*
var fs embed.FS

var images []*ebiten.Image

type Game struct{}

const (
	width  = 370
	height = 320
)

func init() {
	for _, name := range []string{"cry", "embarrass", "faint", "scare"} {
		b, err := fs.ReadFile("images/" + name + ".png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			log.Fatal(err)
		}

		images = append(images, ebiten.NewImageFromImage(img))
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(images[0], nil)
}

func (g *Game) Layout(w, h int) (int, int) {
	return width, height
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
