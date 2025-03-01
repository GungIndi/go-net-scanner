package tools

import (
	"flag"
	"log"

	"github.com/gungindi/port-scanner/internal/data"
)

func GetInput() ([]string, []uint16, string) {

	h := flag.String("host", "", "Target host to scan")
	p := flag.String("ports", "", "Comma-separated list of ports to scan (e.g. '22,80,443')")
	s := flag.String("subdomains", "", "Comma-separated list of subdomains to scan (e.g. 'www,api,app')")

	flag.Parse()

	if *h == "" {
		log.Fatalf("Please provide a host to scan")
	}

	ports, err := data.GetPorts()
	if *p != "" {
		ports = data.ParsePortList(*p)
	} else if err != nil {
		log.Fatalf("Error retrieving ports: %v", err)
	}

	subs, err := data.GetSubdomains()
	if *s != "" {
		subs = data.ParseSubdomainList(*s)
	} else if err != nil {
		log.Fatalf("Error retrieving subdomains: %v", err)
	}

	return subs, ports, *h
}
