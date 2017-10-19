package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Rpm(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		rpmBin := exec.Command("ls", "/bin/rpm")
		rpmBinOut, err := rpmBin.Output()
		if err != nil {
			return
		}
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

		if string(rpmBinOut) != "" {
			d.Packages = rpmSlice
		}
	}
}
