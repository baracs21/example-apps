package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/endpoint", EndpointHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func EndpointHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request from: %s, at: %s", req.Host, time.Now().String())

	username := os.Getenv("APP2_USERNAME")
	password := os.Getenv("APP2_PASSWORD")
	cm := os.Getenv("APP2_CM")

	resp := fmt.Sprintf("HALLO FROM ENDPOINT ;), username: %s, password: %s, cm: %s", username, password, cm)

	_, _ = io.WriteString(w, resp)
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request from: %s, at: %s", req.Host, time.Now().String())
	resp := "UP"
	_, _ = io.WriteString(w, resp)
}
