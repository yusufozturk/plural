package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Rpm(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		rpmqa := exec.Command("rpm", "-qa")
		rpmSort := exec.Command("sort")
		rpmqaOut, err := rpmqa.StdoutPipe()
		if err != nil {
			return
		}
		rpmqa.Start()
		rpmSort.Stdin = rpmqaOut
		rpmOut, err := rpmSort.Output()
		if err != nil {
			return
		}
		rpmStr := string(rpmOut)
		rpmSlice := strings.Split(strings.TrimSpace(rpmStr), "\n")

		if rpmSlice != nil {
			d.Packages = rpmSlice
		}
	}
}
