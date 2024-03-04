package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var DocStyle = lipgloss.NewStyle().
	Padding(1, 2, 1, 2)

var BackdropText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("238")).
	Italic(true)

var HighlightText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("51")).
	Bold(true)

var Highlight = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("242")).
	Background(lipgloss.Color("51")).
	Padding(0, 2)

var Highlight2 = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("51")).
	Background(lipgloss.Color("242")).
	Padding(0, 2)

var Backdrop = lipgloss.NewStyle().
	Foreground(lipgloss.Color("242")).
	Background(lipgloss.Color("236")).
	PaddingLeft(1).
	PaddingRight(2)

var BackdropLight = lipgloss.NewStyle().
	Foreground(lipgloss.Color("248")).
	Background(lipgloss.Color("236")).
	PaddingLeft(2)

var TableBorder = lipgloss.NewStyle().
	Foreground(lipgloss.Color("236"))

var TableBase = lipgloss.NewStyle().
	Padding(0, 1)

var TableHighlight = TableBase.Copy().
	Foreground(lipgloss.Color("242")).
	Background(lipgloss.Color("51")).
	Bold(true)

var TableHead = TableBase.Copy().
	Foreground(lipgloss.Color("255")).
	Bold(true)

var Detail = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("236")).
	Padding(0, 1)

var DefinitionList = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderRight(false).
	BorderBottom(false).
	BorderForeground(lipgloss.Color("236")).
	Padding(0, 1)

var DefinitionTerm = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("255"))

var DefinitionData = lipgloss.NewStyle().
	Foreground(lipgloss.Color("242"))

var Method = lipgloss.NewStyle().
	Background(lipgloss.Color("236")).
	Padding(0, 1)

var GET = Method.Copy().
	Foreground(lipgloss.Color("10"))

var POST = Method.Copy().
	Foreground(lipgloss.Color("220"))

var PUT = Method.Copy().
	Foreground(lipgloss.Color("33"))

var PATCH = Method.Copy().
	Foreground(lipgloss.Color("99"))

var DELETE = Method.Copy().
	Foreground(lipgloss.Color("9"))
