package main

import (
	"time"

	"github.com/gungindi/port-scanner/internal/data"
	"github.com/gungindi/port-scanner/internal/scanner"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	t0 := time.Now()

	ports, err := data.GetPorts()
	if err != nil {
		log.Fatalf("Error retrieving ports: %v", err)
	}

	h := "scanme.nmap.org"
	r := scanner.ScanPorts(h, &ports)

	log.Infof("\nAll Port Scanned!\nExecution time: %v\n", time.Since(t0))

	err = scanner.SaveResult(r, h)
	if err != nil {
		log.Errorf("Failed to save scan results: %v", err)
	}
}
