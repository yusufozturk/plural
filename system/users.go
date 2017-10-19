package system

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func UsersLoggedIn(d *data.PluralJSON) {
	wBin := exec.Command("ls", "/usr/bin/w")
	wBinOut, err := wBin.Output()
	if err != nil {
		return
	}
	wh := exec.Command("w", "-h")
	whAwk := exec.Command("awk", "{print$1\"-\"$2}")
	whOut, err := wh.StdoutPipe()
	if err != nil {
		return
	}
	wh.Start()
	whAwk.Stdin = whOut
	wOut, err := whAwk.Output()
	if err != nil {
		return
	}
	whStr := string(wOut)
	wSlice := strings.Split(strings.TrimSpace(whStr), "\n")

	if string(wBinOut) != "" {
		d.UsersLoggedin = wSlice
	}
}

func Users(d *data.PluralJSON) {
	passGrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
	passGrepOut, err := passGrep.Output()
	if err != nil {
		return
	}
	passStr := string(passGrepOut)
	usersSlice := strings.Split(strings.TrimSpace(passStr), "\n")
	d.Users = usersSlice
}
