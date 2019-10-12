package main

import (tl "github.com/JoelOtter/termloop"
	"strings"
)

type BigText struct {
	Content string
	Texts []*tl.Text
	*tl.Rectangle
	XOffset int
	YOffset int
	maxTextLength int
}

func NewBigText (text string) *BigText {
	retour := new (BigText)
	retour.Content = text
	retour.XOffset = 0
	retour.YOffset = 0

	width, height := TheGame.Screen().Size()

	retour.Rectangle = tl.NewRectangle(0,0, width, height, tl.ColorBlack)

	return retour
}

func (bt *BigText) Draw(screen *tl.Screen) {
	bt.Rectangle.Draw(screen)
	textLines := strings.Split(bt.Content, "\n")
	bt.Texts = nil
	bt.maxTextLength = 0

	for index, line := range textLines {
		newText := tl.NewText(0 - bt.XOffset,index - bt.YOffset, line, tl.ColorWhite, tl.ColorBlack)
		bt.Texts = append(bt.Texts, newText)
		textLen := len(line)
		if (textLen > bt.maxTextLength) {
			bt.maxTextLength = textLen
		}
		newText.Draw(screen)
	}
}

func (bt *BigText) Tick(event tl.Event) {
	//TheGame.Log("event")
	if event.Type == tl.EventKey {
		switch event.Key { // If so, switch on the pressed key.
		default:
			if TheGame.IsBarDisplayed {
				return
			}
			switch event.Ch {
			case 'h':
				bt.MoveLeft(1)
			case 'j':
				bt.MoveDown(1)
			case 'k':
				bt.MoveTop(1)
			case 'l':
				bt.MoveRight(1)
			case ':':
				TheGame.DisplayCommandBar()
			}
		}
	}
}

func (bt *BigText) MoveLeft(index int) {
	bt.MoveX(-1)
}

func (bt *BigText) MoveRight(index int) {
	bt.MoveX(1)
}

func (bt *BigText) MoveTop(index int) {
	bt.MoveY(-1)
}

func (bt *BigText) MoveDown(index int) {
	bt.MoveY(1)
}

func (bt *BigText) MoveX (offset int) {
	if ((offset < 0 && bt.XOffset + offset < 0) || (offset > 0 && bt.XOffset + offset > bt.maxTextLength - 1)) {
		return
	}
	bt.XOffset += offset
}

func (bt *BigText) MoveY (offset int) {
	if ((offset < 0 && bt.YOffset + offset < 0) || (offset > 0 && bt.YOffset + offset > len(bt.Texts) - 1)) {
		return
	}
	bt.YOffset += offset
}