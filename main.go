package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	serverIP := flag.String("host", "0.0.0.0", "Adresse IP")
	flag.Parse()
	StartScan(*serverIP)
	fmt.Println("Scan terminé !")
}

func StartScan(serverIP string) {
	var wg sync.WaitGroup
	ports := make(chan int, 70) 
	for i := 0; i < 100; i++ {
		go func() {
			for port := range ports {
				ScanPort(serverIP, port, &wg)
			}
		}()
	}

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		ports <- i
	}
	close(ports) 
	wg.Wait()
	fmt.Println("Scan terminé !")
}


func ScanPort(serverIP string, port int, wg *sync.WaitGroup) {
	defer wg.Done() 
	address := fmt.Sprintf("%s:%d", serverIP, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return  
	}
	conn.Close()
	color.Green("Port %d is open", port)
}
