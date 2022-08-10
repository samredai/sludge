package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Set colors for a grid that match the system's terminal colors
func SetGridDefaultColors(grid *tview.Grid) {
	grid.SetBackgroundColor(0)
	grid.SetBorderColor(0)
	grid.SetTitleColor(0)
}

// Set colors for a text view that match the system's terminal colors
func SetTextViewDefaultColors(t *tview.TextView) {
	t.SetBackgroundColor(0)
	t.SetBorderColor(0)
	t.SetTextColor(0)
	t.SetTitleColor(0)
}

// Set colors for a list that match the system's terminal colors
func SetListDefaultColors(l *tview.List) {
	l.SetBackgroundColor(0)
	l.SetBorderColor(0)
	l.SetMainTextColor(0)
	l.SetSecondaryTextColor(tcell.Color16)
	l.SetShortcutColor(tcell.Color17)
	l.SetTitleColor(0)
}

// Set colors for a form that match the system's terminal colors
func SetFormDefaultColors(f *tview.Form) {
	f.SetBackgroundColor(0)
	f.SetButtonBackgroundColor(0)
	f.SetButtonTextColor(0)
	f.SetFieldBackgroundColor(0)
	f.SetFieldTextColor(0)
	f.SetLabelColor(0)
	f.SetTitleColor(0)
	f.SetBorder(true)
	f.SetBorderColor(tcell.ColorGreen)
}

// Set colors for a table that match the system's terminal colors
func SetTableDefaultColors(t *tview.Table) {
	var style tcell.Style
	style = tcell.StyleDefault.Attributes(tcell.AttrReverse)
	t.SetBackgroundColor(0)
	t.SetBorderColor(0)
	t.SetBordersColor(0)
	t.SetTitleColor(0)
	t.SetSelectedStyle(style)
}
