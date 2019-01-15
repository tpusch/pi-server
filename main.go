package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
)

func inGig(size uint64) float64 {
	return math.Round(float64(size) / float64(GB))
}

func kbToMb(size uint64) float64 {
	return float64(size) / float64(1024)
}

type StatusPageVariables struct {
	External    DiskString
	ExternalRaw DiskStatus
	Internal    DiskString
	MemStats    MemoryString
}

func handler(w http.ResponseWriter, r *http.Request) {

	pageVars := StatusPageVariables{
		getDiskString(DiskUsage("/mnt/external")),
		DiskUsage("/mnt/external"),
		getDiskString(DiskUsage("/")),
		getMemoryString(memUsage()),
	}

	t, err := template.ParseFiles("status.html")
	if err != nil {
		log.Println("template parsing error: ", err)
	}

	err = t.Execute(w, pageVars)
	if err != nil {
		log.Println("template execution error: ", err)
	}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":80", nil))

}
