package main

import (tl "github.com/JoelOtter/termloop")

type EntityItem struct {
	*tl.Entity
	*Item
	number int
}

func NewEntityItem(x, y int, item *Item) *EntityItem {
	entity := tl.NewEntity(x, y,1,1)
	entity.SetCell(0,0, &tl.Cell{
		Ch: item.Icon,
	})

	entityItem := new (EntityItem)
	entityItem.Entity = entity
	entityItem.Item = item

	return entityItem
}

func (ei *EntityItem) GetNumber() int {
	if (ei.RareItem) {
		return 1
	} else {
		return ei.number
	}
}