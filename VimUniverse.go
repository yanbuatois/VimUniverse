package main

import (
	tl "github.com/JoelOtter/termloop"
)

var game (MyGame)

func main() {
	game = MyGame{Game: tl.NewGame(), currentLevel: 1}

	//level := tl.NewBaseLevel(tl.Cell{
	//	Fg: tl.ColorBlack,
	//	Bg: tl.ColorBlack,
	//	Ch: ' ',
	//})
	//
	//level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))
	//
	//game.Screen().SetLevel(level)
	player := NewPlayer()
	game.SetPlayer(player)
	game.CallParser()
	game.Screen().SetFps(60)
	game.SetDebugOn(true)
	game.SetEndKey(tl.KeyEsc)

	game.Log("run")

	game.Start()
}
