package network

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func DNS() {
	m := data.PluralJSON

	dnsFile := exec.Command("ls", "/etc/resolv.conf")
	dnsFileOut, err := dnsFile.Output()
	if err != nil {
		fmt.Println(dnsFileOut)
	}
	dnsGrep := exec.Command("grep", "nameserver", "/etc/resolv.conf")
	dnsAwk := exec.Command("awk", "{print$2}")
	dnsGrepOut, err := dnsGrep.StdoutPipe()
	if err != nil {
		fmt.Println(dnsGrepOut)
	}
	dnsGrep.Start()
	dnsAwk.Stdin = dnsGrepOut
	dnsOut, err := dnsAwk.Output()
	if err != nil {
		fmt.Println(dnsOut)
	}
	dnsStr := string(dnsOut)
	dnsSlice := strings.Split(strings.TrimSpace(dnsStr), "\n")
	if string(dnsFileOut) != "" {
		m["DNSNameserver"] = dnsSlice
	}
}
