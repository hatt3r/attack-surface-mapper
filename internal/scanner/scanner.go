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
	var mutex sync.Mutex

	ports := make(chan int, 100)

	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(target, ports, &openPorts, &mutex, &wg)
	}

	for port := 1; port <= 1024; port++ {
		ports <- port
	}

	close(ports)

	wg.Wait()

	return openPorts
}

func worker(target string, ports chan int, openPorts *[]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range ports {

		address := fmt.Sprintf("%s:%d", target, port)

		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err == nil {
			conn.Close()

			mutex.Lock()
			*openPorts = append(*openPorts, port)
			mutex.Unlock()
		}
	}
}
