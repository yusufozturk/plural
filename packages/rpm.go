package packages

import (
	"fmt"
	"os/exec"
	"plural/data"
	"strings"
)

func Rpm() {

	m := data.PluralJSON

	rpmBin := exec.Command("ls", "/bin/rpm")
	rpmBinOut, err := rpmBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	rpmqa := exec.Command("rpm", "-qa")
	rpmSort := exec.Command("sort")
	rpmqaOut, err := rpmqa.StdoutPipe()
	rpmqa.Start()
	rpmSort.Stdin = rpmqaOut
	rpmOut, err := rpmSort.Output()
	rpmStr := string(rpmOut)
	rpmSlice := strings.Split(rpmStr, "\n")

	if string(rpmBinOut) != "" {
		m["Packages"] = rpmSlice
	}

}
