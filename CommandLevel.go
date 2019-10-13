package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

var CommandLevel = Command{
	Name:        "level",
	Description: "Display current level.",
	Alias:       nil,
	MinParams:   0,
	MaxParams:   0,
	EnableOnEnd: false,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		return true, "You're currently at the level " + strconv.Itoa(TheGame.CurrentLevel) + ".", tl.ColorWhite, 15
	},
}
