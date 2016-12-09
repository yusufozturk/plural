package network

import (
	"fmt"
	"os/exec"
	"plural/data"
	"strings"
)

func DomainName() {
	m := data.PluralJSON

	hostname := exec.Command("hostname", "-f")
	hostcut := exec.Command("cut", "-d.", "-f", "2-")
	hostnameOut, err := hostname.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	hostname.Start()
	hostcut.Stdin = hostnameOut
	domain, err := hostcut.Output()
	domainStr := string(domain)
	domainTrim := strings.TrimSpace(domainStr)

	m["Domain"] = domainTrim
}
