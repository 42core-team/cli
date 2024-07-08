package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type view int

const (
	mainMenu view = iota
	manageTeams
	tournaments
	startLocalGame
)

type model struct {
	currentView view
	menu        list.Model
}

const listHeight = 10

func initialModel() model {
	items := []list.Item{
		item("Manage Teams"),
		item("Tournaments"),
		item("Start Local Game"),
	}

	const defaultWidth = 20
	menu := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	menu.Title = "Main Menu"
	menu.SetShowStatusBar(false)
	menu.SetFilteringEnabled(false)
	menu.Styles.Title = titleStyle
	menu.Styles.PaginationStyle = paginationStyle
	menu.Styles.HelpStyle = helpStyle

	return model{
		currentView: mainMenu,
		menu:        menu,
	}
}


func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch {
			case msg.String() == "q":
				if m.currentView == mainMenu {
					return m, tea.Quit
				}
				m.currentView = mainMenu
				return m, nil
			case msg.String() == "enter":
				if m.currentView == mainMenu {
					switch m.menu.SelectedItem().(item) {
					case "Manage Teams":
						m.currentView = manageTeams
					case "Tournaments":
						m.currentView = tournaments
					case "Start Local Game":
						m.currentView = startLocalGame
					}
					return m, nil
				}
		}
	}

	if m.currentView == mainMenu {
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	}

	return m, nil
}


func (m model) View() string {
	var content string

	switch m.currentView {
	case manageTeams:
		content = initialManageTeamModel().View()
	case tournaments:
		content = tournamentsView()
	case startLocalGame:
		content = startLocalGameView()
	default:
		content = "\n" + m.menu.View() + "\n"
	}

	return content + "\n"
}

