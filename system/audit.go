package system

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func Audit() {
	m := data.PluralJSON

	auditBin := exec.Command("ls", "/sbin/auditctl")
	auditBinOut, err := auditBin.Output()
	if err != nil {
		fmt.Println(err)
	}
	audit := exec.Command("auditctl", "-l")
	auditOut, err := audit.Output()
	auditStr := string(auditOut)
	auditSlice := strings.Split(auditStr, "\n")

	if string(auditBinOut) != "" {
		m["AuditRules"] = auditSlice
	}
}
