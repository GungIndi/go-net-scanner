package main

import (
	"sync"
	"time"

	"github.com/gungindi/port-scanner/internal/scanner"
	"github.com/gungindi/port-scanner/internal/tools"
	log "github.com/sirupsen/logrus"
)

func main() {
	wg := sync.WaitGroup{}
	log.SetReportCaller(true)
	t0 := time.Now()

	var rp, rs []string

	subs, ports, h := tools.GetInput()
	wg.Add(2)

	go func() {
		defer wg.Done()
		rp = scanner.ScanPorts(&h, &ports)
	}()

	go func() {
		defer wg.Done()
		rs = scanner.ScanSubdomains(&h, &subs)
	}()

	wg.Wait()
	log.Infof("\nExecution time: %v\n", time.Since(t0))

	err := tools.SaveResult(rs, h)
	if err != nil {
		log.Errorf("Failed to save scan results: %v", err)
	}
	err = tools.SaveResult(rp, h)
	if err != nil {
		log.Errorf("Failed to save scan results: %v", err)
	}
}
