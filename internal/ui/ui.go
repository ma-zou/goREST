package ui

import (
	"fmt"
	"goREST/internal/request"

	tea "github.com/charmbracelet/bubbletea"
)

type UI struct {
	data []request.RequestObject
}

func (u *UI) Update(data interface{}) {
	requests := data.([]request.RequestObject)

	p := tea.NewProgram(InitialModel(requests))

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		return
	}
}
