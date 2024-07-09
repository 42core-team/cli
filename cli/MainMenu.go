package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type MainMenuModel struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
	chosen string
}

func NewMainMenuModel() MainMenuModel {
	m := MainMenuModel{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("option").
				Options(huh.NewOptions("Manage Teams", "Tournament", "Leaderboard")...).
				Title("Main menu").
				Description("Choose what you want to do").Value(&m.chosen),
		),
	)
	m.Init()
	return m
}

func (m MainMenuModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
			case "esc", "ctrl+c", "q":
				return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		switch m.form.Get("option") {
		case "Manage Teams":
			return NewManageTeamMenuModel(), nil
		case "Tournament":
			return NewMainMenuModel(), nil
		case "Leaderboard":
			return NewMainMenuModel(), nil
		default:
			fmt.Println("Go fuck yourself")
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MainMenuModel) View() string {
	s := m.styles

	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	body := lipgloss.JoinHorizontal(lipgloss.Top, form)
	return s.Base.Render(body)
}