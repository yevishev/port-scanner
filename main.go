package main

import (
	"fmt"
	"net"
	"time"
)


func main() {
	t := time.Now()
	targetHost := "127.0.0.1"
	startPort := 1
	endPort := 65536
	ch := make(chan struct{}, endPort)
	fmt.Println("start of scanning")

	for port := startPort; port < endPort; port++ {
		go func (port int)  {
			defer func ()  {
				ch <- struct{}{}
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
	for i := 1; i < endPort; i++ {
		<- ch
	}
	fmt.Printf("end of scanning. time: %.2f\n", time.Since(t).Seconds())
}