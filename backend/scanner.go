package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

const scanTimeout = 10

type PortInfo struct {
	Port int
	Open bool
}

func Scan(host string, portsToScan []int) ([]PortInfo, error) {
	wg := sync.WaitGroup{}
	sem := semaphore.NewWeighted(Ulimit())
	ports := make([]PortInfo, 0)
	portsLock := sync.Mutex{}

	for _, port := range portsToScan {
		wg.Add(1)
		sem.Acquire(context.TODO(), 1)

		go func(port int) {
			defer sem.Release(1)
			defer wg.Done()

			isOpen := isPortOpen(host, port)
			pi := PortInfo{Port: port, Open: isOpen}

			portsLock.Lock()
			defer portsLock.Unlock()
			ports = append(ports, pi)
		}(port)
	}

	wg.Wait()

	return ports, nil
}

func isPortOpen(host string, port int) bool {
	tm := time.Second * time.Duration(scanTimeout)
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, tm)
	if conn != nil {
		defer conn.Close()
	}

	if err != nil {
		// log.Println(err)
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(tm + time.Second)
			// log.Printf("Restaring scan for %s on port %d", host, port)
			return isPortOpen(host, port)
		} else {
			// log.Printf("Error scanning %s:%d: %s", host, port, err)
			return false
		}
	}

	if conn == nil {
		return false
	}

	return true
}
