package sidebar

import (
	"strings"
	"testing"
)

func TestChannelIndent_ShiftsSectionRowsOnly(t *testing.T) {
	m := New([]ChannelItem{
		{ID: "C1", Name: "a_rather_long_channel_name", Type: "channel"},
		{ID: "D1", Name: "priya", Type: "dm"},
	})
	m.ToggleCollapse("Channels")
	m.SetChannelIndent(3)
	view := m.View(12, 30)
	plain := func(row string) string { return ansiRe.ReplaceAllString(row, "") }

	ch := plain(rowFor(t, view, "a_rather"))
	if !strings.HasPrefix(ch, "     # a_rather") {
		t.Errorf("channel row not indented by 3: %q", ch)
	}
	if n := len([]rune(ch)); n != 28 {
		t.Errorf("indented row is %d cells, want 28 (name must truncate): %q", n, ch)
	}
	if hdr := plain(rowFor(t, view, "Channels")); strings.HasPrefix(hdr, "     ") {
		t.Errorf("section header indented too: %q", hdr)
	}
	if th := plain(rowFor(t, view, "Threads")); strings.HasPrefix(th, "     ") {
		t.Errorf("Threads row indented: %q", th)
	}
}

func TestChannelIndent_Clamped(t *testing.T) {
	m := New(nil)
	m.SetChannelIndent(40)
	if m.channelIndent != maxChannelIndent {
		t.Errorf("indent = %d, want clamp to %d", m.channelIndent, maxChannelIndent)
	}
	m.SetChannelIndent(-2)
	if m.channelIndent != 0 {
		t.Errorf("indent = %d, want 0", m.channelIndent)
	}
}
