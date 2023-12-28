package main

import (
	"fmt"
)

type GridSize struct {
	width int
	height int
}

var sizeOptions = map[string]GridSize {
	"1": {width: 10, height: 10},
	"2": {width: 20, height: 20},
	"3": {width: 30, height: 30},
	"4": {width: 40, height: 40},
}

var speedOptions = map[string]int {
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
}

func getSizeOption() string {
	menu := CreateMenu("Select grid size")
	for key, value := range sizeOptions {
		menu.AddMenuItem(fmt.Sprintf("%d x %d", value.width, value.height), key)
	}
	return menu.Display()
}

func getSpeedOption() string {
	menu := CreateMenu("Select speed")
	for key, value := range speedOptions {
		menu.AddMenuItem(fmt.Sprintf("%d", value), key)
	}
	return menu.Display()
}

func main() {
	var selectedSizeOption = getSizeOption()
	fmt.Println(selectedSizeOption)

	var selectedSpeedOption = getSpeedOption()
	var game = createSnakeGame(sizeOptions[selectedSizeOption], speedOptions[selectedSpeedOption])
	fmt.Println(selectedSpeedOption)
	game.Run()
}
