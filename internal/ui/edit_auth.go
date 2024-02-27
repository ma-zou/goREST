package ui

import (
	"log"

	"github.com/charmbracelet/huh"
)

func RenderAuthForm() string {
	var authToken string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Auth Token").
				Value(&authToken),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return authToken
}
