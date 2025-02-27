package data

import (
	"encoding/json"
	"fmt"
	"os"
)


type Ports struct {
	Ports []uint16 `json:"ports"`
}

func GetPorts() ([]uint16, error) {
		file, err := os.ReadFile("internal/data/ports.json")

		if err != nil {
			 return nil, fmt.Errorf("failed to open %s: %v", "ports.json", err)
		}

		var ports Ports

		err = json.Unmarshal(file, &ports)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %v", "ports.json", err)
		}

		return ports.Ports, nil
}