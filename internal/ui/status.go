package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Status(maxWidth, height, maxHeight int) string {
	s := Highlight.Render("goREST 🗲")
	c := strings.Builder{}

	statusBox := lipgloss.NewStyle().
		MarginTop(maxHeight - height - 4)

	c.WriteString(BackdropLight.Render("↑ / ↓"))
	c.WriteString(Backdrop.Render("Select"))

	c.WriteString(BackdropLight.Render("↵"))
	c.WriteString(Backdrop.Render("Send"))

	c.WriteString(BackdropLight.Render("n"))
	c.WriteString(Backdrop.Render("New"))

	c.WriteString(BackdropLight.Render("e"))
	c.WriteString(Backdrop.Render("Edit"))

	c.WriteString(BackdropLight.Render("q"))
	c.WriteString(Backdrop.Render("Quit"))

	cr := lipgloss.NewStyle().
		Width(maxWidth - lipgloss.Width(s)).
		Background(lipgloss.Color("236"))

	a := Highlight2.Render("Select Request")

	return statusBox.Render(lipgloss.JoinHorizontal(lipgloss.Left, s, cr.Render(c.String()), a))
}
