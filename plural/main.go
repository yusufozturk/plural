// (C) Copyright 2015 Timothy Marcinowski
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//  http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

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
  "strconv"
  "encoding/json"
  "math"
  "crypto/sha256"
  "flag"
  "plural/networkip"

  "github.com/spf13/viper"
  "github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/disk"
  "github.com/shirou/gopsutil/host"
  "github.com/shirou/gopsutil/load"
  "github.com/shirou/gopsutil/cpu"
  "github.com/dustin/go-humanize"
  "github.com/fsouza/go-dockerclient"
  "github.com/marshyski/GOnetstat"
)

// HTTP client timeout
var timeout = time.Duration(300 * time.Millisecond)

func dialTimeout(network, addr string) (net.Conn, error) {
  return net.DialTimeout(network, addr, timeout)
}

// Command-line flags
var configFlag = flag.String("config", "", "  Set configuration path, default is /opt/plural/conf")
var daemonFlag = flag.Bool("daemon", false, "  Run in daemon mode")
var outputFlag = flag.String("output", "", "  Output JSON file in a directory specified")

func init() {
  flag.StringVar(configFlag, "c", "", "  Set configuration path, default is /opt/plural/conf")
  flag.BoolVar(daemonFlag, "d", false, "  Run in daemon mode")
  flag.StringVar(outputFlag, "o", "", "  Output JSON file in a directory specified")
}

var usage = `Usage: plural [options] <args>

    -d, --daemon       Run in daemon mode
    -c, --config     Set configuration path, default path is /opt/plural/conf
    -o, --output     Output JSON file in a directory specified


Example:       plural -d -c /opt/plural/conf -o /tmp

Documentation:  https://github.com/marshyski/plural/blob/master/README.md

`

func main() {

  flag.Usage = func() {
    fmt.Println(usage)
  }

  flag.Parse()

  for {

    // Docker client remote API
    endpoint := "unix:///var/run/docker.sock"
    dockerClient, _ := docker.NewClient(endpoint)

    // Configuration file settings using key-value
    viper.SetConfigName("plural")
    if *configFlag != "" {
      viper.AddConfigPath(*configFlag)
      } else {
        viper.AddConfigPath("/opt/plural/conf")
      }

      viperReadConfig := viper.ReadInConfig()

      // Default settings if no config file is supplied
      viper.SetDefault("elastic_host", "localhost")
      viper.SetDefault("elastic_port", "9200")
      viper.SetDefault("environment", "dev")
      viper.SetDefault("interval", "300")
      viper.SetDefault("username", "undef")
      viper.SetDefault("password", "undef")
      viper.SetDefault("overwrite", "undef")

      elastic_host := viper.GetString("elastic_host")
      elastic_port := viper.GetString("elastic_port")
      environment := viper.GetString("environment")
      interval := viper.GetInt("interval")
      username := viper.GetString("username")
      password := viper.GetString("password")
      overwrite := viper.GetString("overwrite")

      transport := http.Transport{
        Dial: dialTimeout,
      }

      client := http.Client{
        Transport: &transport,
      }

      cpuCounts, _ := cpu.CPUCounts(true)
      v, _ := mem.VirtualMemory()
      k, _ := disk.DiskUsage("/")
      h, _ := host.HostInfo()
      l, _ := load.LoadAvg()
      memusedprctConv := strconv.FormatFloat(v.UsedPercent, 'f', 6, 64)
      memusedprct := strings.Split(memusedprctConv, ".")[0]
      memfree := humanize.Bytes(v.Free)
      memtotal := humanize.Bytes(v.Total)
      diskusedprctConv := strconv.FormatFloat(k.UsedPercent, 'f', 6, 64)
      diskusedprct := strings.Split(diskusedprctConv, ".")[0]
      diskfree := humanize.Bytes(k.Free)
      disktotal := humanize.Bytes(k.Total)
      timeN := time.Now()
      dateStamp := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", timeN.Year(), timeN.Month(), timeN.Day(), timeN.Hour(), timeN.Minute(), timeN.Second())
      loadoneConv := strconv.FormatFloat(l.Load1, 'f', 6, 64)
      loadone := strings.Split(loadoneConv, ".")[0]
      loadfifteenConv := strconv.FormatFloat(l.Load15, 'f', 6, 64)
      loadfifteen := strings.Split(loadfifteenConv, ".")[0]
      loadfiveConv := strconv.FormatFloat(l.Load5, 'f', 6, 64)
      loadfive := strings.Split(loadfiveConv, ".")[0]

      ipaddress, err := networkip.ExternalIP()
      if err != nil {
        fmt.Println(err.Error())
      }

      // UNIX system commands
      dnsfile := exec.Command("ls", "/etc/resolv.conf")
      dnsfileout, err := dnsfile.Output()
      dnsgrep := exec.Command("grep", "nameserver", "/etc/resolv.conf")
      dnsawk := exec.Command("awk", "{print$2}")
      dnsgrepOut, err := dnsgrep.StdoutPipe()
      dnsgrep.Start()
      dnsawk.Stdin = dnsgrepOut
      dnsOut, err := dnsawk.Output()
      dnsstring := string(dnsOut)
      dnsoutputSlice := strings.Split(dnsstring,"\n")
      dnsjs,_ := json.Marshal(dnsoutputSlice)

      wbin := exec.Command("ls", "/usr/bin/w")
      wbinout, err := wbin.Output()
      wh := exec.Command("w", "-h")
      whawk := exec.Command("awk", "{print$1\"-\"$2}")
      whOut, err := wh.StdoutPipe()
      wh.Start()
      whawk.Stdin = whOut
      wOut, err := whawk.Output()
      whstring := string(wOut)
      whoutputSlice := strings.Split(whstring,"\n")
      whjs,_ := json.Marshal(whoutputSlice)

      passfile := exec.Command("ls", "/etc/passwd")
      passfileout, err := passfile.Output()
      passgrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
      passgrepOut, err := passgrep.Output()
      passstring := string(passgrepOut)
      passoutputSlice := strings.Split(passstring,"\n")
      passjs,_ := json.Marshal(passoutputSlice)

      auditctlbin := exec.Command("ls", "/sbin/auditctl")
      auditctlbinout, err := auditctlbin.Output()
      auditctl := exec.Command("auditctl", "-l")
      auditctlOut, err := auditctl.Output()
      auditctlstring := string(auditctlOut)
      auditctloutputSlice := strings.Split(auditctlstring,"\n")
      auditctljs,_ := json.Marshal(auditctloutputSlice)

      rpmbin := exec.Command("ls", "/bin/rpm")
      rpmbinout, err := rpmbin.Output()
      rpmqa := exec.Command("rpm", "-qa")
      rpmsort := exec.Command("sort")
      rpmqaOut, err := rpmqa.StdoutPipe()
      rpmqa.Start()
      rpmsort.Stdin = rpmqaOut
      rpmOut, err := rpmsort.Output()
      rpmstring := string(rpmOut)
      rpmoutputSlice := strings.Split(rpmstring,"\n")
      rpmjs,_ := json.Marshal(rpmoutputSlice)

      dpkgbin := exec.Command("ls", "/usr/bin/dpkg")
      dpkgbinout, err := dpkgbin.Output()
      dpkg := exec.Command("dpkg", "-l")
      dpkgawk := exec.Command("awk", "/^[a-z]/{print$2\"-\"$3}")
      dpkglOut, err := dpkg.StdoutPipe()
      dpkg.Start()
      dpkgawk.Stdin = dpkglOut
      dpkgOut, err := dpkgawk.Output()
      dpkgstring := string(dpkgOut)
      dpkgoutputSlice := strings.Split(dpkgstring,"\n")
      dpkgjs,_ := json.Marshal(dpkgoutputSlice)

      iptablesbin := exec.Command("ls", "/sbin/iptables")
      iptablesbinout, err := iptablesbin.Output()
      iptablesl := exec.Command("iptables", "-L")
      iptablesgrep := exec.Command("grep", "-v", "^Chain\\|target\\|^$")
      iptableslOut, err := iptablesl.StdoutPipe()
      iptablesl.Start()
      iptablesgrep.Stdin = iptableslOut
      iptablesOut, err := iptablesgrep.Output()
      iptablesstring := string(iptablesOut)
      iptablesoutputSlice := strings.Split(iptablesstring,"\n")
      iptablesjs,_ := json.Marshal(iptablesoutputSlice)

      dockerbin := exec.Command("which", "docker")
      dockerbinout, err := dockerbin.Output()

      pipbin := exec.Command("which", "pip")
      pipbinout, err := pipbin.Output()
      pipfree := exec.Command("pip", "freeze")
      pipsort := exec.Command("sort")
      pipfreeOut, err := pipfree.StdoutPipe()
      pipfree.Start()
      pipsort.Stdin = pipfreeOut
      pipOut, err := pipsort.Output()
      pipstring := string(pipOut)
      pipreplace := strings.Replace(pipstring, "==", "-", -1)
      pipoutSlice := strings.Split(pipreplace,"\n")
      pipjs,_ := json.Marshal(pipoutSlice)

      gembin := exec.Command("which", "gem")
      gembinout, err := gembin.Output()
      gemlist := exec.Command("gem", "list")
      gemgrep := exec.Command("grep", "^[a-zA-Z]")
      gemlistOut, err := gemlist.StdoutPipe()
      gemlist.Start()
      gemgrep.Stdin = gemlistOut
      gemOut, err := gemgrep.Output()
      gemstring := string(gemOut)
      gemreplace := strings.Replace(gemstring, " (", "-", -1)
      gemreplace2 := strings.Replace(gemreplace, ")", "", -1)
    gemoutSlice := strings.Split(gemreplace2,"\n")
    gemjs,_ := json.Marshal(gemoutSlice)

    iproutebin := exec.Command("ls", "/sbin/ip")
    iproutebinout, err := iproutebin.Output()
    iproute := exec.Command("ip", "route")
    iprouteOut, err := iproute.Output()
    iproutestring := string(iprouteOut)
    iprouteoutputSlice := strings.Split(iproutestring,"\n")
    iproutejs,_ := json.Marshal(iprouteoutputSlice)

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
    }

    // ElasticSearch endpoint
    elastic_url := "http://" + elastic_host + ":" + elastic_port + "/" + environment + "/" + h.Hostname

    // JSON file name
    if *outputFlag != "" {
      if string(*outputFlag)[:len(string(*outputFlag))-1] != "/" {
        *outputFlag += "/"
      }
    }

    filename := *outputFlag + h.Hostname + ".json"

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

    if string(auditctlbinout) != "" {
      audit_rules :=`
      "audit_rules": %s,`

      auditctlLine := fmt.Sprintf(audit_rules, string(auditctljs))
      auditctlReplace := strings.Replace(auditctlLine, ",\"\"]", "]", -1)
      writeAuditctl, err := io.WriteString(f, auditctlReplace)
      if err != nil {
        fmt.Println(writeAuditctl, err)
        return
      }
    }

    top := `
    "cpu_count": "%v",
    "diskfree": "%v",
    "disktotal": "%v",
    "diskused": "%v",`

    topLine := fmt.Sprintf(top, cpuCounts, diskfree, disktotal, diskusedprct)
    writeTop, err := io.WriteString(f, topLine)
    if err != nil {
      fmt.Println(writeTop, err)
      return
    }

    if string(dnsfileout) != "" {
      dns_nameserver :=`
      "dns_nameserver": %s,`

      dnsLine := fmt.Sprintf(dns_nameserver, string(dnsjs))
      dnsReplace := strings.Replace(dnsLine, ",\"\"]", "]", -1)
      writeDns, err := io.WriteString(f, dnsReplace)
      if err != nil {
        fmt.Println(writeDns, err)
        return
      }
    }

    if string(dockerbinout) != "" {
      dockerRaw :=`
      "docker": %s,`
      containers, _ := dockerClient.ListContainers(docker.ListContainersOptions{All: false})
      dockerString := `[`

      for _, container := range containers {
        portsRaw := `%v`
        portsString := fmt.Sprintf(portsRaw, container.Ports)
        portsReplace := strings.Replace(portsString, "{", "", -1)
        portsReplace2 := strings.Replace(portsReplace, "}", "", -1)
        portsReplace3 := strings.Replace(portsReplace2, "[", "'", -1)
        portsReplace4 := strings.Replace(portsReplace3, "]", "'", -1)
        containerString :=`"%v, %v, %v",`
        dockerString += fmt.Sprintf(containerString, container.Image, container.Command, portsReplace4)
      }
      dockerString += `""]`

      dockerLine := fmt.Sprintf(dockerRaw, dockerString)
      dockerReplace := strings.Replace(dockerLine, ",\"\"]", "]", -1)
      writeDocker, err := io.WriteString(f, dockerReplace)
      if err != nil {
        fmt.Println(writeDocker, err)
        return
      }
    }

    domain := `
    "domain": "%s",`

    domainLine := fmt.Sprintf(domain, strings.TrimSpace(domainstring))
    writeDomain, err := io.WriteString(f, domainLine)
    if err != nil {
      fmt.Println(writeDomain, err)
      return
    }

    // Local AWS meta-data
    awsResponse, err := client.Get("http://169.254.169.254/latest/")
    if awsResponse != nil && awsResponse.Status == string("200 OK") {
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
      publicip, err := http.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
      defer publicip.Body.Close()
      publicipOut, err := ioutil.ReadAll(publicip.Body)
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
      "ec2_public_ip4": "%s",
      "ec2_security_groups": "%s",`

      awsLine := fmt.Sprintf(aws, amiidOut, availabilityzoneOut, instanceidOut, instancetypeOut, profileOut, publicipOut, securitygroupsOut)
      writeAWS, err := io.WriteString(f, awsLine)
      if err != nil {
        fmt.Println(writeAWS, err)
        return
      }
    }

    environmentOut :=`
    "environment": "%s",`

    envLine := fmt.Sprintf(environmentOut, environment)
    writeEnv, err := io.WriteString(f, envLine)
    if err != nil {
      fmt.Println(writeEnv, err)
      return
    }

    if string(gembinout) != "" {
      gem :=`
      "gem": %s,`

      gemLine := fmt.Sprintf(gem, string(gemjs))
      gemReplace := strings.Replace(gemLine, ",\"\"]", "]", -1)
      writeGem, err := io.WriteString(f, gemReplace)
      if err != nil {
        fmt.Println(writeGem, err)
        return
      }
    }

    if string(iptablesbinout) != "" {
      iptables :=`
      "iptables": %s,`

      iptablesLine := fmt.Sprintf(iptables, string(iptablesjs))
      iptablesReplace := strings.Replace(iptablesLine, ",\"\"]", "]", -1)
      writeIptables, err := io.WriteString(f, iptablesReplace)
      if err != nil {
        fmt.Println(writeIptables, err)
        return
      }
    }

    if string(iproutebinout) != "" {
      ip_route :=`
      "ip_route": %s,`

      iprouteLine := fmt.Sprintf(ip_route, string(iproutejs))
      iprouteReplace := strings.Replace(iprouteLine, ",\"\"]", "]", -1)
      writeIproute, err := io.WriteString(f, iprouteReplace)
      if err != nil {
        fmt.Println(writeIproute, err)
        return
      }
    }

    bottom := `
    "hostname": "%s",
    "ipaddress": "%s",
    "kernelversion": "%s",
    "lastrun": "%s",
    "load15": "%v",
    "load1": "%v",
    "load5": "%v",
    "memoryfree": "%v",
    "memorytotal": "%v",
    "memoryused": "%v",
    "os": "%v",`


    bottomLine := fmt.Sprintf(bottom, h.Hostname, ipaddress, strings.TrimSpace(kernelverstring), dateStamp, loadfifteen, loadone, loadfive, memfree, memtotal, memusedprct, h.OS)
    writeBottom, err := io.WriteString(f, bottomLine)
    if err != nil {
      fmt.Println(writeBottom, err)
      return
    }


    if string(rpmbinout) != "" {
      packages :=`
      "packages": %s,`

      rpmLine := fmt.Sprintf(packages, string(rpmjs))
      rpmReplace := strings.Replace(rpmLine, ",\"\"]", "]", -1)
      writeRpm, err := io.WriteString(f, rpmReplace)
      if err != nil {
        fmt.Println(writeRpm, err)
        return
      }
    }

    if string(dpkgbinout) != "" {
      packages :=`
      "packages": %s,`

      dpkgLine := fmt.Sprintf(packages, string(dpkgjs))
      writeDpkg, err := io.WriteString(f, dpkgLine)
      if err != nil {
        fmt.Println(writeDpkg, err)
        return
      }
    }

    if string(pipbinout) != "" {
      pip :=`
      "pip": %s,`

      pipLine := fmt.Sprintf(pip, string(pipjs))
      pipReplace := strings.Replace(pipLine, ",\"\"]", "]", -1)
      writePip, err := io.WriteString(f, pipReplace)
      if err != nil {
        fmt.Println(writePip, err)
        return
      }
    }

    gonetstat4, err := GOnetstat.Tcp(false)
    tcp4String := `[`
    for _, nettcp := range(gonetstat4) {
      if nettcp.State == "LISTEN" {
        ip_port := fmt.Sprintf("%v:%v", nettcp.Ip, nettcp.Port)
        pid_program := fmt.Sprintf("%v", nettcp.Exe)
        ippidString :=`"%v %v",`
        tcp4String += fmt.Sprintf(ippidString, ip_port, pid_program)
      }
    }
    tcp4String += `""]`
    tcp4Replace := strings.Replace(tcp4String, ",\"\"]", "]", -1)

    gonetstat6, err := GOnetstat.Tcp6(false)
    tcp6String := `[`
    for _, nettcp := range(gonetstat6) {
      if nettcp.State == "LISTEN" {
        ip_port := fmt.Sprintf("%v:%v", nettcp.Ip, nettcp.Port)
        pid_program := fmt.Sprintf("%v", nettcp.Exe)
        ippidString :=`"%v %v",`
        tcp6String += fmt.Sprintf(ippidString, ip_port, pid_program)
      }
    }
    tcp6String += `""]`
    tcp6Replace := strings.Replace(tcp6String, ",\"\"]", "]", -1)

    if err != nil {
       fmt.Println(err)
       return
    }

    beforeLast := `
    "platform": "%v",
    "platformfamily": "%v",
    "platformverison": "%v",
    "tcp4_listen": %v,
    "tcp6_listen": %v,`

    beforeLastLine := fmt.Sprintf(beforeLast, h.Platform, h.PlatformFamily, h. PlatformVersion, tcp4Replace, tcp6Replace)
    writeBeforeLast, err := io.WriteString(f, beforeLastLine)
    if err != nil {
      fmt.Println(writeBeforeLast, err)
      return
    }

    timezoneLast := `
    "timezone": "%s",
    "uptime": "%v",`

    timezoneLastLine := fmt.Sprintf(timezoneLast, strings.TrimSpace(timezonestring), time.Duration(h.Uptime) * time.Second)
    writeTimezoneLast, err := io.WriteString(f, timezoneLastLine)
    if err != nil {
      fmt.Println(writeTimezoneLast, err)
      return
    }

    if string(passfileout) != "" {
      users :=`
      "users": %s,`

      usersLine := fmt.Sprintf(users, string(passjs))
      usersReplace := strings.Replace(usersLine, ",\"\"]", "]", -1)
      writeUsers, err := io.WriteString(f, usersReplace)
      if err != nil {
        fmt.Println(writeUsers, err)
        return
      }
    }

    if string(wbinout) != "" {
      users_loggedin :=`
      "users_loggedin": %s,`

      whLine := fmt.Sprintf(users_loggedin, string(whjs))
      whReplace := strings.Replace(whLine, ",\"\"]", "]", -1)
      writeWh, err := io.WriteString(f, whReplace)
      if err != nil {
        fmt.Println(writeWh, err)
        return
      }
    }

    last := `
    "virtualizationrole": "%v",
    "virtualizationsystem": "%v"
    }`

    lastLine := fmt.Sprintf(last, h.VirtualizationRole, h.VirtualizationSystem)
    writeLast, err := io.WriteString(f, lastLine)
    if err != nil {
      fmt.Println(writeLast, err)
      return
    }

    f.Close()

    // Generate SHA256 SUM of JSON file
    const filechunk = 8192

    file, err := os.Open(filename)

    if err != nil {
      fmt.Println(err.Error())
      return
    }

    defer file.Close()

    info, _ := file.Stat()
    filesize := info.Size()
    blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
    hash := sha256.New()

    for i := uint64(0); i < blocks; i++ {
      blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
      buf := make([] byte, blocksize)

      file.Read(buf)
      io.WriteString(hash, string(buf))
    }

    if viperReadConfig != nil {
      fmt.Println(dateStamp, h.Hostname, "INFO no config file used, using default configuration")
    }

    fmt.Printf("%s %s INFO %s SHA256 checksum is %x\n",dateStamp, h.Hostname, file.Name(), hash.Sum(nil))

    // Check to see if ElasticSearch server is up
    elasticResponse, err := client.Get(elastic_url)
    if elasticResponse != nil {
      jsonStr,err := ioutil.ReadFile(filename)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      fmt.Println(dateStamp, h.Hostname, "INFO elasticsearch endpoint:", elastic_url)
      if overwrite == "enable" {
        reqDelete, err := http.NewRequest("DELETE", elastic_url, nil)
        if username != "undef" {
          reqDelete.SetBasicAuth(username, password)
        }
        respDelete, err := http.DefaultClient.Do(reqDelete)
        fmt.Println(dateStamp, h.Hostname, "DELETE elasticsearch type status:", respDelete.Status)
        if err != nil {
          fmt.Println(err.Error())
        }
      }
      reqPost, err := http.NewRequest("POST", elastic_url, bytes.NewBuffer(jsonStr))
      if password != "undef" {
        reqPost.SetBasicAuth(username, password)
      }
      reqPost.Header.Set("Content-Type", "application/json")

      clientReq := &http.Client{}
      respPost, err := clientReq.Do(reqPost)
      if err != nil {
        fmt.Println(err.Error())
      }
      defer respPost.Body.Close()
      fmt.Println(dateStamp, h.Hostname, "POST json elasticsearch type status:", respPost.Status)
      postBody, _ := ioutil.ReadAll(respPost.Body)
      fmt.Println(dateStamp, h.Hostname, "POST response body:", string(postBody))
      } else {
        fmt.Println(dateStamp, h.Hostname, "FAIL unable to connect to elasticeearch server:", "http://" + elastic_host + ":" + elastic_port)
      }

      if !*daemonFlag {
        break
      }

      // Sleep / interval time for daemon
      time.Sleep(time.Duration(interval) * time.Second)

    }
  }
