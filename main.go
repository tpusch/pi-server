package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

func inGig(size uint64) float64 {
	return float64(size) / float64(GB)
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	disk := DiskUsage("/mnt/external")
	fmt.Printf("All: %.2f GB\n", inGig(disk.All))
	fmt.Printf("Used: %.2f GB\n", inGig(disk.Used))
	fmt.Printf("Free: %.2f GB\n", inGig(disk.Free))

	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Starting Web Server")
	log.Fatal(http.ListenAndServe(":80", nil))
}
