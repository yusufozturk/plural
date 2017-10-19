package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func IPTables(d *data.PluralJSON) {

	if runtime.GOOS == "linux" {
		iptableBin := exec.Command("ls", "/sbin/iptables")
		iptableBinOut, err := iptableBin.Output()
		if err != nil {
			return
		}
		iptableL := exec.Command("iptables", "-L")
		iptableGrep := exec.Command("grep", "-v", "^Chain\\|target\\|^$")
		iptableLOut, err := iptableL.StdoutPipe()
		if err != nil {
			return
		}
		iptableL.Start()
		iptableGrep.Stdin = iptableLOut
		iptableOut, err := iptableGrep.Output()
		if err != nil {
			return
		}
		iptableStr := string(iptableOut)
		iptableSlice := strings.Split(strings.TrimSpace(iptableStr), "\n")

		if string(iptableBinOut) != "" {
			d.Iptables = iptableSlice
		}
	}
}
