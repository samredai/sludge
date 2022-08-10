package main

import (
	"context"
	"flag"
	"time"

	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/controllers"
	"github.com/samredai/sludge/internal/models"
	"github.com/samredai/sludge/internal/views"
)

type SludgeApp struct {
	TView *tview.Application
	Pages *tview.Pages
}

func main() {
	app := tview.NewApplication()

	// Parse arguments
	username := flag.String("username", "", "A username for filtering SLUDGE data")
	account := flag.String("account", "", "An account for filtering SLUDGE data")
	refreshInterval := flag.Duration("refresh-interval", 60*time.Second, "Interval in seconds for refreshing dashboards")
	flag.Parse()

	ctx := context.Background()
	ctx = context.WithValue(ctx, models.USERNAME_KEY, *username)
	ctx = context.WithValue(ctx, models.ACCOUNT_KEY, *account)
	ctx = context.WithValue(ctx, models.APP_KEY, app)

	// Get all jobs from the database and populate the jobs table
	jobs := models.GetJobs(ctx)

	jobsTable := views.CreateJobsTable(jobs)
	jobDetails := views.CreateJobDetails()

	nodes := models.GetAllNodes()
	nodesTable := views.CreateNodesTable(nodes)
	nodeDetails := views.CreateNodeDetails()

	statusView := views.CreateStatusView()

	reportsView := views.CreateReportsView()

	pages := tview.NewPages()
	ctx = context.WithValue(ctx, models.PAGE_HANDLER_KEY, pages)

	jobPage := views.JobPage(ctx, statusView, jobsTable, jobDetails)
	nodePage := views.NodePage(ctx, statusView, nodesTable, nodeDetails)
	reportsPage := views.ReportsPage(ctx, statusView, reportsView)

	pages.AddPage(views.JobPageID, jobPage, true, true)
	pages.AddPage(views.NodePageID, nodePage, true, false)
	pages.AddPage(views.ReportsPageID, reportsPage, true, false)

	// Configure the responses to pressing certain keys
	controllers.ConfigureKeys(app, pages)
	// Configure the responses to selecting things on the screen
	controllers.ConfigureJobSelections(jobsTable, jobPage, jobDetails)
	controllers.ConfigureNodeSelections(nodesTable, nodePage, nodeDetails)

	// Important that the table's title is set last since this returns a Box
	jobsTable.SetBorder(true).SetTitle("Jobs")
	nodesTable.SetBorder(true).SetTitle("Nodes")

	// Periodically refresh the jobs table and redraw the app
	go views.RefreshSlurmStatusView(ctx, app, statusView, *refreshInterval)
	go views.RefreshJobsTable(ctx, app, jobsTable, *refreshInterval)
	go views.RefreshNodeTable(app, nodesTable, *refreshInterval)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
