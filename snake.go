package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Direction int

const (
	U Direction = iota
	D
	L
	R
)

type GameCell struct {
	X int
	Y int
}

type SnakeGame struct {
	Size      GridSize
	Direction Direction
	Speed     int
	Over      bool
	Snake     []GameCell
	Fruit     GameCell
	Score     int
}

func createSnakeGame(size GridSize, speed int) *SnakeGame {
	var game = &SnakeGame{
		Size:      size,
		Speed:     speed,
		Snake:     getStartSnake(size),
		Over:      false,
		Direction: U,
		Score:     0,
	}
	game.Fruit = game.createFruitCell()
	return game
}

func (game *SnakeGame) getSnakeParts(lineNumber int) []GameCell {
	parts := []GameCell{}
	for _, part := range game.Snake {
		if part.Y == lineNumber {
			parts = append(parts, part)
		}
	}
	return parts
}

func (game *SnakeGame) createSnakePart(direction Direction) GameCell {
	var x = game.Snake[0].X
	var y = game.Snake[0].Y
	if direction == U {
		return GameCell{X: x, Y: y - 1}
	} else if direction == D {
		return GameCell{X: x, Y: y + 1}
	} else if direction == L {
		return GameCell{X: x - 1, Y: y}
	} else if direction == R {
		return GameCell{X: x + 1, Y: y}
	}
	return GameCell{X: x, Y: y - 1}
}

func (game *SnakeGame) createFruitCell() GameCell {
	for {
		randX := rand.Intn(game.Size.width)
		randY := rand.Intn(game.Size.height)
		var cell = GameCell{X: randX, Y: randY}
		if !(game.isInSnake(cell)) {
			return cell
		}
	}
}

func getStartSnake(size GridSize) []GameCell {
	first := GameCell{
		X: size.width / 2,
		Y: size.height - 4,
	}
	second := GameCell{
		X: size.width / 2,
		Y: size.height - 3,
	}
	third := GameCell{
		X: size.width / 2,
		Y: size.height - 2,
	}
	parts := []GameCell{first, second, third}
	return parts
}

func (game *SnakeGame) renderGame(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", game.Size.height+1)
	}
	fmt.Printf("\r%dx%d, speed: %d, %sscore: %d%s%s", game.Size.height, game.Size.width, game.Speed, colors["cyan"], game.Score, colors["clear"], "\n")

	for i := 0; i < game.Size.height; i++ {
		var s = strings.Repeat(".", game.Size.width)
		var parts = game.getSnakeParts(i)
		if game.Fruit.Y == i {
			s = s[:game.Fruit.X] + "#" + s[game.Fruit.X+1:]
		}
		for _, part := range parts {
			s = s[:part.X] + "o" + s[part.X+1:]
		}
		fmt.Printf("\r%s%s", s, "\n")
	}
}

func (game *SnakeGame) isInSnake(newPart GameCell) bool {
	for _, part := range game.Snake {
		if part.X == newPart.X && part.Y == newPart.Y {
			return true
		}
	}
	return false
}

func (game *SnakeGame) updateState() string {
	part := game.createSnakePart(game.Direction)

	// snake intersect itself exit condition
	if game.isInSnake(part) {
		game.Over = true
		return "intersect"
	}

	// hit the fruit
	if part.X == game.Fruit.X && part.Y == game.Fruit.Y {
		game.Fruit = game.createFruitCell()
		game.Snake = append([]GameCell{part}, game.Snake...)
		game.Score = game.Score + game.Speed
	} else {
		game.Snake = append([]GameCell{part}, game.Snake[:len(game.Snake)-1]...)
	}

	// snake hits the wall exit condition
	if part.X >= game.Size.width || part.X < 0 {
		game.Over = true
		return "wall"
	}
	if part.Y >= game.Size.height || part.Y < 0 {
		game.Over = true
		return "wall"
	}
	return "error"
}

func (game *SnakeGame) gameLoop(wg *sync.WaitGroup) string {
	defer wg.Done()
	for {
		var offset = 1.0 / float32(game.Speed)
		time.Sleep(time.Duration(offset * float32(time.Second)))
		game.updateState()
		if game.Over {
			return ""
		}
		if !(game.Over) {
			game.renderGame(true)
		}
	}
}

func (game *SnakeGame) inputLoop(wg *sync.WaitGroup) string {
	defer wg.Done()
	for {
		if game.Over {
			return ""
		}
		keyCode := GetInput()
		if keyCode == escape {
			game.Over = true
			return ""
		} else if keyCode == enter {
			game.Over = true
			return ""
		} else if keyCode == up {
			if game.Direction != D {
				game.Direction = U
			}
		} else if keyCode == down {
			if game.Direction != U {
				game.Direction = D
			}
		} else if keyCode == left {
			if game.Direction != R {
				game.Direction = L
			}

		} else if keyCode == right {
			if game.Direction != L {
				game.Direction = R
			}
		}
	}
}

func (game *SnakeGame) Run() string {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()
	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	game.renderGame(false)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go game.gameLoop(&wg)
	wg.Add(1)
	go game.inputLoop(&wg)
	wg.Wait()
	fmt.Printf("\n\r%sGAME OVER%s", colors["red"], colors["clear"])
	fmt.Printf("\n")

	return ""
}
