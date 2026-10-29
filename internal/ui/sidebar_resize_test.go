package ui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/gammons/slk/internal/ui/sidebar"
)

func newResizeApp(enabled bool) *App {
	a := NewApp()
	a.width = 160
	a.height = 40
	a.sidebarVisible = true
	a.sidebar.SetItems([]sidebar.ChannelItem{{ID: "C1", Name: "general", Type: "channel"}})
	a.SetSidebarStyle(SidebarStyle{MouseResize: enabled})
	_ = a.View()
	return a
}

func TestSidebarResize_DragMovesTheBorder(t *testing.T) {
	a := newResizeApp(true)
	before := a.sidebar.Width()
	border := a.layout.SidebarEnd() - 1

	a.Update(tea.MouseClickMsg{X: border, Y: 5, Button: tea.MouseLeft})
	a.Update(tea.MouseMotionMsg{X: border + 10, Y: 5, Button: tea.MouseLeft})
	a.Update(tea.MouseReleaseMsg{X: border + 10, Y: 5, Button: tea.MouseLeft})
	_ = a.View()

	if got := a.sidebar.Width(); got != before+10 {
		t.Fatalf("width after a 10-column drag = %d, want %d", got, before+10)
	}
	if got := a.layout.SidebarEnd() - 1; got != border+10 {
		t.Errorf("border at %d, want it under the pointer at %d", got, border+10)
	}
	if a.resizingSidebar {
		t.Error("release left the resize gesture active")
	}
}

func TestSidebarResize_OffByDefault(t *testing.T) {
	a := newResizeApp(false)
	before := a.sidebar.Width()
	border := a.layout.SidebarEnd() - 1

	a.Update(tea.MouseClickMsg{X: border, Y: 5, Button: tea.MouseLeft})
	a.Update(tea.MouseMotionMsg{X: border + 10, Y: 5, Button: tea.MouseLeft})
	a.Update(tea.MouseReleaseMsg{X: border + 10, Y: 5, Button: tea.MouseLeft})

	if got := a.sidebar.Width(); got != before {
		t.Errorf("width changed to %d with mouse_resize off", got)
	}
}
