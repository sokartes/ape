package ape

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationPlayer struct {
	SpriteSheet, CurrentFrame *ebiten.Image
	Animations                map[string]*Animation
	Paused                    bool
	CurrentFrameIndex         int

	currentState string
	tick         float64
}

func (ap *AnimationPlayer) NewAnimation(stateName string, x, y, w, h, frameCount int) *Animation {
	subImages := []*ebiten.Image{}
	frameRect := image.Rect(x, y, x+w, y+h)
	for i := 0; i < frameCount; i++ {
		subImages = append(subImages, ap.SpriteSheet.SubImage(frameRect).(*ebiten.Image))
		frameRect.Min.X += w
		frameRect.Max.X += w
	}
	anim := &Animation{
		FPS:    15.0,
		Frames: subImages,
		Name:   stateName,
	}
	ap.currentState = stateName
	ap.Animations[stateName] = anim
	return anim
}

func (ap *AnimationPlayer) SetFPS(fps float64) {
	for _, anim := range ap.Animations {
		anim.FPS = fps
	}
}

func (ap *AnimationPlayer) AddAnimation(a *Animation) {
	ap.Animations[a.Name] = a
}

func (ap *AnimationPlayer) State() string {
	return ap.currentState
}

func (ap *AnimationPlayer) CurrentStateFPS() float64 {
	return ap.Animations[ap.State()].FPS
}

func (ap *AnimationPlayer) SetState(state string) {
	if ap.currentState != state {
		ap.currentState = state
		ap.tick = 0
		ap.CurrentFrameIndex = 0
	}
}

func (ap *AnimationPlayer) PauseAtFrame(frameIndex int) {
	ap.Paused = true
	ap.CurrentFrameIndex = frameIndex
}

func (ap *AnimationPlayer) Update() {
	if !ap.Paused {
		ap.tick += ap.Animations[ap.currentState].FPS / 60.0
		ap.CurrentFrameIndex = int(math.Floor(ap.tick))
		if ap.CurrentFrameIndex >= ap.Animations[ap.currentState].FrameCount() {
			ap.tick = 0
			ap.CurrentFrameIndex = 0
		}
	}

	// update image
	ap.CurrentFrame = ap.Animations[ap.currentState].Frames[ap.CurrentFrameIndex]
}

func NewAnimationPlayer(spriteSheet *ebiten.Image) *AnimationPlayer {
	return &AnimationPlayer{
		SpriteSheet:       spriteSheet,
		Paused:            false,
		Animations:        make(map[string]*Animation),
		CurrentFrameIndex: 0,
	}

}
