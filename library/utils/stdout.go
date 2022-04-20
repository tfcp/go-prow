package utils

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func TableStdout(title []string, contents [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	for _, content := range contents {
		table.Append(content)
	}
	table.Render()
}
