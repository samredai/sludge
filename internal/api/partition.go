package api

import (
	"fmt"
	"os"
	"os/exec"
)

// Run the find-best-partition tool for a given script
func FindBestPartition(script string) (string, error) {
	if _, err := os.Stat(script); err == nil {
		out, err := exec.Command("find-best-partition", "-f", script, "-o", "check").Output()
		if err != nil {
			return "", fmt.Errorf("Unexpected output format of find-best-partition")
		}
		return string(out), nil
	} else {
		return "", err
	}
}
