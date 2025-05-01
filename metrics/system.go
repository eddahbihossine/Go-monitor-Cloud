package metrics

import (
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/disk"
)


type SystemStats struct {
	CPUPercent float64 `json:"cpu_percent"`
	MemoryUsed uint64  `json:"memory_used_mb"`
	MemoryTotal uint64 `json:"memory_total_mb"`
	DiskUsed   uint64  `json:"disk_used_mb"`
	DiskTotal  uint64  `json:"disk_total_mb"`
}

func GetSystemStats() SystemStats {
	cpuPercent, _ := cpu.Percent(0, false)
	vm, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")

	return SystemStats{
		CPUPercent:  cpuPercent[0],
		MemoryUsed:  vm.Used / (1024 * 1024),
		MemoryTotal: vm.Total / (1024 * 1024),
		DiskUsed:    diskStat.Used / (1024 * 1024),
		DiskTotal:   diskStat.Total / (1024 * 1024),
	}
}
