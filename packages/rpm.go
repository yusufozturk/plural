package packages

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Rpm() {

	m := data.PluralJSON

	rpmBin := exec.Command("ls", "/bin/rpm")
	rpmBinOut, _ := rpmBin.Output()
	rpmqa := exec.Command("rpm", "-qa")
	rpmSort := exec.Command("sort")
	rpmqaOut, _ := rpmqa.StdoutPipe()
	rpmqa.Start()
	rpmSort.Stdin = rpmqaOut
	rpmOut, _ := rpmSort.Output()
	rpmStr := string(rpmOut)
	rpmSlice := strings.Split(strings.TrimSpace(rpmStr), "\n")

	if string(rpmBinOut) != "" {
		m["Packages"] = rpmSlice
	}

}
