package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on UDP port 8080
	addr := "0.0.0.0:8080"
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP server listening on", addr)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read from the connection
		n, clientAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}

		// Log the incoming message
		message := string(buffer[:n])
		fmt.Printf("Received message: %s from %s\n", message, clientAddr)

		// Send a response back to the client
		response := []byte("Message received")
		_, err = conn.WriteTo(response, clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
