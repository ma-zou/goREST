package ui

import (
	"fmt"
	"goREST/internal/request"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Table(selected int, data []request.RequestObject) string {

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(TableBorder).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return TableHead
			case row == selected:
				return TableHighlight
			default:
				return TableBase
			}
		}).
		Headers(
			"URL",
			"Method",
			"Headers",
			"Params",
			"Body",
			"Auth",
		)

	for _, d := range data {
		t.Row(
			d.URL,
			d.Method,
			fmt.Sprintf("%d", len(d.Headers)),
			fmt.Sprintf("%d", len(d.Params)),
			d.Body,
			d.Auth.Type,
		)
	}

	return t.String()
}
