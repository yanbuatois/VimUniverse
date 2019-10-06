package main

import (
	tl "github.com/JoelOtter/termloop"
	"os"
)

type MyGame struct {
	*tl.Game
	*CommandBar
	isBarDisplayed bool
	*InfoText
	*Player
	currentLevel int
}

func (game *MyGame) DisplayCommandBar() {
	game.HideInfoText()
	if game.CommandBar == nil {
		game.CommandBar = NewCommandBar()
	}
	game.Screen().AddEntity(game.CommandBar)
	game.isBarDisplayed = true
}

func (game *MyGame) HideCommandBar() {
	game.Screen().RemoveEntity(game.CommandBar)
	game.isBarDisplayed = false
	game.CommandBar.SetText("")
}

func (game *MyGame) DisplayInfoText(text string, textColor tl.Attr) {
	game.DisplayInfoTextWithTime(text, textColor, -1)
}

func (game *MyGame) DisplayInfoTextWithTime(text string, textColor tl.Attr, time float64) {
	game.InfoText = new(InfoText)
	game.InfoText.Text = tl.NewText(1, 1, text, textColor, tl.ColorBlack)
	game.InfoText.time = time
	game.Screen().AddEntity(game.InfoText)
}

func (game *MyGame) End() {
	os.Exit(0)
}

func (game *MyGame) HideInfoText() {
	game.Screen().RemoveEntity(game.InfoText)
}

func (game *MyGame) SetPlayer(player *Player) {
	game.Player = player
}

func (game *MyGame) CallParser() {
	LoadLevel(game.currentLevel)
}
