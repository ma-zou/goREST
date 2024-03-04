package ui

import (
	"fmt"
	"goREST/internal/request"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type model struct {
	requests []request.RequestObject
	cursor   int
}

func InitialModel(requests []request.RequestObject) model {
	return model{
		requests: requests,
		cursor:   1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			fmt.Print("new request")
		case "up", "k":
			if m.cursor > 1 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.requests) {
				m.cursor++
			}
		case "enter", " ":
			fmt.Print(m.cursor)
		}
	}

	return m, nil
}

func (m model) View() string {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))
	uiWidth := physicalWidth

	doc := strings.Builder{}

	t := Table(m.cursor, m.requests)
	d := Details(m.requests[m.cursor-1])

	doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, t, d))

	uiWidth = lipgloss.Width(doc.String())
	uiHeight := lipgloss.Height(doc.String())

	doc.WriteString("\n\n")

	doc.WriteString(Status(uiWidth, uiHeight, physicalHeight))

	if physicalWidth > 0 {
		DocStyle = DocStyle.MaxWidth(physicalWidth)
	}

	return DocStyle.Render(doc.String())
}
