package MapEntities

import tl "github.com/JoelOtter/termloop"

type WaterSquare struct {
	*tl.Entity
}

func NewWaterSquare(x, y int) *WaterSquare {
	waterSquare := new(WaterSquare)

	waterSquare.Entity = tl.NewEntity(x, y, 1, 1)
	waterSquare.SetCell(0, 0, &tl.Cell{
		Fg: tl.ColorBlack,
		Bg: tl.ColorWhite,
		Ch: '~',
	})

	return waterSquare
}
