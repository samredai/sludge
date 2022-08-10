package views

import (
	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/api"
)

// Update a SLURM status display using a result from a ping of the cluster
func SetSlurmStatus(status api.SlurmPingResult, textview *tview.TextView) {
	var statusMessage string
	if status.Status == -1 {
		statusMessage = "⚪ Getting SLURM status"
	} else if status.Status == 0 {
		statusMessage = "🟢 SLURM Cluster Responsive"
	} else if status.Status == 1 {
		statusMessage = "🔴 SLURM Cluster Unresponsive"
	} else {
		statusMessage = "🟠 Unable to detect SLURM cluster status"
	}
	textview.SetTextAlign(tview.AlignLeft).
		SetText(statusMessage + " (" + status.PingDuration.String() + ")")
}
