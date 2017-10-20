package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Pip(d *data.PluralJSON) {

	pipFree := exec.Command("pip", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		return
	}
	pipFree.Start()
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		return
	}
	pipStr := string(pipOut)
	pipReplace := strings.Replace(pipStr, "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if pipSlice != nil {
		d.Pip = pipSlice
	}

}
