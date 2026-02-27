package stats

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
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

func GetStats() (CPUStats, MemStats, error) {
	var cpuStats CPUStats
	var memStats MemStats

	// --- CPU Info (MHz) ---
	infos, err := cpu.Info()
	if err != nil {
		return cpuStats, memStats, err
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

	// --- CPU Usage ---
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return cpuStats, memStats, err
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

	// --- Memory ---
	vm, err := mem.VirtualMemory()
	if err != nil {
		return cpuStats, memStats, err
	}

	memStats.UsedPercent = vm.UsedPercent
	memStats.Total = vm.Total
	memStats.Used = vm.Used

	return cpuStats, memStats, nil
}
