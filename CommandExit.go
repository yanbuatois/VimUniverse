package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

var CommandExit = Command{
	Name:        "exit",
	Description: "Exits by the door.",
	Alias:       nil,
	MinParams:   0,
	MaxParams:   0,
	EnableOnEnd: false,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		upLevel := func() (bool, string, tl.Attr, float64) {
			TheGame.CurrentLevel++
			TheGame.CallParser()
			return true, "Congratulations. You're now at the level " + strconv.Itoa(TheGame.CurrentLevel) + "! However, you lost all of your belongings while running...", tl.ColorWhite, 0
		}
		player := TheGame.Player
		switch player.onDoor {
		case OnLockedDoor:
			ok := player.UseItem("key")
			if !ok {
				return true, "You need to have a key to open this door.", tl.ColorWhite, 0
			} else {
				return upLevel()
			}
		case OnOpenDoor:
			return upLevel()
		default:
			return true, "You're not on an exit door.", tl.ColorWhite, 0
		}
	},
}
