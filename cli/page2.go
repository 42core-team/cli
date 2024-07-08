package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Page2Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewPage2Model() Page2Model {
	m := Page2Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("class").
				Options(huh.NewOptions("List Teams", "New Team")...).
				Title("Main menu").
				Description("Choose what you want to do"),
		),
	)
	return m
}

func (m Page2Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Page2Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		return NewModel(), nil
	}

	return m, tea.Batch(cmds...)
}

func (m Page2Model) View() string {
	s := m.styles

	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	body := lipgloss.JoinHorizontal(lipgloss.Top, form)

	return s.Base.Render(body)

}
