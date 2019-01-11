package main

import (
	"fmt"
	"log"
	"net/http"
)

func inGig(size uint64) float64 {
	return float64(size) / float64(GB)
}

func kbToMb(size uint64) float64 {
	return float64(size) / float64(1024)
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := `<head>
<meta http-equiv="refresh" content="5">
</head>
<h1>PI Stats</h1>
<pre><font size="5">`
	message += addDisk("External", DiskUsage("/mnt/external"))
	message += addDisk("Internal", DiskUsage("/"))
	message += addMem(memUsage())
	message += "</font></pre>"

	_, _ = fmt.Fprint(w, message, r.URL.Path[1:])
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
