package docker

import (
	"bytes"
	"fmt"
	"strings"

	"plural/data"
	"plural/utils"

	"github.com/fsouza/go-dockerclient"
)

func Containers() {

	if utils.SingleOut("which", "docker") != "" {

		endpoint := "unix:///var/run/docker.sock"
		dockerClient, _ := docker.NewClient(endpoint)

		containers, _ := dockerClient.ListContainers(docker.ListContainersOptions{All: false})
		ctrSlice := []string{}

		for _, container := range containers {
			portsRaw := `%v`
			portsString := fmt.Sprintf(portsRaw, container.Ports)
			portsReplace := strings.Replace(portsString, "{", "", -1)
			portsReplace2 := strings.Replace(portsReplace, "}", "", -1)
			portsReplace3 := strings.Replace(portsReplace2, "[", "'", -1)
			portsReplace4 := strings.Replace(portsReplace3, "]", "'", -1)
			ctrStr := new(bytes.Buffer)
			ctrStr.WriteString(container.Image)
			ctrStr.WriteString(", ")
			ctrStr.WriteString(container.Command)
			ctrStr.WriteString(", ")
			ctrStr.WriteString(portsReplace4)
			ctrSlice = append(ctrSlice, ctrStr.String())
		}
		m := data.PluralJSON
		m["Docker"] = ctrSlice
	}
}
