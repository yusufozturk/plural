package system

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func UsersLoggedIn() {
	m := data.PluralJSON
	wBin := exec.Command("ls", "/usr/bin/w")
	wBinOut, _ := wBin.Output()
	wh := exec.Command("w", "-h")
	whAwk := exec.Command("awk", "{print$1\"-\"$2}")
	whOut, _ := wh.StdoutPipe()
	wh.Start()
	whAwk.Stdin = whOut
	wOut, _ := whAwk.Output()
	whStr := string(wOut)
	wSlice := strings.Split(strings.TrimSpace(whStr), "\n")

	if string(wBinOut) != "" {
		m["UsersLoggedin"] = wSlice
	}
}

func Users() {
	m := data.PluralJSON
	passGrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
	passGrepOut, _ := passGrep.Output()
	passStr := string(passGrepOut)
	usersSlice := strings.Split(strings.TrimSpace(passStr), "\n")
	m["Users"] = usersSlice
}
