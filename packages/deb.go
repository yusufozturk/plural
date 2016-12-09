package packages

import (
	"fmt"
	"os/exec"
	"plural/data"
	"strings"
)

func Deb() {
	m := data.PluralJSON

	dpkgBin := exec.Command("ls", "/usr/bin/dpkg")
	dpkgBinOut, err := dpkgBin.Output()
	dpkg := exec.Command("dpkg", "-l")
	dpkgAwk := exec.Command("awk", "/^[a-z]/{print$2\"-\"$3}")
	dpkgLOut, err := dpkg.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	dpkg.Start()
	dpkgAwk.Stdin = dpkgLOut
	dpkgOut, err := dpkgAwk.Output()
	if err != nil {
		fmt.Println(err)
	}
	dpkgStr := string(dpkgOut)
	dpkgSlice := strings.Split(dpkgStr, ",")

	if string(dpkgBinOut) != "" {
		m["Packages"] = dpkgSlice
	}
}
