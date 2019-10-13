package main

import (
	tl "github.com/JoelOtter/termloop"
)

var CommandRetry = Command{
	Name:        "retry",
	Description: "Restart the game.",
	Alias:       []string{"relaunch"},
	MinParams:   0,
	MaxParams:   0,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		game.Relaunch()

		return true, "The game was successfully relaunched.", tl.ColorWhite, 30
	},
}
