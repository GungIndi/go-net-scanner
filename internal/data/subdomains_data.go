package data

import (
	"fmt"
	"os"
	"strings"
)

type Subdomains []string

func GetSubdomains() ([]string, error) {
	data, err := os.ReadFile("internal/data/subdomains.txt")

	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", "names.txt", err)
	}

	var sub Subdomains
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		sub = append(sub, string(line))
	}

	return sub, nil
}

func ParseSubdomainList(input string) []string {
	var subs []string
	subs = append(subs, strings.Split(input, ",")...)
	return subs
}
