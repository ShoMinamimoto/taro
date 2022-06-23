package kiss

import "github.com/charmbracelet/lipgloss"
import tea "github.com/charmbracelet/bubbletea"

type Wrapper struct {
	Style  lipgloss.Style
	Height int
	Width  int
}

var (
	HotPink = lipgloss.AdaptiveColor{Light: "#F94BCE", Dark: "#C6189B"}
)

func (w Wrapper) Resize(msg tea.WindowSizeMsg) Wrapper {
	w.Height = msg.Height - w.Style.GetVerticalFrameSize()
	w.Width = msg.Width - w.Style.GetHorizontalFrameSize()
	return w
}

func (w Wrapper) SetStyle(style lipgloss.Style) Wrapper {
	//account for frame size differences
	w.Height += w.Style.GetVerticalFrameSize()
	w.Height -= style.GetVerticalFrameSize()
	w.Width += w.Style.GetHorizontalFrameSize()
	w.Width -= style.GetHorizontalFrameSize()

	w.Style = style
	return w
}

func NewDefaultWrapper() Wrapper {
	return Wrapper{
		Style: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(HotPink).
			Padding(1, 2).
			Margin(1, 2),
		Height: 18,
		Width:  70,
	}
}

func (w Wrapper) Render(string string) string {
	return w.Style.Render(string)
}
