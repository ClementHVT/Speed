package stats

func GetStats() (CPUStats, MemStats, DiskStats, error) {
	cpuStats, err := GetCPUStats()
	if err != nil {
		return CPUStats{}, MemStats{}, DiskStats{}, err
	}

	memStats, err := GetMemStats()
	if err != nil {
		return CPUStats{}, MemStats{}, DiskStats{}, err
	}

	diskStats, err := GetDiskStats()
	if err != nil {
		return CPUStats{}, MemStats{}, DiskStats{}, err
	}

	return cpuStats, memStats, diskStats, nil
}
