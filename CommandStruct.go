package main

import (tl "github.com/JoelOtter/termloop")

type Command struct {
	Name string
	Description string
	Alias []string
	MinParams int
	MaxParams int
	Execute func(call string, params []string, fullMessage string, game *MyGame) (bool, string, tl.Attr, float64)
}

var AllCommands []Command

func init() {
	AllCommands = []Command{
		CommandQuit,
		CommandPick,
		CommandHelp,
		CommandLife,
		CommandInventory,
	}
}