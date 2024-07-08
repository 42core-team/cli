package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	manageTeamsMenu view = iota
	listTeams
	newTeam
)

type ManageTeamModel struct {
	currentView view
	menu        list.Model
}

func initialManageTeamModel() ManageTeamModel {
	items := []list.Item{
		item("List Teams"),
		item("New Team"),
	}

	const defaultWidth = 20
	menu := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	menu.Title = "Main Menu"
	menu.SetShowStatusBar(false)
	menu.SetFilteringEnabled(false)
	menu.Styles.Title = titleStyle
	menu.Styles.PaginationStyle = paginationStyle
	menu.Styles.HelpStyle = helpStyle

	return ManageTeamModel{
		currentView: mainMenu,
		menu:        menu,
	}
}


func (m ManageTeamModel) Init() tea.Cmd {
	return nil
}

func (m ManageTeamModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
					case "List Teams":
						m.currentView = listTeams
					case "New Team":
						m.currentView = newTeam
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


func (m ManageTeamModel) View() string {
	var content string

	switch m.currentView {
	case listTeams:
		content = initialManageTeamModel().View()
	case newTeam:
		content = tournamentsView()
	default:
		content = "\n" + m.menu.View() + "\n"
	}

	return content + "\n"
}

