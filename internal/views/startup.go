package views

import (
	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/api"
	"github.com/samredai/sludge/internal/models"
)

// Initialize a jobs table on application init
func CreateJobsTable(c *models.Collection[api.JobDetails]) *tview.Table {
	jobsTable := tview.NewTable().SetFixed(1, 1).SetSelectable(true, false)
	SetTableDefaultColors(jobsTable)
	PopulateTable(c, jobsTable)
	return jobsTable
}

// Initialize a nodes table on application init
func CreateNodesTable(c *models.Collection[api.NodeInfo]) *tview.Table {
	nodeTable := tview.NewTable().SetFixed(1, 1).SetSelectable(true, false)
	SetTableDefaultColors(nodeTable)
	PopulateTable(c, nodeTable)
	return nodeTable
}

// Initialize the job details view on application init
func CreateJobDetails() *tview.TextView {
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(Introduction)
	SetTextViewDefaultColors(tv)
	return tv
}

// Initialize the node details view on application init
func CreateNodeDetails() *tview.TextView {
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(Introduction)
	SetTextViewDefaultColors(tv)
	return tv
}

// Initialize the report view on application init
func CreateReportsView() *tview.TextView {
	tv := tview.NewTextView()
	tv.SetTextAlign(tview.AlignCenter).
		SetText(api.ReportPageInstructions)
	SetTextViewDefaultColors(tv)
	tv.SetDynamicColors(true)
	tv.SetTextAlign(tview.AlignCenter)
	return tv
}

// Initialize the SLURM cluster status view on application init
func CreateStatusView() *tview.TextView {
	tv := tview.NewTextView()
	SetSlurmStatus(models.GetSlurmStatus(), tv)
	SetTextViewDefaultColors(tv)
	return tv
}
