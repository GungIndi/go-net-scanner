package data

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ports []uint16

func GetPorts() ([]uint16, error) {
	data, err := os.ReadFile("internal/data/ports.txt")

	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", "ports.txt", err)
	}

	var ports Ports
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		port, err := strconv.ParseUint(strings.TrimSpace(line), 10, 16)
		if err != nil {
			fmt.Printf("Skipping invalid port: %s\n", line)
			continue
		}
		ports = append(ports, uint16(port))
	}

	return ports, nil
}

func ParsePortList(input string) []uint16 {
	var ports []uint16
	for _, p := range strings.Split(input, ",") {
		var port uint16
		fmt.Sscanf(p, "%d", &port)
		ports = append(ports, port)
	}
	return ports
}