package stats

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// GetStats returns CPU and memory usage percentages
func GetStats() (float64, float64, uint64, uint64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	if len(percent) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("no CPU data returned")
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return percent[0], v.UsedPercent, v.Total, v.Used, nil
}
