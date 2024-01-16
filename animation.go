package ape

import "github.com/hajimehoshi/ebiten/v2"

type Animation struct {
	Frames []*ebiten.Image
	FPS    float64
	Name   string
}

func (a *Animation) FrameCount() int {
	return len(a.Frames)
}

func NewAnimation(frames []*ebiten.Image, name string) *Animation {
	return &Animation{
		Frames: frames,
		Name:   name,
		FPS:    15,
	}
}
