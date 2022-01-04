package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"strconv"
)
import "code.cloudfoundry.org/bytefmt"

type Disk struct {
	Name        string `json:"name,omitempty"`
	Label       string `json:"label,omitempty"`
	Used        string `json:"used,omitempty"`
	Total       string `json:"total,omitempty"`
	Available   string `json:"available,omitempty"`
	UsedPercent float64    `json:"used_percent,omitempty"`
}

func GetDisks() []Disk {
	parts, _ := disk.Partitions(false)
	var disks = []Disk{}
	for _, part := range parts {
		label, _ := disk.Label(part.Device)
		usage, _ := disk.Usage(part.Mountpoint)
		percent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", usage.UsedPercent), 64)
		disks = append(disks, Disk{
			Name:        part.Device,
			Label:       label,
			Used:        bytefmt.ByteSize(usage.Used),
			Total:       bytefmt.ByteSize(usage.Total),
			Available:   bytefmt.ByteSize(usage.Total - usage.Used),
			UsedPercent: percent,
		})
	}

	return disks
}