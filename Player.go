package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	Level *tl.BaseLevel
	Lives int
	onItem *EntityItem
	inventory [](*Item)
	CanFloat bool
	floating bool
}

func (player *Player) Tick(event tl.Event) {
	//TheGame.Log("event")
	if event.Type == tl.EventKey {
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyCtrlC:
			TheGame.DisplayInfoText("Type :q and press <Enter> to abandon your progression and exit VimUniverse.", tl.ColorWhite)
		default:
			if TheGame.IsBarDisplayed || TheGame.IsBigTextDisplayed {
				return
			}
			switch event.Ch {
			case 'h':
				player.MoveLeft(1)
			case 'j':
				player.MoveDown(1)
			case 'k':
				player.MoveTop(1)
			case 'l':
				player.MoveRight(1)
			case ':':
				TheGame.DisplayCommandBar()
			}
		}
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Collide(collision tl.Physical) {
	if _, rect := collision.(*Wall); rect {
		player.CancelMove()
	}
	if _, waterSquare := collision.(*WaterSquare); waterSquare {
		if (player.CanFloat) {
			if (!player.floating) {
				player.UseBoat()
			}
		} else {
			player.CancelMove()
		}
	}
	if entityItem, ok := collision.(*EntityItem); ok {
		if (player.onItem != entityItem) {
			player.onItem = entityItem
			TheGame.DisplayInfoTextWithTime("You walked on an item ("+entityItem.Item.Name+"). Type \":pick\" to pick it up.", tl.ColorWhite, 10)
		}
	} else {
		player.onItem = nil
	}
}

func (player *Player) MoveTop(nb int) {
	player.Move(0, -1*nb)
}

func (player *Player) CancelMove() {
	player.SetPosition(player.prevX, player.prevY)
}

func (player *Player) MoveLeft(nb int) {
	player.Move(-1*nb, 0)
}

func (player *Player) MoveRight(nb int) {
	player.Move(nb, 0)
}

func (player *Player) MoveDown(nb int) {
	player.Move(0, nb)
}

func (player *Player) Move(x, y int) {
	player.prevX, player.prevY = player.Position()
	player.SetPosition(player.prevX+x, player.prevY+y)
	if (player.floating) {
		player.UseBoat()
	}
}

func (player *Player) UseBoat() {
	if (!player.floating) {
		player.floating = true
		TheGame.Log(strconv.FormatBool(player.floating))
		player.SetCell(0, 0, &tl.Cell {Fg: tl.ColorBlack, Ch: '⌴'})
	} else {
		player.floating = false
		TheGame.Log("retour")
		player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '옷'})
	}
}

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)
	player.Lives = 3
	player.CanFloat = false
	player.floating = false

	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '옷'})

	return player
}
