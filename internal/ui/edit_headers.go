package ui

import (
	"log"

	"github.com/charmbracelet/huh"
)

func RenderHeadersForm() (bool, string, string) {
	var key string
	var value string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Header Key").
				Inline(true).
				Value(&key),
			huh.NewInput().
				Title("Header Value").
				Inline(true).
				Value(&value),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if key == "" || value == "" {
		return false, "", ""
	}

	return true, key, value
}
