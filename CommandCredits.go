package main

import tl "github.com/JoelOtter/termloop"

var CommandCredits = Command{
	Name:        "credits",
	Description: "Display the people participating to the project.",
	Alias:       nil,
	MinParams:   0,
	MaxParams:   0,
	EnableOnEnd: true,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		text := "Yan Buatois: Development and game design.\n"
		text += "Particular thanks to JoelOtter, who developed Termloop, the Framework used by this game.\n"
		text += "Thanks for playing to this game."

		TheGame.DisplayBigText(text)
		return false, "", tl.ColorWhite, 0
	},
}
