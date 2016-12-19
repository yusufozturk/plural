package packages

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Pip() {

	m := data.PluralJSON

	pipBin := exec.Command("which", "pip")
	pipBinOut, err := pipBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	pipFree := exec.Command("pip", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	pipFree.Start()
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		fmt.Println(err)
	}
	pipStr := string(pipOut)
	pipReplace := strings.Replace(pipStr, "==", "-", -1)
	pipSlice := strings.Split(pipReplace, "\n")

	if string(pipBinOut) != "" {
		m["Pip"] = pipSlice
	}

}
