package main

import tl "github.com/JoelOtter/termloop"

type Wall struct {
	*tl.Rectangle
}

func NewWall(x, y int) *Wall {
	wall := new(Wall)
	wall.Rectangle = tl.NewRectangle(x, y, 1, 1, tl.ColorWhite)

	return wall
}
