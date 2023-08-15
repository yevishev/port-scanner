package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)


func main() {
	var (
		targetHost string
		startPort = 1
		endPort = 65536
	)
	fmt.Print("enter host address\n")
	fmt.Scan(&targetHost)
	
	fmt.Println("start of scanning")
	t := time.Now()
	var wg sync.WaitGroup
	for port := startPort; port < endPort; port++ {
		wg.Add(1)
		go func (port int)  {
			defer func ()  {
				wg.Done()
			}()
			address := fmt.Sprintf("%s:%d", targetHost, port)
			conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
			if err != nil {
				return
			}
			fmt.Printf("Port - %d - opened\n", port)
			conn.Close()
		}(port)
	}
	wg.Wait()
	fmt.Printf("end of scanning. time: %.2f\n", time.Since(t).Seconds())
}