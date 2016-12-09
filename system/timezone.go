package system

import (
	"fmt"
	"os/exec"
	"plural/data"
	"strings"
)

func TimeZone() {
	m := data.PluralJSON

	timezone := exec.Command("date", "+%Z")
	timezoneOut, err := timezone.Output()
	if err != nil {
		fmt.Println(err)
	}
	timezoneStr := string(timezoneOut)
	timezoneTrim := strings.TrimSpace(timezoneStr)

	m["Timezone"] = timezoneTrim
}
