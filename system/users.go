package system

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func UsersLoggedIn() {
	m := data.PluralJSON

	wBin := exec.Command("ls", "/usr/bin/w")
	wBinOut, err := wBin.Output()
	if err != nil {
		fmt.Println(wBinOut)
	}
	wh := exec.Command("w", "-h")
	whAwk := exec.Command("awk", "{print$1\"-\"$2}")
	whOut, err := wh.StdoutPipe()
	if err != nil {
		fmt.Println(whOut)
	}
	wh.Start()
	whAwk.Stdin = whOut
	wOut, err := whAwk.Output()
	if err != nil {
		fmt.Println(wOut)
	}
	whStr := string(wOut)
	wSlice := strings.Split(whStr, ",")

	if string(wBinOut) != "" {
		m["UsersLoggedin"] = wSlice
	}
}

func Users() {
	m := data.PluralJSON

	passFile := exec.Command("ls", "/etc/passwd")
	passFileOut, err := passFile.Output()
	if err != nil {
		fmt.Println(passFileOut)
	}
	passGrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
	passGrepOut, err := passGrep.Output()
	passStr := string(passGrepOut)
	usersSlice := strings.Split(passStr, ",")
	m["Users"] = usersSlice
}
