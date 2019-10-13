package main

import (
	tl "github.com/JoelOtter/termloop"
)

func main() {
	TheGame = MyGame{Game: tl.NewGame(), CurrentLevel: 1}

	//level := tl.NewBaseLevel(tl.Cell{
	//	Fg: tl.ColorBlack,
	//	Bg: tl.ColorBlack,
	//	Ch: ' ',
	//})
	//
	//level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))
	//
	//game.Screen().SetLevel(level)
	TheGame.SetDebugOn(true)
	player := NewPlayer()
	TheGame.SetPlayer(player)
	TheGame.CallParser()
	TheGame.Screen().SetFps(60)
	if !TheGame.DebugOn() {
		TheGame.SetEndKey(tl.KeyEsc)
	}

	TheGame.Log("run")

	TheGame.Start()
}
