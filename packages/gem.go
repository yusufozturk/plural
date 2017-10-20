package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Gem(d *data.PluralJSON) {

	gemList := exec.Command("gem", "list")
	gemGrep := exec.Command("grep", "^[a-zA-Z]")
	gemListOut, err := gemList.StdoutPipe()
	if err != nil {
		return
	}
	gemList.Start()
	gemGrep.Stdin = gemListOut
	gemOut, err := gemGrep.Output()
	if err != nil {
		return
	}
	gemStr := string(gemOut)
	gemReplace := strings.Replace(gemStr, " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(strings.TrimSpace(gemReplace2), "\n")

	if gemSlice != nil {
		d.Gem = gemSlice
	}
}
