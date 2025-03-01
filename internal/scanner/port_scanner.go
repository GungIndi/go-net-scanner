package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(p uint16, h string, c chan string, lwg *sync.WaitGroup) {
	t0 := time.Now()
	defer lwg.Done()
	addr := fmt.Sprintf("%s:%d", h, p)
	conn, err := net.DialTimeout("tcp", addr, time.Second*1)

	if err != nil {
		return
	} else {
		if time.Since(t0) > time.Millisecond*300 {
			c <- fmt.Sprintf("⚠️   [PORT] %s:%d is slow to respond! [Response Time: %v]", h, p, time.Since(t0))
		} else {
			c <- fmt.Sprintf("✅  [PORT] %s:%d is Open! [Response Time: %v]", h, p, time.Since(t0))
		}
		conn.Close()
	}
}

func ScanPorts(h *string, ports *[]uint16) []string {
	c := make(chan string, len(*ports))

	lwg := sync.WaitGroup{}

	var r []string

	for _, p := range *ports {
		lwg.Add(1)
		go scanPort(p, *h, c, &lwg)
	}

	go func() {
		lwg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
		r = append(r, i)
	}
	return r
}
