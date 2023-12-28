package main

import (
	"fmt"
)

type Menu struct {
	Prompt         string
	CursorPosition int
	Items          []*MenuItem
}

type MenuItem struct {
	Text string
	Id   string
}

func CreateMenu(prompt string) *Menu {
	return &Menu{
		Prompt: prompt,
		Items:  make([]*MenuItem, 0),
	}
}

func (m *Menu) AddMenuItem(text string, id string) *Menu {
	menuItem := &MenuItem{
		Text: text,
		Id:   id,
	}
	m.Items = append(m.Items, menuItem)
	return m
}

func (m *Menu) renderMenu(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", len(m.Items)-1)
	}
	for index, item := range m.Items {
		var newLine = "\n"
		if index == len(m.Items)-1 {
			newLine = " "
		}
		menuItemText := item.Text
		cursor := "  "
		if index == m.CursorPosition {
			cursor = " >"
		}
		fmt.Printf("\r%s %s%s", cursor, menuItemText, newLine)
	}
}

func (m *Menu) Display() string {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()

	// Clear the screen
	// fmt.Printf("\033[2J")
	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	fmt.Printf("%s\n", m.Prompt+":")
	m.renderMenu(false)

	for {
		keyCode := GetInput()
		if keyCode == escape {
			return ""
		} else if keyCode == enter {
			menuItem := m.Items[m.CursorPosition]
			fmt.Println("\r")
			return menuItem.Id
		} else if keyCode == up {
			m.CursorPosition = (m.CursorPosition + len(m.Items) - 1) % len(m.Items)
			m.renderMenu(true)
		} else if keyCode == down {
			m.CursorPosition = (m.CursorPosition + 1) % len(m.Items)
			m.renderMenu(true)
		}
	}
}
