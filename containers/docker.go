package containers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/marshyski/plural/data"

	"github.com/fsouza/go-dockerclient"
)

func Docker(d *data.PluralJSON) {

	endpoint := "unix:///var/run/docker.sock"
	dockerClient, err := docker.NewClient(endpoint)
	if err != nil {
		return
	}

	containers, _ := dockerClient.ListContainers(docker.ListContainersOptions{All: true})
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
	d.Docker = ctrSlice
}
