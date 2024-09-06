package ui

import (
	// "github.com/abdheshnayak/typing-game/apps/client/internal/domain"
	"fmt"

	"github.com/abdheshnayak/typing-game/apps/client/internal/domain"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"

	// "github.com/rivo/tview"
	"go.uber.org/fx"
)

var HomeModule = fx.Module("home",
	fx.Invoke(func(d domain.Domain) {

		p := tea.NewProgram(initialModel(d))

		if _, err := p.Run(); err != nil {
			panic(err)
		}

		// app := tview.NewApplication()
		//
		// textview := tview.NewTextView().SetText("hello world")
		//
		// if err := app.SetRoot(textview, true).EnableMouse(true).Run(); err != nil {
		// 	panic(err)
		// }

	}),
)

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
	domain   domain.Domain
}

func initialModel(d domain.Domain) model {
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
		domain:   d,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	// fmt.Println("hello world")

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}

		case tea.KeyCtrlC:
			return m, tea.Quit

		default:
			m.domain.UserDebugLog(msg.String())
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func log() {

}
