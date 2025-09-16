package tables

import (	
	"github.com/jedib0t/go-pretty/v6/table"	
	"os"
)

func PrintTable (tableRowHeader table.Row,tableCaption string ,rows []table.Row)  {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(rows)
	tw.SetCaption(tableCaption)
	tw.Render()
}
