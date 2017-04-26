package system

import (
	"os/exec"
	"strings"

	"github.com/marshyski/plural/data"
)

func TimeZone() {
	m := data.PluralJSON

	timezone := exec.Command("date", "+%Z")
	timezoneOut, _ := timezone.Output()
	timezoneStr := string(timezoneOut)
	timezoneTrim := strings.TrimSpace(timezoneStr)

	m["Timezone"] = timezoneTrim
}
