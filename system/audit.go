package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func Audit(d *data.PluralJSON) {
	if runtime.GOOS == "linux" {
		audit := exec.Command("auditctl", "-l")
		auditOut, err := audit.Output()
		if err != nil {
			return
		}
		auditStr := string(auditOut)
		auditSlice := strings.Split(strings.TrimSpace(auditStr), "\n")

		if auditSlice != nil {
			d.AuditRules = auditSlice
		}
	}
}
