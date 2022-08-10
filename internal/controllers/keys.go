package controllers

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/views"
)

// Handle key presses
func ConfigureKeys(app *tview.Application, pages *tview.Pages) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		switch key {
		case tcell.KeyCtrlC:
			app.Stop()
		case tcell.KeyCtrlD:
			app.Stop()
		case tcell.KeyCtrlJ:
			pages.SwitchToPage(views.JobPageID)
		case tcell.KeyCtrlN:
			pages.SwitchToPage(views.NodePageID)
		case tcell.KeyCtrlR:
			pages.SwitchToPage(views.ReportsPageID)
		}
		return event
	})
}
