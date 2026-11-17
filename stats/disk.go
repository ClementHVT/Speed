package stats

import (
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskStats struct {
	UsedPercent float64
	Total       uint64
	Used        uint64
}


func GetDiskStats() (DiskStats, error) {
	diskStats := DiskStats{}

	// Disk (cross-platform main disk detection)
	partitions, err := disk.Partitions(false)
	if err != nil {
		return diskStats, err
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

	return diskStats, nil
}