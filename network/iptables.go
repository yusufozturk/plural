package network

import (
	"fmt"
	"os/exec"
	"plural/data"
	"strings"
)

func IPTables() {
	m := data.PluralJSON

	iptableBin := exec.Command("ls", "/sbin/iptables")
	iptableBinOut, err := iptableBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	iptableL := exec.Command("iptables", "-L")
	iptableGrep := exec.Command("grep", "-v", "^Chain\\|target\\|^$")
	iptableLOut, err := iptableL.StdoutPipe()
	iptableL.Start()
	iptableGrep.Stdin = iptableLOut
	iptableOut, err := iptableGrep.Output()
	iptableStr := string(iptableOut)
	iptableSlice := strings.Split(iptableStr, ",")

	if string(iptableBinOut) != "" {
		m["Iptables"] = iptableSlice
	}
}
