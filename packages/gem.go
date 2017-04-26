package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Gem() {
	m := data.PluralJSON

	gemBin := exec.Command("which", "gem")
	gemBinOut, _ := gemBin.Output()
	gemList := exec.Command("gem", "list")
	gemGrep := exec.Command("grep", "^[a-zA-Z]")
	gemListOut, _ := gemList.StdoutPipe()
	gemList.Start()
	gemGrep.Stdin = gemListOut
	gemOut, _ := gemGrep.Output()
	gemStr := string(gemOut)
	gemReplace := strings.Replace(gemStr, " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(strings.TrimSpace(gemReplace2), "\n")

	if string(gemBinOut) != "" {
		m["Gem"] = gemSlice
	}
}
