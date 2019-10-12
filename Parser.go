package main

import (
	tl "github.com/JoelOtter/termloop"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadLevel(level int) {
	path := "levels" + string(os.PathSeparator) + strconv.Itoa(level)
	file, err := os.Open(path)

	type StartingPos struct {
		x int
		y int
	}
	currentLine := 0
	currentColumn := 0
	var startingPos StartingPos
	check(err)

	readBuffer := make([]byte, 32*1024)

	newLevel := tl.NewBaseLevel(tl.Cell{
		Fg: tl.ColorWhite,
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	for {
		n, err := file.Read(readBuffer)
		if err == io.EOF {
			break
		} else if err != nil {
			check(err)
		}
		readText := string(readBuffer[:n])
		for _, char := range readText {
			if char == '\n' {
				currentColumn = 0
				currentLine++
			} else {
				switch char {
				case ' ':
					// Do nothing
				case '~':
					waterSquare := NewWaterSquare(currentColumn, currentLine)
					newLevel.AddEntity(waterSquare)
				case 'S':
					startingPos = StartingPos{
						x: currentColumn,
						y: currentLine,
					}
				case 'W':
					wall := NewWall(currentColumn, currentLine)
					newLevel.AddEntity(wall)
				case 'B':
					boat := NewBoatItem()
					boatEntity := NewEntityItem(currentColumn,currentLine,boat)
					newLevel.AddEntity(boatEntity)
				default:
					continue
				}
				currentColumn++
			}
		}
	}

	TheGame.Player.Level = newLevel
	TheGame.Screen().SetLevel(newLevel)
	newLevel.AddEntity(TheGame.Player)
	TheGame.Player.SetPosition(startingPos.x, startingPos.y)
}
