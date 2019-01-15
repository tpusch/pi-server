package main

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	Total uint64 `json:"all"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

type DiskString struct {
	Total string
	Used  string
	Free  string
}

func getDiskString(status DiskStatus) (out DiskString) {
	out.Total = fmt.Sprintf("%.2f GB", inGig(status.Total))
	out.Used = fmt.Sprintf("%.2f GB", inGig(status.Used))
	out.Free = fmt.Sprintf("%.2f GB", inGig(status.Free))
	return
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.Total = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.Total - disk.Free
	return
}

func addDisk(name string, status DiskStatus) string {
	format := `%s:
Total: %.2f GB
Used: %.2f GB
Free: %.2f GB`

	return fmt.Sprintf(format, name, inGig(status.Total), inGig(status.Used), inGig(status.Free))
}
