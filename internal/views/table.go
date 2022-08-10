package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/samredai/sludge/internal/models"
)

// Given a collection and a table, populate the table with the items from the collection
func PopulateTable[T any](c *models.Collection[T], table *tview.Table) {
	table.Clear()
	if len(c.Items) > 0 {
		for column, text := range c.Headers(c.Items[0]) {
			table.SetCell(0, column, &tview.TableCell{
				Text:       text,
				Align:      tview.AlignCenter,
				Attributes: tcell.AttrBold,
			})
		}
		for row, item := range c.Items {
			ref := c.Row(item)[0]
			for column, text := range c.Row(item) {
				attr := tcell.AttrNone
				if column == 0 {
					attr = tcell.AttrBold
				}
				align := tview.AlignLeft
				if column == 0 || column >= 4 {
					align = tview.AlignRight
				}
				table.SetCell(row+1, column, &tview.TableCell{
					Text:       text,
					Color:      tcell.ColorDefault,
					Align:      align,
					Reference:  ref,
					Attributes: attr,
				})
			}
		}
	}
}
