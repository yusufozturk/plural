package system

import (
	"github.com/marshyski/plural/config"
	"github.com/marshyski/plural/data"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	gb   = 1073741824
	days = 86400
)

func Stats(d *data.PluralJSON) {

	cpuinfo, _ := cpu.Info()
	cpuCount := cpuinfo[0].Cores
	v, _ := mem.VirtualMemory()
	k, _ := disk.Usage("/")
	h, _ := host.Info()
	l, _ := load.Avg()
	memUsed := v.Used / gb
	memFree := v.Free / gb
	memTotal := v.Total / gb
	diskUsed := k.Used / gb
	diskFree := k.Free / gb
	diskTotal := k.Total / gb

	d.CPUCount = cpuCount
	d.Memoryused = memUsed
	d.Memoryfree = memFree
	d.Memorytotal = memTotal
	d.Diskused = diskUsed
	d.Diskfree = diskFree
	d.Disktotal = diskTotal
	d.Load1 = l.Load1
	d.Load5 = l.Load5
	d.Load15 = l.Load15
	d.Hostname = h.Hostname
	d.Platform = h.Platform
	d.Platformfamily = h.PlatformFamily
	d.Platformverison = h.PlatformVersion
	d.Kernelversion = h.KernelVersion
	d.Virtualizationrole = h.VirtualizationRole
	d.Os = h.OS
	d.Uptime = h.Uptime / days
	d.Environment = config.ConfigStr("environment")
}
