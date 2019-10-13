package main

import tl "github.com/JoelOtter/termloop"

type OpenDoor struct {
	*tl.Entity
}

func NewOpenDoor(x, y int) *OpenDoor {
	ld := new(OpenDoor)
	ld.Entity = tl.NewEntity(x, y, 1, 1)
	ld.SetCell(0, 0, &tl.Cell{
		Fg: tl.ColorBlack,
		Bg: tl.ColorWhite,
		Ch: 'H',
	})

	return ld
}
