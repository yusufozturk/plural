package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Snaps(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		snap := exec.Command("snap", "list")
		snapAwk := exec.Command("awk", "/^[a-z]/{print$1\"-\"$2}")
		snapOut, err := snap.StdoutPipe()
		if err != nil {
			return
		}
		snap.Start()
		snapAwk.Stdin = snapOut
		snapLOut, err := snapAwk.Output()
		if err != nil {
			return
		}
		snapStr := string(snapLOut)
		snapSlice := strings.Split(strings.TrimSpace(snapStr), "\n")

		if snapSlice != nil {
			d.Snaps = snapSlice
		}
	}
}
