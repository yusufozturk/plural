package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Deb(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		dpkgBin := exec.Command("ls", "/usr/bin/dpkg")
		dpkgBinOut, err := dpkgBin.Output()
		if err != nil {
			return
		}
		dpkg := exec.Command("dpkg", "-l")
		dpkgAwk := exec.Command("awk", "/^[a-z]/{print$2\"-\"$3}")
		dpkgLOut, err := dpkg.StdoutPipe()
		if err != nil {
			return
		}
		dpkg.Start()
		dpkgAwk.Stdin = dpkgLOut
		dpkgOut, err := dpkgAwk.Output()
		if err != nil {
			return
		}
		dpkgStr := string(dpkgOut)
		dpkgSlice := strings.Split(strings.TrimSpace(dpkgStr), "\n")

		if string(dpkgBinOut) != "" {
			d.Packages = dpkgSlice
		}
	}
}
