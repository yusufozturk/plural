package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func UsersLoggedIn(d *data.PluralJSON) {
	if runtime.GOOS != "windows" {
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

		if wSlice != nil {
			d.UsersLoggedin = wSlice
		}
	}
}

func Users(d *data.PluralJSON) {
	if runtime.GOOS != "windows" {
		passGrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
		passGrepOut, err := passGrep.Output()
		if err != nil {
			return
		}
		passStr := string(passGrepOut)
		usersSlice := strings.Split(strings.TrimSpace(passStr), "\n")
		if usersSlice != nil {
			d.Users = usersSlice
		}
	}
}
