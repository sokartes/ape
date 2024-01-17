package ape

import "github.com/hajimehoshi/ebiten/v2"

// Animation for AnimationPlayer
type Animation struct {
	Frames []*ebiten.Image
	FPS    float64
	Name   string
}

// FrameCount returns the frame count of the animation
func (a *Animation) FrameCount() int {
	return len(a.Frames)
}

// NewAnimation returns new Animation
func NewAnimation(frames []*ebiten.Image, name string) *Animation {
	return &Animation{
		Frames: frames,
		Name:   name,
		FPS:    15,
	}
}
