package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// Checks availibility
func CheckPort(host string, port int) (bool, time.Duration) {
	address := fmt.Sprintf("%s:%d", host, port)
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return false, 0
	}

	conn.Close()
	connectDuration := time.Since(startTime)
	return true, connectDuration
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Invalid port number: %v\n", err)
		return
	}

	fmt.Printf("%.22v | Address: %s:%d - start checking...\n", time.Now(), host, portNumber)
	// Open file for logging
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer f.Close()

	// Logging
	logger := log.New(f, "", log.LstdFlags)
	for {
		ok, lat := CheckPort(host, portNumber)
		fLat := float64(lat) / float64(time.Millisecond)
		if ok {
			if fLat > 250 {
				logger.Printf("%.22v | High latency: %.f\n", time.Now(), fLat)
			}
			fmt.Printf("%.22v | OK (%.f ms)\n", time.Now(), fLat)
			time.Sleep(5 * time.Second)
		} else {
			fmt.Printf("%.22v | NOT OK\n\a", time.Now())
			logger.Printf("%.22v | Port %d is sleeping\n", time.Now(), portNumber)
			time.Sleep(2 * time.Second)
		}
	}
}
