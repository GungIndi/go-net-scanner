package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(p uint16, h string, c chan string, wg *sync.WaitGroup) {
	t0 := time.Now()
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", h, p)
	conn, err := net.DialTimeout("tcp", addr, time.Second*1)
	if err != nil {
		c <- fmt.Sprintf("Port %d is Closed! : %v", p, time.Since(t0))
		return
	} else {
		defer conn.Close()
		c <- fmt.Sprintf("Port %d is Open! : %v", p, time.Since(t0))
	}
}

func ScanPorts(h string, ports *[]uint16) []string {
	wg := sync.WaitGroup{}
	c := make(chan string, len(*ports))

	var r []string

	for _, p := range *ports {
		wg.Add(1)
		go scanPort(p, h, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
		r = append(r, i)
	}
	return r
}
