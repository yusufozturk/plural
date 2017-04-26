package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Deb() {
	m := data.PluralJSON

	dpkgBin := exec.Command("ls", "/usr/bin/dpkg")
	dpkgBinOut, _ := dpkgBin.Output()
	dpkg := exec.Command("dpkg", "-l")
	dpkgAwk := exec.Command("awk", "/^[a-z]/{print$2\"-\"$3}")
	dpkgLOut, _ := dpkg.StdoutPipe()
	dpkg.Start()
	dpkgAwk.Stdin = dpkgLOut
	dpkgOut, _ := dpkgAwk.Output()
	dpkgStr := string(dpkgOut)
	dpkgSlice := strings.Split(strings.TrimSpace(dpkgStr), "\n")

	if string(dpkgBinOut) != "" {
		m["Packages"] = dpkgSlice
	}
}
