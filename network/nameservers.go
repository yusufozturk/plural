package network

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func DNS(d *data.PluralJSON) {

	dnsFile := exec.Command("ls", "/etc/resolv.conf")
	dnsFileOut, err := dnsFile.Output()
	if err != nil {
		return
	}
	dnsGrep := exec.Command("grep", "nameserver", "/etc/resolv.conf")
	dnsAwk := exec.Command("awk", "{print$2}")
	dnsGrepOut, err := dnsGrep.StdoutPipe()
	if err != nil {
		return
	}
	dnsGrep.Start()
	dnsAwk.Stdin = dnsGrepOut
	dnsOut, err := dnsAwk.Output()
	if err != nil {
		return
	}
	dnsStr := string(dnsOut)
	dnsSlice := strings.Split(strings.TrimSpace(dnsStr), "\n")
	if string(dnsFileOut) != "" {
		d.DNSNameserver = dnsSlice
	}
}
