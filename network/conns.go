package network

import (
	"bytes"
	"runtime"

	"github.com/marshyski/plural/data"

	"github.com/drael/GOnetstat"
)

func Conns(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		gonetstat4 := GOnetstat.Tcp()
		tcp4 := []string{}
		for _, nettcp := range gonetstat4 {
			if nettcp.State == "LISTEN" {
				ipPort := new(bytes.Buffer)
				ipPort.WriteString(nettcp.Ip)
				ipPort.WriteString(":")
				ipPort.WriteString(string(nettcp.Port))

				tcp4Str := new(bytes.Buffer)
				tcp4Str.WriteString(ipPort.String())
				tcp4Str.WriteString(nettcp.Exe)
				tcp4 = append(tcp4, tcp4Str.String())
			}
		}

		d.TCP4Listen = tcp4

		gonetstat6 := GOnetstat.Tcp6()
		tcp6 := []string{}
		for _, nettcp := range gonetstat6 {
			if nettcp.State == "LISTEN" {
				ipPort := new(bytes.Buffer)
				ipPort.WriteString(nettcp.Ip)
				ipPort.WriteString(":")
				ipPort.WriteString(string(nettcp.Port))

				tcp6Str := new(bytes.Buffer)
				tcp6Str.WriteString(ipPort.String())
				tcp6Str.WriteString(string(nettcp.Exe))
				tcp6 = append(tcp6, tcp6Str.String())
			}
		}

		d.TCP6Listen = tcp6
	}
}
