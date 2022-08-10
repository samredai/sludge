package views

import (
	"context"
	"time"

	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/models"
)

// Continuously refresh the status view on a given interval
func RefreshSlurmStatusView(ctx context.Context, app *tview.Application, statusView *tview.TextView, refreshInterval time.Duration) {
	for {
		time.Sleep(refreshInterval)
		status := models.GetSlurmStatus()
		SetSlurmStatus(status, statusView)
		app.Draw()
	}
}

// Continuously refresh the jobs table on a given interval
func RefreshJobsTable(ctx context.Context, app *tview.Application, jobsTable *tview.Table, refreshInterval time.Duration) {
	for {
		time.Sleep(refreshInterval)
		jobs := models.GetJobs(ctx)
		PopulateTable(jobs, jobsTable)
		app.Draw()
	}
}

// Continuously refresh the nodes table on a given interval
func RefreshNodeTable(app *tview.Application, nodesTable *tview.Table, refreshInterval time.Duration) {
	for {
		time.Sleep(refreshInterval)
		nodes := models.GetAllNodes()
		PopulateTable(nodes, nodesTable)
		app.Draw()
	}
}
