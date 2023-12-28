package main

import (
	"log"

	"github.com/pkg/term"
)

// Raw input keycodes
var up byte = 65
var down byte = 66
var left byte = 68
var right byte = 67

var escape byte = 27
var enter byte = 13

var keys = map[byte]bool{
	up:    true,
	down:  true,
	left:  true,
	right: true,
}

// getInput will read raw input from the terminal
// It returns the raw ASCII value inputted
func GetInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	t.Restore()
	t.Close()

	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
	// See: https://en.wikipedia.org/wiki/ANSI_escape_code
	if read == 3 {
		if _,  ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}
