package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Audit(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		auditBin := exec.Command("ls", "/sbin/auditctl")
		auditBinOut, err := auditBin.Output()
		if err != nil {
			return
		}
		audit := exec.Command("auditctl", "-l")
		auditOut, err := audit.Output()
		if err != nil {
			return
		}
		auditStr := string(auditOut)
		auditSlice := strings.Split(strings.TrimSpace(auditStr), "\n")

		if string(auditBinOut) != "" {
			d.AuditRules = auditSlice
		}
	}
}
