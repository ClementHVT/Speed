package stats

import (
	"github.com/shirou/gopsutil/v4/mem"
)

type MemStats struct {
	UsedPercent float64
	Total       uint64
	Used        uint64
}

func GetMemStats() (MemStats, error) {
	memStats := MemStats{}
	vm, err := mem.VirtualMemory()
	if err != nil {
		return memStats, err
	}

	memStats.UsedPercent = vm.UsedPercent
	memStats.Total = vm.Total
	memStats.Used = vm.Used

	return memStats, nil
}