package main

import (
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/NunoFrRibeiro/hello-tui/internal/model"
)

func main() {
	p := tea.NewProgram(model.NewModel())

	if _, err := p.Run(); err != nil {
		slog.Error("error running program: %v", err)
		os.Exit(1)
	}
}
