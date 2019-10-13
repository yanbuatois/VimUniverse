package main

import (
	tl "github.com/JoelOtter/termloop"
	"strings"
)

func (game *MyGame) RunCommand(command string) (bool, string, tl.Attr, float64) {
	displayMessage := true
	textColor := tl.ColorWhite
	time := float64(30)

	commandElements := strings.Split(command, " ")
	commandName := (commandElements[:1][0])[1:]
	commandParams := commandElements[1:]

	message := "Not a VimUniverse command  " + commandName

	implementedCommands := AllCommands

	for _, comm := range implementedCommands {
		if (comm.Name == commandName || Contains(comm.Alias, commandName)) && (!game.Ended || comm.EnableOnEnd) {
			nbParams := len(commandParams)
			if nbParams < comm.MinParams {
				return true, "Missing parameters for call " + commandName + ".", tl.ColorWhite, 30
			} else if nbParams > comm.MaxParams {
				return true, "Trailing characters", tl.ColorWhite, 30
			} else {
				return comm.Execute(commandName, commandParams, command, game)
			}
		}
	}

	//switch command {
	//case ":q", ":q!":
	//	TheGame.End()
	//case ":life":
	//	displayMessage = true
	//	message = strconv.Itoa(Player.Lives) + " Lives remaining."
	//default:
	//	displayMessage = true
	//	message = "Not a TheGame command: \"" + command + "\"."
	//}

	return displayMessage, message, textColor, time
}
