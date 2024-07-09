package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type ManageTeamMenuModel struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewManageTeamMenuModel() ManageTeamMenuModel {
	m := ManageTeamMenuModel{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("option").
				Options(huh.NewOptions("List Teams", "New Team")...).
				Title("Manage team menu").
				Description("Choose what you want to do"),
		),
	)
	m.Init()
	return m
}

func (m ManageTeamMenuModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m ManageTeamMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "List Teams":
			return NewMainMenuModel(), nil
		case "New Team":
			return NewMainMenuModel(), nil
		default:
			fmt.Println("Go fuck yourself")
		}
	}

	return m, tea.Batch(cmds...)
}

func (m ManageTeamMenuModel) View() string {
	s := m.styles

	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	body := lipgloss.JoinHorizontal(lipgloss.Top, form)

	return s.Base.Render(body)

}
