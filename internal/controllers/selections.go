package controllers

import (
	"strconv"

	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/models"
	"github.com/samredai/sludge/internal/views"
)

// Handle selecting a job from the jobs table
func ConfigureJobSelections(jobstable *tview.Table, grid *tview.Grid, jobDetails *tview.TextView) {
	// When a job is selected in the jobs table, update the job details on
	// the main section of the grid
	jobstable.SetSelectionChangedFunc(func(row int, column int) {
		c := jobstable.GetCell(row, 0)
		if ref, ok := c.GetReference().(string); ok {
			// Remove the current job details from the grid
			grid.RemoveItem(jobDetails)
			// Create a new text view and populate it with the jobs details
			jobDetails = tview.NewTextView()
			jobID, err := strconv.Atoi(ref)
			if err == nil {
				jobinfo := models.GetJobInfo(jobID)
				jobDetails.SetTextAlign(tview.AlignLeft).SetDynamicColors(true).SetText("[green]" + jobinfo)
				views.SetTextViewDefaultColors(jobDetails)
				// Add the new job details to the main section of the grid
				grid.AddItem(jobDetails, 1, 1, 1, 2, 0, 100, false)
			}
		}
	})
}

// Handle selecting a node from the nodes table
func ConfigureNodeSelections(nodestable *tview.Table, grid *tview.Grid, nodeDetails *tview.TextView) {
	// When a node is selected in the nodes table, update the node details
	nodestable.SetSelectionChangedFunc(func(row int, column int) {
		c := nodestable.GetCell(row, 0)
		if ref, ok := c.GetReference().(string); ok {
			// Remove the current node details from the grid
			grid.RemoveItem(nodeDetails)
			// Create a new text view and populate it with the nodes details
			nodeDetails = tview.NewTextView()
			nodeinfo := models.GetNodeInfo(ref)
			nodeDetails.SetTextAlign(tview.AlignLeft).SetDynamicColors(true).SetText("[green]" + nodeinfo)
			views.SetTextViewDefaultColors(nodeDetails)
			// Add the new node details to the main section of the grid
			grid.AddItem(nodeDetails, 1, 1, 1, 2, 0, 100, false)

		}
	})
}
