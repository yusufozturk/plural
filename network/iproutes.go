package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func IPRoutes(d *data.PluralJSON) {

	if runtime.GOOS == "linux" {
		iprteBin := exec.Command("ls", "/sbin/ip")
		iprteBinOut, err := iprteBin.Output()
		if err != nil {
			return
		}
		iproute := exec.Command("ip", "route")
		iprteOut, err := iproute.Output()
		if err != nil {
			return
		}
		iprteStr := string(iprteOut)
		iprteSlice := strings.Split(strings.TrimSpace(iprteStr), "\n")

		if string(iprteBinOut) != "" {
			d.IPRoute = iprteSlice
		}
	}

	// if runtime.GOOS == "darwin" {
	// 	iprteBin := exec.Command("ls", "/usr/sbin/netstat")
	// 	iprteBinOut, err := iprteBin.Output()
	// 	if err != nil {
	// 		return
	// 	}
	// 	iproute := exec.Command("netstat", "-rn")
	// 	iprteOut, err := iproute.Output()
	// 	if err != nil {
	// 		return
	// 	}
	// 	iprteStr := string(iprteOut)
	// 	iprteSlice := strings.Split(strings.TrimSpace(iprteStr), "\n")

	// 	if string(iprteBinOut) != "" {
	// 		d.IPRoute = iprteSlice
	// 	}
	// }
}
