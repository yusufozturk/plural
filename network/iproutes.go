package network

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func IPRoutes() {
	m := data.PluralJSON

	iprteBin := exec.Command("ls", "/sbin/ip")
	iprteBinOut, err := iprteBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	iproute := exec.Command("ip", "route")
	iprteOut, err := iproute.Output()
	iprteStr := string(iprteOut)
	iprteSlice := strings.Split(iprteStr, ",")

	if string(iprteBinOut) != "" {
		m["IPRoute"] = iprteSlice
	}
}
