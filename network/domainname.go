package network

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
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
