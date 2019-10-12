package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

var CommandPick Command = Command{
	Name:        "pick",
	Description: "Pick the item on which you are.",
	Alias:       []string{},
	MinParams:   0,
	MaxParams:   0,
	Execute: func(call string, params []string, fullMessage string, game *MyGame) (b bool, s string, attr tl.Attr, f float64) {
		player := game.Player
		if player.onItem == nil {
			return true, "You're not on an item.", tl.ColorWhite, 15
 		} else {
 			alreadyPicked := false
 			item := player.onItem
 			playerItem := item.Item
 			if (!item.RareItem) {
				for index, uIt := range player.inventory {
					if (uIt.Id == item.Id) {
						player.inventory[index].currentStack += item.currentStack
						if (player.inventory[index].currentStack > player.inventory[index].maxStack) {
							player.inventory[index].currentStack = player.inventory[index].maxStack
						}
						alreadyPicked = true
						playerItem = player.inventory[index]
					}
				}
			}

			if (!alreadyPicked) {
				player.inventory = append(player.inventory, item.Item)
			}
 			TheGame.Level.RemoveEntity(item)
 			item.Item.PickItem()
 			player.onItem = nil
 			text := "You picked up "
 			if (!item.RareItem) {
 				text += strconv.Itoa(item.GetCurrentStack()) + "x "
			}
			text +=  item.Name + ". It goes now in your inventory."
			if (!item.RareItem) {
				text += " You now have " + strconv.Itoa(playerItem.GetCurrentStack()) + " items of this kind."
			}
 			return true, text, tl.ColorWhite, 10
		}
	},
}
