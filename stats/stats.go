package stats

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type CPUStats struct {
	PerCoreUsage []float64
	Usage        float64
	PerCoreMhz   []float64
	AvgMHz       float64
}

type MemStats struct {
	UsedPercent float64
	Total       uint64
	Used        uint64
}

type DiskStats struct {
	UsedPercent float64
	Total       uint64
	Used        uint64
}

func GetStats() (CPUStats, MemStats, DiskStats, error) {
	var cpuStats CPUStats
	var memStats MemStats
	var diskStats DiskStats

	// CPU Info
	infos, err := cpu.Info()
	if err != nil {
		return cpuStats, memStats, diskStats, err
	}

	coreCount := len(infos)
	if coreCount > 0 {
		cpuStats.PerCoreMhz = make([]float64, coreCount)

		var totalMHz float64
		for i, info := range infos {
			cpuStats.PerCoreMhz[i] = info.Mhz
			totalMHz += info.Mhz
		}

		cpuStats.AvgMHz = totalMHz / float64(coreCount)
	}

	// CPU Usage
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return cpuStats, memStats, diskStats, err
	}

	if len(percent) > 0 {
		cpuStats.PerCoreUsage = make([]float64, len(percent))

		var totalUsage float64
		for i, p := range percent {
			cpuStats.PerCoreUsage[i] = p
			totalUsage += p
		}

		cpuStats.Usage = totalUsage / float64(len(percent))
	}

	// Memory
	vm, err := mem.VirtualMemory()
	if err != nil {
		return cpuStats, memStats, diskStats, err
	}

	memStats.UsedPercent = vm.UsedPercent
	memStats.Total = vm.Total
	memStats.Used = vm.Used

	// Disk (cross-platform main disk detection)
	partitions, err := disk.Partitions(false)
	if err != nil {
		return cpuStats, memStats, diskStats, err
	}

	for _, p := range partitions {
		// Skip pseudo/virtual filesystems
		if p.Fstype == "" {
			continue
		}

		usage, err := disk.Usage(p.Mountpoint)
		if err == nil && usage.Total > 0 {
			diskStats.UsedPercent = usage.UsedPercent
			diskStats.Total = usage.Total
			diskStats.Used = usage.Used
			break
		}
	}

	return cpuStats, memStats, diskStats, nil
}
