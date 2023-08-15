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
	fmt.Println("start scanning")

	for port := startPort; port < endPort; port++ {
		address := fmt.Sprintf("%s:%d", targetHost, port)
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
		if err != nil {
			continue
		}
		fmt.Printf("Port - %d - opened\n", port)
		conn.Close()
	}
	fmt.Printf("end of scanning. time: %.2f\n", time.Since(t).Seconds())
}