package network

import (
	"os"
	"strings"

	"github.com/marshyski/plural/data"
)

func DomainName(d *data.PluralJSON) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	d.Domain = strings.TrimSuffix(hostname, ".")
}
