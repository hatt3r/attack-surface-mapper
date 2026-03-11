package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanTarget(target string) []int {

	var openPorts []int
	var wg sync.WaitGroup
	// allows to wait for all goroutines to finish
	var mutex sync.Mutex
	// prevents race conditions for multiple goroutines accessing

	ports := make(chan int, 100)

	for i := 0; i < cap(ports); i++ {
		wg.Add(1)

		go worker(target, ports, &openPorts, &mutex, &wg)
		// worker goroutines starts for each channel
	}

	for port := 1; port <= 1024; port++ {
		ports <- port
	}

	close(ports)

	wg.Wait()

	return openPorts
}

// goroutine worker that scans ports from channel
func worker(target string, ports chan int, openPorts *[]int, mutex *sync.Mutex, wg *sync.WaitGroup) {

	defer wg.Done()
	// signals that the goroutine has finished its work when the function returns

	for port := range ports {

		address := fmt.Sprintf("%s:%d", target, port)

		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err == nil {
			conn.Close()

			mutex.Lock()
			//helps to prevent race conditions when multiple goroutines access at the same time
			*openPorts = append(*openPorts, port)
			mutex.Unlock()
		}
	}
}
