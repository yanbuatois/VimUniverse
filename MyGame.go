package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"os"
)

var TheGame (MyGame)

type MyGame struct {
	*tl.Game
	*CommandBar
	IsBarDisplayed bool
	*InfoText
	*Player
	CurrentLevel       int
	IsBigTextDisplayed bool
	*BigText
	Ended   bool
	endText *EndBigText
}

func (game *MyGame) DisplayCommandBar() {
	if game.IsBarDisplayed {
		game.HideCommandBar()
	}
	game.HideInfoText()
	game.CommandBar = NewCommandBar()
	game.Screen().AddEntity(game.CommandBar)
	game.IsBarDisplayed = true
}

func (game *MyGame) HideCommandBar() {
	game.IsBarDisplayed = false
	if game.CommandBar != nil {
		game.Screen().RemoveEntity(game.CommandBar)
		game.CommandBar.SetText("")
		game.CommandBar = nil
	}
}

func (game *MyGame) DisplayInfoText(text string, textColor tl.Attr) {
	game.DisplayInfoTextWithTime(text, textColor, -1)
}

func (game *MyGame) DisplayInfoTextWithTime(text string, textColor tl.Attr, time float64) {
	if game.InfoText != nil {
		game.HideInfoText()
	}
	game.InfoText = new(InfoText)
	game.InfoText.Text = tl.NewText(1, 1, text, textColor, tl.ColorBlack)
	game.InfoText.time = time
	game.Screen().AddEntity(game.InfoText)
}

func (g *MyGame) DisplayBigText(text string) {
	if len(text) <= 0 {
		return
	}
	if g.IsBigTextDisplayed {
		g.HideBigText()
	}
	g.IsBigTextDisplayed = true
	g.BigText = NewBigText(text)
	g.Screen().AddEntity(g.BigText)
}

func (g *MyGame) HideBigText() {
	g.IsBigTextDisplayed = false
	if g.BigText != nil {
		g.Screen().RemoveEntity(g.BigText)
		g.BigText = nil
	}
}

func (game *MyGame) End() {
	os.Exit(0)
}

func (game *MyGame) HideInfoText() {
	if game.InfoText != nil {
		game.Screen().RemoveEntity(game.InfoText)
		game.InfoText = nil
	}
}

func (game *MyGame) SetPlayer(player *Player) {
	game.Player = player
}

func (game *MyGame) CallParser() {
	success := LoadLevel(game.CurrentLevel)

	if !success && game.CurrentLevel > 1 {
		game.Win()
	} else if !success {
		fmt.Print("No level files found.")
		game.End()
	}
}

func (game *MyGame) NextLevel() {
	game.CurrentLevel++
	game.Player.inventory = [](*Item){}
	game.CallParser()
}

func (game *MyGame) Win() {
	game.endGame()
	game.endText = NewEndBigText(WonText)
	game.Screen().AddEntity(game.endText)
}

func (game *MyGame) Lose() {
	game.endGame()
	game.endText = NewEndBigText(LostText)
	game.Screen().AddEntity(game.endText)
}

func (game *MyGame) endGame() {
	game.Ended = true
	game.HideBigText()
	game.HideInfoText()
	game.HideCommandBar()
}

func (game *MyGame) Relaunch() {
	game.HideInfoText()
	game.HideBigText()
	game.HideCommandBar()
	game.Ended = false
	game.CurrentLevel = 1
	game.Player = NewPlayer()
	game.Screen().RemoveEntity(game.endText)
	game.endText = nil
	game.CallParser()
}
