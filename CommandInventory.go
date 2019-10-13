package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

var CommandInventory = Command{
	Name:        "inv",
	Description: "Display the player inventory",
	Alias:       []string{"inventory"},
	MinParams:   0,
	MaxParams:   0,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		player := TheGame.Player
		text := ""
		if len(player.inventory) <= 0 {
			return true, "Your inventory is empty.", tl.ColorWhite, 30
		} else {
			for _, item := range player.inventory {
				text += string(item.Icon) + " " + item.Name

				if !item.RareItem {
					text += " (x" + strconv.Itoa(item.GetCurrentStack()) + ")"
				}

				text += ": " + item.Description + "\n"
			}
		}

		TheGame.DisplayBigText(text)

		return false, "", tl.ColorWhite, 0
	},
}
