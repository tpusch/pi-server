package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
	Used  uint64 `json:"used"`
}

func getValue(line string) uint64 {
	cols := strings.Fields(line)
	i, _ := strconv.Atoi(cols[1])
	return uint64(i)
}

func memUsage() (mem Memory) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "MemTotal") {
			mem.Total = getValue(text)
		} else if strings.Contains(text, "MemFree") {
			mem.Free = getValue(text)
		}
	}

	mem.Used = mem.Total - mem.Free

	return
}

func addMem(memory Memory) string {
	format := `MemStatus:
Total: %.2f MB
Used: %.2f MB
Free: %.2f MB

`
	return fmt.Sprintf(format, kbToMb(memory.Total), kbToMb(memory.Used), kbToMb(memory.Free))
}
