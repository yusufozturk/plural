package main

import (
    "fmt"
    "net/http"
    "net"
    "time"
    "os"
    "io"
    "os/exec"
    "io/ioutil"
    "strings"
    "bytes"
    "github.com/spf13/viper"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/load"
)

// HTTP client timeout
var timeout = time.Duration(300 * time.Millisecond)

func dialTimeout(network, addr string) (net.Conn, error) {
        return net.DialTimeout(network, addr, timeout)
}

func main() {

  for {

    // Configuration file settings using key-value
    viper.SetConfigName("plural")
    viper.AddConfigPath("/opt/plural/conf")
    err := viper.ReadInConfig()
    if err != nil {
       fmt.Println("No Configuration File Using DEFAULTS")
    }

    // Default settings if no config file is supplied
    viper.SetDefault("elastic_host", "localhost")
    viper.SetDefault("elastic_port", "9200")
    viper.SetDefault("environment", "dev")

    elastic_host := viper.GetString("elastic_host")
    elastic_port := viper.GetString("elastic_port")
    environment := viper.GetString("environment")

    transport := http.Transport{
       Dial: dialTimeout,
    }

    client := http.Client{
       Transport: &transport,
    }

    // UNIX system commands
    kernelver := exec.Command("uname","-r")
    kernelverout, err := kernelver.Output()
    kernelverstring := string(kernelverout)
    timezone := exec.Command("date","+%Z")
    timezoneout, err := timezone.Output()
    timezonestring := string(timezoneout)

    hostname := exec.Command("hostname","-f")
    hostcut := exec.Command("cut","-d.","-f","2-")
    hostnameOut, err := hostname.StdoutPipe()
    hostname.Start()
    hostcut.Stdin = hostnameOut
    domainname, err := hostcut.Output()
    domainstring := string(domainname)

    if err != nil {
       fmt.Println(err.Error())
       return
    }

    v, _ := mem.VirtualMemory()
    k, _ := disk.DiskUsage("/")
    h, _ := host.HostInfo()
    l, _ := load.LoadAvg()

    // ElasticSearch endpoint
    elastic_url := "http://" + elastic_host + ":" + elastic_port + "/" + environment + "/" + h.Hostname

    // JSON file name
    filename := h.Hostname + ".json"

    // Create JSON file
    f, err := os.Create(filename)
    if err != nil {
       fmt.Println(err.Error())
       return
    }
    n, err := io.WriteString(f, "{")
    if err != nil {
       fmt.Println(n, err)
       return
    }

    top := `
    "diskfree": "%v",
    "disktotal": "%v",
    "diskused": "%v",
    "domain": "%s",`

    topLine := fmt.Sprintf(top, k.Free, k.Total, k.UsedPercent, strings.TrimSpace(domainstring))
    writeTop, err := io.WriteString(f, topLine)
    if err != nil {
       fmt.Println(writeTop, err)
       return
    }

    // Local meta-data AWS
    awsResponse, err := client.Get("http://169.254.169.254/latest")
    if awsResponse != nil {
       amiid, err := http.Get("http://169.254.169.254/latest/meta-data/ami-id")
       defer amiid.Body.Close()
       amiidOut, err := ioutil.ReadAll(amiid.Body)
       instanceid, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
       defer instanceid.Body.Close()
       instanceidOut, err := ioutil.ReadAll(instanceid.Body)
       instancetype, err := http.Get("http://169.254.169.254/latest/meta-data/instance-type")
       defer instancetype.Body.Close()
       instancetypeOut, err := ioutil.ReadAll(instancetype.Body)
       availabilityzone, err := http.Get("http://169.254.169.254/latest/meta-data/placement/availability-zone")
       defer availabilityzone.Body.Close()
       availabilityzoneOut, err := ioutil.ReadAll(availabilityzone.Body)
       securitygroups, err := http.Get("http://169.254.169.254/latest/meta-data/security-groups")
       defer securitygroups.Body.Close()
       securitygroupsOut, err := ioutil.ReadAll(securitygroups.Body)
       profile, err := http.Get("http://169.254.169.254/latest/meta-data/profile")
       defer profile.Body.Close()
       profileOut, err := ioutil.ReadAll(profile.Body)
       if err != nil {
          fmt.Println(err.Error())
          return
       }

       aws := `
    "ec2_ami_id": "%s",
    "ec2_availability_zone": "%s",
    "ec2_instance_id": "%s",
    "ec2_instance_type": "%s",
    "ec2_profile": "%s",
    "ec2_security_groups": "%s",`

       awsLine := fmt.Sprintf(aws, amiidOut, availabilityzoneOut, instanceidOut, instancetypeOut, profileOut, securitygroupsOut)
       writeAWS, err := io.WriteString(f, awsLine)
       if err != nil {
          fmt.Println(writeAWS, err)
          return
       }
    }

    bottom := `
    "hostname": "%s",
    "kernelversion": "%s",
    "load15": "%v",
    "load1": "%v",
    "load5": "%v",
    "memoryfree": "%v",
    "memorytotal": "%v",
    "memoryused": "%v",
    "os": "%v",
    "platform": "%v",
    "platformfamily": "%v",
    "platformverison": "%v",
    "timezone": "%s",
    "uptime": "%v",
    "virtualizationrole": "%v",
    "virtualizationsystem": "%v"
  }`

    bottomLine := fmt.Sprintf(bottom, h.Hostname, strings.TrimSpace(kernelverstring), l.Load15, l.Load1, l.Load5, v.Free, v.Total, v.UsedPercent, h.OS, h.Platform, h.PlatformFamily, h. PlatformVersion, strings.TrimSpace(timezonestring), h.Uptime, h.VirtualizationRole, h.VirtualizationSystem)
    writeLine, err := io.WriteString(f, bottomLine)
    if err != nil {
       fmt.Println(writeLine, err)
       return
    }

    f.Close()

    // Check to see if ElasticSearch server is up
    elasticResponse, err := client.Get(elastic_url)
    if elasticResponse != nil {
       jsonStr,err := ioutil.ReadFile(filename)
       if err != nil {
          fmt.Println(err.Error())
          return
       }
       fmt.Println("ElasticSearch EndPoint:", elastic_url)
       reqDelete, err := http.NewRequest("DELETE", elastic_url, nil)
       respDelete, err := http.DefaultClient.Do(reqDelete)
       fmt.Println("Delete ElasticSearch Type Status:", respDelete.Status)
       reqPost, err := http.NewRequest("POST", elastic_url, bytes.NewBuffer(jsonStr))
       reqPost.Header.Set("Content-Type", "application/json")

       clientReq := &http.Client{}
       respPost, err := clientReq.Do(reqPost)
       if err != nil {
          fmt.Println(err.Error())
       }
       defer respPost.Body.Close()

       fmt.Println("POST JSON ElasticSearch Type Status:", respPost.Status)
       postBody, _ := ioutil.ReadAll(respPost.Body)
       fmt.Println("POST Response Body:", string(postBody))
    } else {
       fmt.Println("Unable to Connect to ElasticSearch Server:", "http://" + elastic_host + ":" + elastic_port)
    }

    // Sleep time for, for loop
    time.Sleep(5 * time.Minute)

  }
}
