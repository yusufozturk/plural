package packages

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Gem() {
	m := data.PluralJSON

	gemBin := exec.Command("which", "gem")
	gemBinOut, err := gemBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	gemList := exec.Command("gem", "list")
	gemGrep := exec.Command("grep", "^[a-zA-Z]")
	gemListOut, err := gemList.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	gemList.Start()
	gemGrep.Stdin = gemListOut
	gemOut, err := gemGrep.Output()
	gemStr := string(gemOut)
	gemReplace := strings.Replace(gemStr, " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(gemReplace2, "\n")

	if string(gemBinOut) != "" {
		m["Gem"] = gemSlice
	}
}
