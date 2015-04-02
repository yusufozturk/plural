package main

import (
     "fmt"
     "net/http"
     "github.com/shirou/gopsutil/mem"
     "github.com/shirou/gopsutil/disk"
     "github.com/shirou/gopsutil/cpu"
     "github.com/shirou/gopsutil/host"
     "github.com/shirou/gopsutil/load"
     "github.com/shirou/gopsutil/net"
)

func all(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8000", nil)
}
