package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func IPRoutes(d *data.PluralJSON) {

	if runtime.GOOS == "linux" {
		iproute := exec.Command("ip", "route")
		iprteOut, err := iproute.Output()
		if err != nil {
			return
		}
		iprteStr := string(iprteOut)
		iprteSlice := strings.Split(strings.TrimSpace(iprteStr), "\n")

		if iprteSlice != nil {
			d.IPRoute = iprteSlice
		}
	}

	if runtime.GOOS == "darwin" {
		net := exec.Command("netstat", "-rn")
		netGrep := exec.Command("grep", "-v", "Internet\\|Routing\\|Destination\\|^$")
		netOut, err := net.StdoutPipe()
		if err != nil {
			return
		}
		net.Start()
		netGrep.Stdin = netOut
		netsOut, err := netGrep.Output()
		if err != nil {
			return
		}
		netsStr := string(netsOut)
		netsSlice := strings.Split(strings.TrimSpace(netsStr), "\n")

		if netsSlice != nil {
			d.IPRoute = netsSlice
		}
	}
}
