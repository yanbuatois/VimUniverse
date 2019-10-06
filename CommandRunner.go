package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

func (game *MyGame) RunCommand(command string) (bool, string, tl.Attr, float64) {
	displayMessage := false
	message := ""
	textColor := tl.ColorWhite
	time := float64(30)

	switch command {
	case ":q", ":q!":
		game.End()
	case ":life":
		displayMessage = true
		message = strconv.Itoa(game.Player.lives) + " lives remaining."
	default:
		displayMessage = true
		message = "Not a game command: \"" + command + "\"."
	}

	return displayMessage, message, textColor, time
}
