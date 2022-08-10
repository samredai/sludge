package views

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/models"
)

// Creates the footer view that displays pages and the shortcut keys to switch to them
func footer() *tview.TextView {
	s := "[green]"
	s += "[Ctrl+J: Jobs]   "
	s += "[Ctrl+N: Nodes]   "
	s += "[Ctrl+R: Reports]   "
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetText(s)
	SetTextViewDefaultColors(tv)
	return tv
}

// Creates the form for submitting a request to find the best partition for a given script
func FindBestPartitionForm(ctx context.Context) *tview.Flex {
	app := ctx.Value(models.APP_KEY).(*tview.Application)
	pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)

	findBestPartitionForm := tview.NewForm().AddInputField("Script", "", 200, nil, nil)
	findBestPartitionForm.AddButton("Find Best Partition", func() {
		pages.RemovePage(FindBestPartitionFormID)
		app.Suspend(func() {
			script := findBestPartitionForm.GetFormItemByLabel("Script").(*tview.InputField).GetText()

			cmd := exec.Command("find-best-partition", "-f", script, "-o", "check")
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			_ = cmd.Run()

			waitMessage := exec.Command("echo", "\n\nPress Enter to return to sludge...")
			waitMessage.Stdout = os.Stdout
			waitMessage.Stdin = os.Stdin
			_ = waitMessage.Run()

			wait := exec.Command("read", "-p", "")
			wait.Stdout = os.Stdout
			wait.Stdin = os.Stdin
			_ = wait.Run()
		})
	}).AddButton("Close", func() {
		pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)
		pages.RemovePage(FindBestPartitionFormID)
	})
	findBestPartitionForm.SetBorder(true).SetTitle("Find the best partition to submit a script").SetTitleAlign(tview.AlignLeft)
	SetFormDefaultColors(findBestPartitionForm)

	formWindow := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(findBestPartitionForm, 0, 1, true).
			AddItem(nil, 0, 1, false), 0, 1, true).
		AddItem(nil, 0, 1, false)
	return formWindow
}

// Creates the form for submitting a script to the SLURM cluster using sbatch
func SubmitScript(ctx context.Context) *tview.Flex {
	app := ctx.Value(models.APP_KEY).(*tview.Application)
	pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)

	submitScriptForm := tview.NewForm().AddInputField("Script", "", 200, nil, nil)
	submitScriptForm.AddButton("Submit", func() {
		pages.RemovePage(SubmitScriptFormID)
		app.Suspend(func() {
			script := submitScriptForm.GetFormItemByLabel("Script").(*tview.InputField).GetText()
			args := strings.Fields(script)
			cmd := exec.Command("sbatch", args...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			_ = cmd.Run()

			waitMessage := exec.Command("echo", "\n\nPress Enter to return to sludge...")
			waitMessage.Stdout = os.Stdout
			waitMessage.Stdin = os.Stdin
			_ = waitMessage.Run()

			wait := exec.Command("read", "-p", "")
			wait.Stdout = os.Stdout
			wait.Stdin = os.Stdin
			_ = wait.Run()
		})
	}).AddButton("Close", func() {
		pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)
		pages.RemovePage(SubmitScriptFormID)
	})
	submitScriptForm.SetBorder(true).SetTitle("Submit a script to SLURM").SetTitleAlign(tview.AlignLeft)
	SetFormDefaultColors(submitScriptForm)

	formWindow := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(submitScriptForm, 0, 1, true).
			AddItem(nil, 0, 1, false), 0, 1, true).
		AddItem(nil, 0, 1, false)
	return formWindow
}

// Create the action pane which includes forms to run common actions
func ActionPane(ctx context.Context) *tview.List {
	actionPane := tview.NewList().ShowSecondaryText(false).
		AddItem("Submit Script", "", 's', func() {

			pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)
			formWindow := SubmitScript(ctx)

			pages.AddPage(SubmitScriptFormID, formWindow, true, true)
		}).
		AddItem("Find Best Partition", "", 'f', func() {

			pages := ctx.Value(models.PAGE_HANDLER_KEY).(*tview.Pages)
			formWindow := FindBestPartitionForm(ctx)

			pages.AddPage(FindBestPartitionFormID, formWindow, true, true)
		})
	SetListDefaultColors(actionPane)
	return actionPane
}

// Create the job page which includes active jobs and details about them
func JobPage(ctx context.Context, statusView *tview.TextView, left *tview.Table, center *tview.TextView) *tview.Grid {
	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(Logo + RepoURL)
	SetTextViewDefaultColors(header)
	topright := ActionPane(ctx)
	grid := tview.NewGrid().
		SetRows(8, 0, 1).
		SetColumns(60, 0, 0).
		SetBorders(true).
		AddItem(statusView, 0, 0, 1, 1, 0, 0, false).
		AddItem(header, 0, 1, 1, 1, 0, 0, false).
		AddItem(topright, 0, 2, 1, 1, 0, 0, false).
		AddItem(footer(), 2, 0, 1, 3, 0, 0, false)
	grid.AddItem(left, 1, 0, 1, 1, 0, 100, false).
		AddItem(center, 1, 1, 1, 2, 0, 100, false)
	SetGridDefaultColors(grid)
	return grid
}

// Create the node page which includes SLURM cluster nodes and details about them
func NodePage(ctx context.Context, statusView *tview.TextView, left *tview.Table, center *tview.TextView) *tview.Grid {

	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(Logo + RepoURL)
	SetTextViewDefaultColors(header)
	topright := ActionPane(ctx)
	grid := tview.NewGrid().
		SetRows(8, 0, 1).
		SetColumns(60, 0, 0).
		SetBorders(true).
		AddItem(statusView, 0, 0, 1, 1, 0, 0, false).
		AddItem(header, 0, 1, 1, 1, 0, 0, false).
		AddItem(topright, 0, 2, 1, 1, 0, 0, false).
		AddItem(footer(), 2, 0, 1, 3, 0, 0, false)
	grid.AddItem(left, 1, 0, 1, 1, 0, 100, false).
		AddItem(center, 1, 1, 1, 2, 0, 100, false)
	SetGridDefaultColors(grid)
	return grid
}

// Create the reports page which allows running various reports against the SLURM cluster
func ReportsPage(ctx context.Context, statusView *tview.TextView, reportView *tview.TextView) *tview.Grid {
	reportTable := tview.NewTable()

	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(Logo + RepoURL)
	SetTextViewDefaultColors(header)
	topright := ActionPane(ctx)
	grid := tview.NewGrid().
		SetRows(8, 0, 1).
		SetColumns(60, 0).
		SetBorders(true).
		AddItem(statusView, 0, 0, 1, 1, 0, 0, false).
		AddItem(header, 0, 1, 1, 1, 0, 0, false).
		AddItem(topright, 0, 2, 1, 1, 0, 0, false).
		AddItem(footer(), 2, 0, 1, 3, 0, 0, false).
		AddItem(reportView, 1, 1, 1, 2, 0, 100, false)
	SetGridDefaultColors(grid)

	list := tview.NewList().ShowSecondaryText(false).
		AddItem("Cluster Utilization", "", 'c', func() {
			PopulateReportTable(ctx, "Cluster", "Utilization", reportTable)
			reportTable.SetBorder(true).SetTitle("Cluster Utilization")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Cluster Account Utilization By User", "", 'a', func() {
			PopulateReportTable(ctx, "Cluster", "AccountUtilizationByUser", reportTable)
			reportTable.SetBorder(true).SetTitle("Cluster AccountUtilizationByUser")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Cluster User Utilization By Account", "", 'u', func() {
			PopulateReportTable(ctx, "Cluster", "UserUtilizationByAccount", reportTable)
			reportTable.SetBorder(true).SetTitle("Cluster UserUtilizationByAccount")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Job Sizes By Account", "", 'j', func() {
			PopulateReportTable(ctx, "Job", "SizesByAccount", reportTable)
			reportTable.SetBorder(true).SetTitle("Job SizesByAccount")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Job Sizes By Account And WcKey", "", 'e', func() {
			PopulateReportTable(ctx, "Job", "SizesByAccountAndWcKey", reportTable)
			reportTable.SetBorder(true).SetTitle("Job SizesByAccountAndWcKey")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Job Sizes By Wckey", "", 'l', func() {
			PopulateReportTable(ctx, "Job", "SizesByWckey", reportTable)
			reportTable.SetBorder(true).SetTitle("Job SizesByWckey")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("Reservation Utilization", "", 'r', func() {
			PopulateReportTable(ctx, "Reservation", "Utilization", reportTable)
			reportTable.SetBorder(true).SetTitle("Reservation Utilization")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		}).
		AddItem("User TopUsage", "", 't', func() {
			var app *tview.Application
			app = ctx.Value(models.APP_KEY).(*tview.Application)
			go app.Draw()
			PopulateReportTable(ctx, "User", "TopUsage", reportTable)
			reportTable.SetBorder(true).SetTitle("User TopUsage")
			grid.RemoveItem(reportView).AddItem(reportTable, 1, 1, 1, 2, 0, 100, false)
		})
	SetListDefaultColors(list)
	grid.AddItem(list, 1, 0, 1, 1, 0, 0, false)

	return grid
}
