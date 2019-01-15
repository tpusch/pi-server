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
	Total float64 `json:"total"`
	Free  float64 `json:"free"`
	Used  float64 `json:"used"`
}

type MemoryString struct {
	Total string
	Free  string
	Used  string
}

func getMemoryString(memory Memory) (memString MemoryString) {
	memString.Total = fmt.Sprintf("%.2f MB", memory.Total)
	memString.Used = fmt.Sprintf("%.2f MB", memory.Used)
	memString.Free = fmt.Sprintf("%.2f MB", memory.Free)
	return
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
			mem.Total = kbToMb(getValue(text))
		} else if strings.Contains(text, "MemFree") {
			mem.Free = kbToMb(getValue(text))
		}
	}

	mem.Used = mem.Total - mem.Free

	return
}
