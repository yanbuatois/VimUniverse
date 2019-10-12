package main

import tl "github.com/JoelOtter/termloop"

var CommandHelp Command = Command{
	Name:        "help",
	Description: "Display all available commands",
	Alias:       nil,
	MinParams:   0,
	MaxParams:   0,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		commands := AllCommands
		text := ""
		if (len(commands) <= 0) {
			text = "No command available"
		} else {
			for _, cmd := range commands {
				text += ":" + cmd.Name + " " + cmd.Description + "\n"
			}
		}
		TheGame.DisplayBigText(text)

		return false, "", tl.ColorWhite, 0
	},
}
