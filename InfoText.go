package main

import (
	tl "github.com/JoelOtter/termloop"
)

type InfoText struct {
	*tl.Text
	time float64
}

func (infoText *InfoText) Draw(screen *tl.Screen) {
	if infoText.time > 0 {
		infoText.time -= screen.TimeDelta()
		if infoText.time <= 0 {
			game.HideInfoText()
		}
	}

	_, screenHeight := screen.Size()

	var offset int
	if game.isBarDisplayed {
		offset = 2
	} else {
		offset = 1
	}

	infoText.SetPosition(0, screenHeight-offset)

	infoText.Text.Draw(screen)
}
