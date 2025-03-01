package tools

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func SaveResult(r []string, h string) error {

	o := "logs/scan_result.txt"
	t := fmt.Sprintf("\n[%v] New scan result for %v...\n\n", time.Now().Format("2006-01-02 15:04:05"), h)

	file, err := os.OpenFile(o, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	d := t + strings.Join(r, "\n") + "\n"

	if _, err := file.WriteString(d); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Successfully written to", o)
	return nil
}
