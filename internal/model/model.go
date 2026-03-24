package model

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NunoFrRibeiro/hello-tui/internal/constants"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			PaddingLeft(2).
			PaddingTop(1).
			PaddingBottom(1)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF007C")).
				Bold(true).
				PaddingLeft(2)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			PaddingLeft(4)

	resultBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			MarginTop(1).
			MarginLeft(2).
			Width(50)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			PaddingLeft(2).
			PaddingTop(1)
)

type Model struct {
	Languages        []constants.Language
	Cursor           int
	SelectedLanguage *int
	Quit             bool
}

func NewModel() Model {
	return Model{
		Languages:        constants.AvailableLanguages,
		Cursor:           0,
		SelectedLanguage: nil,
		Quit:             false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			m.Quit = true
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Languages)-1 {
				m.Cursor++
			}

		case "enter", " ":
			selected := m.Cursor
			m.SelectedLanguage = &selected

		case "r":
			m.SelectedLanguage = nil
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	var builder strings.Builder

	builder.WriteString(titleStyle.Render("🌍 Hello World - Language Selector"))
	builder.WriteString("\n\n")

	if m.SelectedLanguage != nil {
		lang := m.Languages[*m.SelectedLanguage]
		result := fmt.Sprintf("Language; %s (%s)\n\n%s", lang.Name, lang.Code, lang.Translation)

		builder.WriteString(resultBoxStyle.Render(result))
		builder.WriteString("\n\n")
		builder.WriteString(helpStyle.Render("Press 'r' to select another language or 'q' to quit"))

	} else {
		builder.WriteString(normalItemStyle.Render("Select a language:"))
		builder.WriteString("\n\n")

		for i, lang := range m.Languages {

			cursor := " "

			if m.Cursor == i {
				cursor = "→ "
			}

			line := fmt.Sprintf("%s %s", cursor, lang.Name)

			if m.Cursor == i {
				builder.WriteString(selectedItemStyle.Render(line))
			} else {
				builder.WriteString(normalItemStyle.Render(line))
			}

			builder.WriteString("\n")
		}

		builder.WriteString("\n")
		builder.WriteString(helpStyle.Render("↑/k: up • ↓/j: down • enter/space: select • q: quit"))
	}

	return tea.NewView(builder.String())
}
