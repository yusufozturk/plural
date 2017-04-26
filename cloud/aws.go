package cloud

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/marshyski/plural/data"
)

func awsClient(route string) string {
	url := "http://169.254.169.254/latest/" + route
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body)
}

func Aws() {

	m := data.PluralJSON

	awsResponse, _ := http.Get("http://169.254.169.254/latest/")
	if awsResponse != nil && awsResponse.Status == string("200 OK") {
		m["Ec2AmiID"] = awsClient("ami-id")
		m["Ec2InstanceID"] = awsClient("instance-id")
		m["Ec2InstanceType"] = awsClient("instance-type")
		m["Ec2AvailabilityZone"] = awsClient("placement/availability-zone")
		m["Ec2Profile"] = awsClient("profile")
		m["Ec2PublicIP4"] = awsClient("public-ipv4")
		securityGroupSplit := strings.Split(strings.TrimSpace(awsClient("security-groups")), "\n")
		m["Ec2SecurityGroups"] = securityGroupSplit

	}

}
