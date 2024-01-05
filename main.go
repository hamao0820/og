package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/*
var fs embed.FS

type Game struct {
	image *ebiten.Image
}

const (
	width  = 370
	height = 320
)

func NewGame() (*Game, error) {
	names := []string{"cry", "faint", "scare"}
	name := names[rand.Intn(len(names))]
	b, err := fs.ReadFile("images/" + name + ".png")
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	reversed := ebiten.NewImage(width, height)

	// Rotate 180 degrees
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(math.Pi)
	op.GeoM.Translate(width, height)

	reversed.DrawImage(ebiten.NewImageFromImage(img), op)

	return &Game{
		image: reversed,
	}, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.image, nil)
}

func (g *Game) Layout(w, h int) (int, int) {
	return width, height
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowClosingHandled(true)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	if err := ebiten.RunGameWithOptions(game, op); err != nil {
		log.Fatal(err)
	}
}
