package main

import tl "github.com/JoelOtter/termloop"

type LockedDoor struct {
	*tl.Entity
}

func NewLockedDoor(x, y int) *LockedDoor {
	ld := new(LockedDoor)
	ld.Entity = tl.NewEntity(x, y, 1, 1)
	ld.SetCell(0, 0, &tl.Cell{
		Fg: tl.ColorBlack,
		Bg: tl.ColorWhite,
		Ch: 'ℍ',
	})

	return ld
}
