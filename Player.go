package main

import (
	"VimUniverse/MapEntities"
	//"strconv"
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
	lives int
}

func (player *Player) Tick(event tl.Event) {
	//game.Log("event")
	if event.Type == tl.EventKey {
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyCtrlC:
			game.DisplayInfoText("Type :q and press <Enter> to abandon your progression and exit VimUniverse.", tl.ColorWhite)
		default:
			if game.isBarDisplayed {
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
				game.DisplayCommandBar()
			}
		}
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Collide(collision tl.Physical) {
	if _, rect := collision.(*tl.Rectangle); rect {
		player.CancelMove()
	} else if _, waterSquare := collision.(*MapEntities.WaterSquare); waterSquare {
		player.CancelMove()
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
}

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)
	player.lives = 3

	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '옷'})

	return player
}
