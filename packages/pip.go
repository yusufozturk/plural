package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Pip() {

	m := data.PluralJSON

	pipBin := exec.Command("which", "pip")
	pipBinOut, _ := pipBin.Output()
	pipFree := exec.Command("pip", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, _ := pipFree.StdoutPipe()
	pipFree.Start()
	pipSort.Stdin = pipFreeOut
	pipOut, _ := pipSort.Output()
	pipStr := string(pipOut)
	pipReplace := strings.Replace(pipStr, "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if string(pipBinOut) != "" {
		m["Pip"] = pipSlice
	}

}
