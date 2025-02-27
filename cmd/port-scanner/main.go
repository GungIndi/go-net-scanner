package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gungindi/port-scanner/internal/data"
	"github.com/gungindi/port-scanner/internal/scanner"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	t0 := time.Now()

	var ports []uint16

	h := flag.String("host", "scanme.nmap.org", "Target host to scan")
	p := flag.String("ports", "", "Comma-separated list of ports to scan (e.g. '22,80,443')")
	flag.Parse()

	fmt.Print(*p)
	if *p != "" {
		ports = data.ParsePortList(*p)
	} else {
		var err error
		ports, err = data.GetPorts()
		if err != nil {
			log.Fatalf("Error retrieving ports: %v", err)
		}
	}

	r := scanner.ScanPorts(*h, &ports)

	log.Infof("\nAll Port Scanned!\nExecution time: %v\n", time.Since(t0))

	err := scanner.SaveResult(r, *h)
	if err != nil {
		log.Errorf("Failed to save scan results: %v", err)
	}
}
