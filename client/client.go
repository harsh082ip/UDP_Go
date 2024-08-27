package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	// Define the server address to connect to
	serverAddr := "127.0.0.1:8080"

	// Resolve the UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Dial UDP connection to the server
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	for i := 1; i <= 100000; i++ {
		message := []byte(strconv.Itoa(i))
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Set a timeout for receiving the server's response
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		// Receive response from the server
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			// If there's a timeout or other error, print the message as not received
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Printf("Message not received by server: %s\n", string(message))
			} else {
				fmt.Println("Error receiving response:", err)
			}
			continue
		}

		// Print the server's response
		fmt.Printf("Received response: %s\n", string(buffer[:n]))
	}
}
