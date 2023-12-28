package main

import (
	"fmt"
	"strings"
)

type Direction int

const (
    U Direction = iota
    D
    L
    R
)

type IDirection interface {
	Get() Direction
}

type BodyPart struct {
	X int
	Y int
}

type SnakeGame struct {
	Size GridSize
	direction IDirection 
	Speed int
	Snake []BodyPart
}

func createSnakeGame(size GridSize, speed int) *SnakeGame {
	return &SnakeGame{
		Size: size,
		Speed: speed,
		Snake: getStartSnake(size),
	}
}

func (game *SnakeGame) getSnakeParts (lineNumber int) []BodyPart {
	parts := []BodyPart{}
	for _, part := range game.Snake {
		if part.Y == lineNumber {
			parts = append(parts, part)
		}
	}
	return parts
}

func (game *SnakeGame) createSnakePart (direction Direction) BodyPart {
	var x = game.Snake[0].X
	var y = game.Snake[0].Y
	if direction == U {
		return BodyPart{X: x, Y: y-1}
	} else if direction == D {
		return BodyPart{X: x, Y: y+1}
	} else if direction == L {
		return BodyPart{X: x-1, Y: y}
	} else if direction == R {
		return BodyPart{X: x+1, Y: y}
	}
	return BodyPart{X: x, Y: y-1} 
}

func getStartSnake(size GridSize) []BodyPart {
	first := BodyPart{
		X: size.width/2,
		Y: size.height-4,
	}
	second := BodyPart{
		X: size.width/2,
		Y: size.height-3,
	}
	third := BodyPart{
		X: size.width/2,
		Y: size.height-2,
	}
	parts := []BodyPart{first, second, third}
	return parts
}

func (game *SnakeGame) renderGame(redraw bool){
	if redraw {
		fmt.Printf("\033[%dA", game.Size.height)
	}
	//fmt.Printf("\r%s", strings.Repeat("_", game.Size.width))
	for i := 0; i < game.Size.height; i++ {
		var s = strings.Repeat(".", game.Size.width)
		var parts = game.getSnakeParts(i)
		for _, part := range parts {
			s = s[:part.X] + "x" + s[part.X+1:]
		}
		fmt.Printf("\r%s%s", s, "\n")
	}
	//fmt.Printf("\r%s", strings.Repeat("_", game.Size.width))
}

func (game *SnakeGame) Run() string {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()
	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	game.renderGame(false)
	for {
		keyCode := GetInput()
		if keyCode == escape {
			return ""
		} else if keyCode == enter {
			return ""
		} else if keyCode == up {
			part := game.createSnakePart(U)
			game.Snake = append([]BodyPart{part}, game.Snake...)
			game.renderGame(true)
		} else if keyCode == down {
			part := game.createSnakePart(D)
			game.Snake = append([]BodyPart{part}, game.Snake...)
			game.renderGame(true)
		} else if keyCode == left {
			part := game.createSnakePart(L)
			game.Snake = append([]BodyPart{part}, game.Snake...)
			game.renderGame(true)
		} else if keyCode == right {
			part := game.createSnakePart(R)
			game.Snake = append([]BodyPart{part}, game.Snake...)
			game.renderGame(true)
		}
	}
}