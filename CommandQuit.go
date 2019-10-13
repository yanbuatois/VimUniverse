package main

import (
	tl "github.com/JoelOtter/termloop"
)

var CommandQuit Command = Command{
	Name:        "q",
	Description: "Quit the game or the currently open interface.",
	Alias:       []string{"q!", "quit", "quit!"},
	MinParams:   0,
	MaxParams:   0,
	EnableOnEnd: true,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		if !TheGame.IsBigTextDisplayed {
			game.End()
		} else {
			game.HideBigText()
		}
		return false, "", tl.ColorWhite, 0
	},
}
