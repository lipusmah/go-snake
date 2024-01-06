package main

import (
	"fmt"
	"sort"
	"strconv"
)

var sizeOptions = map[string]GridSize{
	"1": {width: 10, height: 10},
	"2": {width: 20, height: 10},
	"3": {width: 20, height: 20},
	"4": {width: 40, height: 20},
	"5": {width: 30, height: 30},
	"6": {width: 40, height: 40},
	"7": {width: 60, height: 40},
}

var speedOptions = [12]int{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
}

var colors = map[string]string{
	"red":   "\x1b[1;91m",
	"cyan":  "\x1b[1;36m",
	"clear": "\x1b[m",
}

type GridSize struct {
	width  int
	height int
}

func getSizeOption() string {
	keys := make([]string, 0, len(sizeOptions))

	for k := range sizeOptions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	menu := CreateMenu("Select grid size")
	for _, v := range keys {
		var value = sizeOptions[v]
		menu.AddMenuItem(fmt.Sprintf("%d x %d", value.width, value.height), v)
	}
	return menu.Display()
}

func getSpeedOption() string {
	menu := CreateMenu("Select speed")
	for v := range speedOptions {
		var value = speedOptions[v]
		menu.AddMenuItem(fmt.Sprintf("%d", value), strconv.Itoa(v))
	}
	return menu.Display()
}

func main() {
	fmt.Printf(colors["red"])
	fmt.Printf("\nGO-Snake\n")
	fmt.Printf(colors["clear"])

	var selectedSizeOption = getSizeOption()
	fmt.Printf("\033[%dA", len(sizeOptions)+1)
	fmt.Printf("\033[0J")

	var selectedSpeedOption = getSpeedOption()
	fmt.Printf("\033[%dA", len(speedOptions)+1)
	fmt.Printf("\033[0J")

	speed, err := strconv.Atoi(selectedSpeedOption)
	if err != nil {
		fmt.Printf("Error with configuration")
	}
	var game = createSnakeGame(sizeOptions[selectedSizeOption], speed)
	game.Run()
	fmt.Printf("\n")
}
