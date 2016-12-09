package cloud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"plural/data"
	"strings"
)

func awsClient(route string) string {
	url := "http://169.254.169.254/latest/" + route
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		return string(body)
	}
	return "error"
}

func Aws() {

	m := data.PluralJSON

	awsResponse, _ := http.Get("http://169.254.169.254/latest/")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if awsResponse != nil && awsResponse.Status == string("200 OK") {
		m["Ec2AmiID"] = awsClient("ami-id")
		m["Ec2InstanceID"] = awsClient("instance-id")
		m["Ec2InstanceType"] = awsClient("instance-type")
		m["Ec2AvailabilityZone"] = awsClient("placement/availability-zone")
		m["Ec2Profile"] = awsClient("profile")
		m["Ec2PublicIP4"] = awsClient("public-ipv4")
		securityGroupSplit := strings.Split(awsClient("security-groups"), ",")
		m["Ec2SecurityGroups"] = securityGroupSplit

	}

}
