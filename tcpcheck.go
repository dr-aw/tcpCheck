package main

import (
	"flag"
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

var latency = flag.Int("l", 250, "latency in ms to log")
var tOkInt = flag.Int("tok", 5, "duration in seconds after OK connect")
var tNotOkInt = flag.Int("tnot", 2, "duration in seconds after NOT OK connect")

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: go run tcpCheck.go [-l latency] [-tok] [-tnot] <host> <port>")
		return
	}
	if *latency < 1 || *tOkInt < 1 || *tNotOkInt < 1 {
		fmt.Println("Only positive flags are permitted")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Invalid port number: %v\n", err)
		return
	}

	tcpLogger(host, portNumber)
}

func tcpLogger(host string, portNumber int) {
	fmt.Printf("%.22v | Address: %s:%d - start checking...\n", time.Now(), host, portNumber)
	fmt.Printf("Latency for warning: %d ms, tOk: %d s, tNotOk: %d s\n________________\n", *latency, *tOkInt, *tNotOkInt)

	// Open file for logging
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer f.Close()

	// Logging
	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("\nLogging started at %.22v.\n________________")
	for {
		ok, lat := CheckPort(host, portNumber)
		fLat := float64(lat) / float64(time.Millisecond)
		if ok {
			if fLat > float64(*latency) {
				logger.Printf("%.22v | High latency: %.f\n", time.Now(), fLat)
			}
			fmt.Printf("%.22v | OK (%.f ms)\n", time.Now(), fLat)
			time.Sleep(time.Duration(*tOkInt) * time.Second)
		} else {
			fmt.Printf("%.22v | NOT OK!\n\a", time.Now())
			logger.Printf("%.22v | Port %d is sleeping!\n", time.Now(), portNumber)
			time.Sleep(time.Duration(*tNotOkInt) * time.Second)
		}
	}
}
