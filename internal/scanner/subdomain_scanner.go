package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanSubdomain(s string, h string, c chan string, lwg *sync.WaitGroup) {
	t0 := time.Now()
	defer lwg.Done()

	addr := fmt.Sprintf("%s.%s", s, h)
	_, err := net.LookupHost(addr)

	if err != nil {
		return
	} else {
		if time.Since(t0) > time.Millisecond*300 {
			c <- fmt.Sprintf("⚠️   [SUBDOMAIN] %s is slow to respond! [Response Time: %v]", addr, time.Since(t0))
		} else {
			c <- fmt.Sprintf("✅  [SUBDOMAIN] %s is responding! [Response Time: %v]", addr, time.Since(t0))
		}
	}
}

func ScanSubdomains(h *string, subdomains *[]string) []string {
	c := make(chan string, len(*subdomains))

	lwg := sync.WaitGroup{}
	var r []string

	for _, p := range *subdomains {
		lwg.Add(1)
		go scanSubdomain(p, *h, c, &lwg)
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
