package views

import (
	"context"

	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/models"
)

// Run a specific SLURM report and add the results to a table
func PopulateReportTable(ctx context.Context, group string, report string, t *tview.Table) {
	SetTableDefaultColors(t)
	switch reportName := group + " " + report; reportName {
	case "Cluster Utilization":
		c := models.GetClusterUtilizationReport(ctx)
		PopulateTable(c, t)
	case "Cluster AccountUtilizationByUser":
		c := models.GetClusterAccountUtilizationByUserReport(ctx)
		PopulateTable(c, t)
	case "Cluster UserUtilizationByAccount":
		c := models.GetClusterUserUtilizationByAccountReport(ctx)
		PopulateTable(c, t)
	case "Job SizesByAccount":
		c := models.GetJobSizesByAccountReport(ctx)
		PopulateTable(c, t)
	case "Reservation Utilization":
		c := models.GetReservationUtilizationReport(ctx)
		PopulateTable(c, t)
	case "User TopUsage":
		c := models.GetUserTopUsageReport(ctx)
		PopulateTable(c, t)
	default:
		c := models.GetClusterUtilizationReport(ctx)
		PopulateTable(c, t)
	}
}
