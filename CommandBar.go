package main

import (
	tl "github.com/JoelOtter/termloop"
)

type CommandBar struct {
	*tl.Rectangle
	content string
	*tl.Text
	firstTime bool
}

func NewCommandBar() *CommandBar {
	element := new(CommandBar)
	element.content = ""
	element.firstTime = true
	element.Rectangle = tl.NewRectangle(1, 1, 1, 1, tl.ColorBlack)
	element.Text = tl.NewText(1, 1, element.content, tl.ColorWhite, tl.ColorBlack)
	element.AddChar(':')

	return element
}

func (bar *CommandBar) Tick(event tl.Event) {
	if (bar.firstTime && !TheGame.IsBigTextDisplayed) {
		bar.firstTime = false
		return
	}
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyBackspace, tl.KeyBackspace2:
			bar.RemoveLastChar()
		case tl.KeyEnter:
			bar.ExecuteCommand()
		default:
			bar.AddChar(event.Ch)
		}
	}
}

func (bar *CommandBar) SetText(text string) {
	bar.content = text
	bar.Text.SetText(text)
	if text == "" && TheGame.IsBarDisplayed {
		TheGame.HideCommandBar()
	}
}

func (bar *CommandBar) AddChar(char rune) {
	bar.SetText(bar.content + string(char))
}

func (bar *CommandBar) RemoveLastChar() {
	length := len(bar.content)
	if length > 0 {
		bar.SetText(bar.content[:length-1])
	}
}

func (bar *CommandBar) ExecuteCommand() {
	command := bar.content
	bar.SetText("")
	if ok, message, color, time := TheGame.RunCommand(command); ok {
		TheGame.DisplayInfoTextWithTime(message, color, time)
	}
}

func (bar *CommandBar) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()

	bar.SetPosition(0, screenHeight-1)
	bar.SetSize(screenWidth, 1)

	bar.Rectangle.Draw(screen)
	bar.Text.Draw(screen)
}

func (bar *CommandBar) SetPosition(x, y int) {
	bar.Rectangle.SetPosition(x, y)
	bar.Text.SetPosition(x, y)
}