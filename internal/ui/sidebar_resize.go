package ui

import tea "charm.land/bubbletea/v2"

// reduceSidebarResize drags the sidebar's right border when
// [sidebar] mouse_resize is on: a left press on the border starts the
// drag, motion moves the border to the pointer, release saves the width
// like the [ and ] keys do. It claims mouse messages only while it owns
// the gesture, so clicks elsewhere and text-selection drags pass on.
var reduceSidebarResize reducerFunc = func(a *App, msg tea.Msg) (tea.Cmd, bool) {
	switch m := msg.(type) {
	case tea.MouseClickMsg:
		if !a.sidebarMouseResize || !a.sidebarVisible || m.Button != tea.MouseLeft ||
			a.mode.IsModalOverlay() || a.bootstrap.IsLoading() {
			return nil, false
		}
		if m.X != a.layout.SidebarEnd()-1 || m.Y >= a.height-1 {
			return nil, false
		}
		a.resizingSidebar = true
		return nil, true
	case tea.MouseMotionMsg:
		if !a.resizingSidebar {
			return nil, false
		}
		a.sidebar.SetWidth(m.X - a.layout.RailWidth() - 1)
		return nil, true
	case tea.MouseReleaseMsg:
		if !a.resizingSidebar {
			return nil, false
		}
		a.resizingSidebar = false
		if a.settings != nil {
			a.settings.SaveSidebarWidth(a.sidebar.Width())
		}
		return nil, true
	}
	return nil, false
}
