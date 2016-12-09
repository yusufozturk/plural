package system

import (
	"plural/config"
	"plural/data"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func Stats() {

	m := data.PluralJSON

	cpuCount, _ := cpu.Counts(true)
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
	loadoneConv := strconv.FormatFloat(l.Load1, 'f', 6, 64)
	load1 := strings.Split(loadoneConv, ".")[0]
	loadfifteenConv := strconv.FormatFloat(l.Load15, 'f', 6, 64)
	load15 := strings.Split(loadfifteenConv, ".")[0]
	loadfiveConv := strconv.FormatFloat(l.Load5, 'f', 6, 64)
	load5 := strings.Split(loadfiveConv, ".")[0]

	m["CPUCount"] = string(cpuCount)
	m["Memoryused"] = memUsed
	m["Memoryfree"] = memFree
	m["Memorytotal"] = memTotal
	m["Diskused"] = diskUsed
	m["Diskfree"] = diskFree
	m["Disktotal"] = diskTotal
	m["Load1"] = load1
	m["Load5"] = load5
	m["Load15"] = load15
	m["Hostname"] = h.Hostname
	m["Platform"] = h.Platform
	m["Platformfamily"] = h.PlatformFamily
	m["Platformverison"] = h.PlatformVersion
	m["Kernelversion"] = h.KernelVersion
	m["Os"] = h.OS
	m["Uptime"] = string(time.Duration(h.Uptime) * time.Second)
	m["Environment"] = config.ConfigStr("environment")
}
