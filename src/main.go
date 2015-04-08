package main

import (
    "log"
    "fmt"
    "net/http"
    "os/exec"
    "io/ioutil"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/net"
)

func all(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
       errorHandler(w, r, http.StatusNotFound)
       return
    }

    kernelver := exec.Command("uname","-r")
    kernelverout, err := kernelver.Output()
    timezone  := exec.Command("date","+%Z")
    timezoneout, err := timezone.Output()

    hostname := exec.Command("hostname","-f")
    hostcut := exec.Command("cud","-d.","-f","2-")
    hostnameOut, _ := hostname.StdoutPipe()
    hostname.Start()
    hostcut.Stdin = hostnameOut
    domainname, _ := hostcut.Output()

    if err != nil {
       fmt.Println(err.Error())
       return
    }

    v, _ := mem.VirtualMemory()
    k, _ := disk.DiskUsage("/")
    c, _ := cpu.CPUInfo()
    h, _ := host.HostInfo()
    l, _ := load.LoadAvg()
    n, _ := net.NetInterfaces()

   // AWS 

   response, err := http.Get("http://169.254.169.254/latest")
   if response.Status == string("200 OK") {
      if err != nil {
         fmt.Printf("%s", err)
       } else {
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
             fmt.Printf("%s", err)
          }
          fmt.Fprintf(w, "ec2_ami_id: %s\n", string(amiidOut))
          fmt.Fprintf(w, "ec2_instance_id: %s\n", string(instanceidOut))
          fmt.Fprintf(w, "ec2_instance_type: %s\n", string(instancetypeOut))
          fmt.Fprintf(w, "ec2_availability_zone: %s\n", string(availabilityzoneOut))
          fmt.Fprintf(w, "ec2_security_groups: %s\n", string(securitygroupsOut))
          fmt.Fprintf(w, "ec2_profile: %s\n", string(profileOut))
      }
    }

    fmt.Fprintf(w, "memorytotal: %v\nmemoryfree: %v\nmemoryused: %f%%\n", v.Total, v.Free, v.UsedPercent)
    fmt.Fprintf(w, "disktotal: %v\ndiskfree: %v\ndiskused: %f%%\n", k.Total, k.Free, k.UsedPercent)
    fmt.Fprintf(w, "cpu: %v\n", c)
    fmt.Fprintf(w, "hostname: %v\nuptime: %v\nos: %v\nplatform: %v\nplatformfamily: %v\nplatformverison: %v\nvirtualizationsystem: %v\nvirtualizationrole: %v\n", h.Hostname, h.Uptime, h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion, h.VirtualizationSystem, h.VirtualizationRole)
    fmt.Fprintf(w, "load1: %v\nload5: %v\nload15: %v\n", l.Load1, l.Load5, l.Load15)
    fmt.Fprintf(w, "networkinterfaces: %v\n", n)
    fmt.Fprintf(w, "kernelversion: %s", string(kernelverout))
    fmt.Fprintf(w, "timezone: %s", string(timezoneout))
    fmt.Fprintf(w, "domain: %s", domainname)
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
       handler.ServeHTTP(w, r)
    })
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
       fmt.Fprint(w, "404")
    }
}

func main() {
     http.HandleFunc("/", all)
     err := http.ListenAndServe(":8000", Log(http.DefaultServeMux))
     if err != nil {
        log.Fatal("ListenAndServe: ", err)
     }
}
