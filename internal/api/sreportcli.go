package api

import (
	"fmt"
	"os/exec"
	"strings"
)

var ReportPageInstructions = `
Select any report on the left to view it here.
[green]
       .--.                   .---.
   .---|__|           .-.     |~~~|
.--|===|--|_          |_|     |~~~|--.
|  |===|  |'\     .---!~|  .--| S |--|
|%%|   |  |.'\    |===| |--|%%| L |  |
|%%|   |  |\.'\   |   | |__|  | U |  |
|  |   |  | \  \  |===| |==|  | R |  |
|  |   |__|  \.'\ |   |_|__|  | M |__|
|  |===|--|   \.'\|===|~|--|%%|~~~|--|
^--^---'--^    '-''---^-^--^--^---'--'
`

// Run a SLURM report for a given group and report name
func GetParsableReport(group string, report string) ([][]string, error) {
	var data [][]string

	if group == "" && report == "" {
		return data, nil
	}
	out, err := exec.Command("sreport", group, report, "--parsable", "--noheader", "--quiet").Output()
	if err != nil {
		return data, fmt.Errorf("Unexpected output format of sreport")
	}

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		if line != "" {
			row := strings.Split(line, "|")
			data = append(data, row)
		}
	}

	return data, nil
}
