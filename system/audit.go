package system

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Audit() {
	m := data.PluralJSON

	auditBin := exec.Command("ls", "/sbin/auditctl")
	auditBinOut, _ := auditBin.Output()
	audit := exec.Command("auditctl", "-l")
	auditOut, _ := audit.Output()
	auditStr := string(auditOut)
	auditSlice := strings.Split(strings.TrimSpace(auditStr), "\n")

	if string(auditBinOut) != "" {
		m["AuditRules"] = auditSlice
	}
}
