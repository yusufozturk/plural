package system

import (
	"strconv"
	"strings"
	"time"

	"github.com/marshyski/plural/config"
	"github.com/marshyski/plural/data"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func Stats(d *data.PluralJSON) {

	cpuinfo, _ := cpu.Info()
	cpuCount := cpuinfo[0].Cores
	v, _ := mem.VirtualMemory()
	k, _ := disk.Usage("/")
	h, _ := host.Info()
	l, _ := load.Avg()
	memusedConv := strconv.FormatFloat(v.UsedPercent, 'f', 6, 64)
	memUsed := strings.Split(memusedConv, ".")[0]
	memFree := humanize.Bytes(v.Free)
	memTotal := humanize.Bytes(v.Total)
	diskusedConv := strconv.FormatFloat(k.UsedPercent, 'f', 6, 64)
	diskUsed := strings.Split(diskusedConv, ".")[0]
	diskFree := humanize.Bytes(k.Free)
	diskTotal := humanize.Bytes(k.Total)

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
	// Convert to days
	d.Uptime = int64(time.Duration(h.Uptime) / 86400)
	d.Environment = config.ConfigStr("environment")
}
