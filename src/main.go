package main

import (
    "log"
    "fmt"
    "net/http"
    "os/exec"
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


    fmt.Fprintf(w, "memorytotal: %v\nmemoryfree: %v\nmemoryused: %f%%\n", v.Total, v.Free, v.UsedPercent)
    fmt.Fprintf(w, "disktotal: %v\ndiskfree: %v\ndiskused: %f%%\n", k.Total, k.Free, k.UsedPercent)
    fmt.Fprintf(w, "cpu: %v\n", c)
    fmt.Fprintf(w, "hostname: %v\nuptime: %v\nos: %v\nplatform: %v\nplatformfamily: %v\nplatformverison: %v\nvirtualizationsystem: %v\nvirtualizationrole: %v\n", h.Hostname, h.Uptime, h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion, h.VirtualizationSystem, h.VirtualizationRole)
    fmt.Fprintf(w, "load1: %v\nload5: %v\nload15: %v\n", l.Load1, l.Load5, l.Load15)
    fmt.Fprintf(w, "networkinterfaces: %v\n", n)
    fmt.Fprintf(w, "kernelversion: %s\n", string(kernelverout))
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
