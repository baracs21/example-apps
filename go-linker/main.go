package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/call", CallHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func CallHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request from: %s, at: %s", req.Host, time.Now().String())

	url := req.URL.Query().Get("url")

	fmt.Printf("URL: %s", url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	_, _ = io.WriteString(w, string(responseData))
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request from: %s, at: %s", req.Host, time.Now().String())

	// Get the IP addresses associated with the system
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	ipAddress := ""

	// Loop through the addresses to find the non-localhost IPv4 address
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddress += ipnet.IP.String()
			}
		}
	}

	resp := "UP, IP-Addresses: " + ipAddress
	_, _ = io.WriteString(w, resp)
}
