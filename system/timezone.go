package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/marshyski/plural/data"
)

func TimeZone(d *data.PluralJSON) {
	if runtime.GOOS != "windows" {
		timezone := exec.Command("date", "+%Z")
		timezoneOut, err := timezone.Output()
		if err != nil {
			return
		}
		timezoneStr := string(timezoneOut)
		timezoneTrim := strings.TrimSpace(timezoneStr)

		d.Timezone = timezoneTrim
	}
}
