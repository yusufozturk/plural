package main

import (
     "fmt"
     "github.com/shirou/gopsutil/mem"
     "github.com/shirou/gopsutil/disk"
//     "github.com/shirou/gopsutil/cpu"
     "github.com/shirou/gopsutil/host"
)

func main() {
     v, _ := mem.VirtualMemory()
     k, _ := disk.DiskUsage("/")
//     c, _ := cpu.CPUInfo()
     h, _ := host.HostInfo()
     // almost every return value is a struct
     fmt.Printf("mem_total: %v\nmem_free: %v\nmem_used: %f%%\n", v.Total, v.Free, v.UsedPercent)
     fmt.Printf("disk_total: %v\ndisk_free: %v\ndisk_used: %f%%\n", k.Total, k.Free, k.UsedPercent)
//     fmt.Printf("cpu: %v\ncpu_vendor: %v\ncpu_family: %v\ncpu_cores: %v\ncpu_model: %f%%\n", c.CPU, c.VendorID, c.Family, c.Cores, c.Model)
     fmt.Printf("hostname: %v\nuptime: %v\nos: %v\nplatform: %v\nplatform_family: %v\nplatform_ver: %v\nvirt_sys: %v\nvirt_role: %v\n", h.Hostname, h.Uptime, h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion, h.VirtualizationSystem, h.VirtualizationRole)
}
