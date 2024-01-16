package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sokartes/ape"
)

var runnerImage *ebiten.Image

type Game struct {
	AnimPlayer *ape.AnimationPlayer
}

func (g *Game) Update() error {
	g.AnimPlayer.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.AnimPlayer.Paused = !g.AnimPlayer.Paused

	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS += 0.1
		if g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS >= 120 {
			g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS = 120
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS -= 0.1
		if g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS <= 0.0 {
			g.AnimPlayer.Animations[g.AnimPlayer.State()].FPS = 0.0
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.AnimPlayer.SetState("jump")

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.AnimPlayer.SetState("idle")

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.AnimPlayer.SetState("run")

	}
	return nil
}

var controls string = `
W = jump state
D = run state
S = idle state
Q = Decrease FPS
E = Increase FPS
P = Play/Pause
`

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	op.GeoM.Translate(32, 0)
	screen.DrawImage(g.AnimPlayer.CurrentFrame, op)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"Controls: %s\nFrameIndex: %d\nState: %s\nFPS: %f",
			controls,
			g.AnimPlayer.CurrentFrameIndex,
			g.AnimPlayer.State(),
			g.AnimPlayer.CurrentStateFPS()),
	)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 350, 350
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	game := &Game{
		AnimPlayer: ape.NewAnimationPlayer(runnerImage),
	}
	game.AnimPlayer.NewAnimation("idle", 0, 0, 32, 32, 5)
	game.AnimPlayer.NewAnimation("run", 0, 32, 32, 32, 8)
	game.AnimPlayer.NewAnimation("jump", 0, 64, 32, 32, 4)
	game.AnimPlayer.SetState("run")

	ebiten.SetWindowSize(350, 350)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
