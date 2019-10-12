package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

var CommandLife = Command{
	Name:        ":life",
	Description: "Display remaining lifes amount",
	Alias:       nil,
	MinParams:   0,
	MaxParams:   0,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		return true, "You have " + strconv.Itoa(TheGame.Player.Lives) + " remaining live(s).", tl.ColorWhite, 30
	},
}