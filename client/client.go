package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	// Define the server address to connect to
	serverAddr := "212.2.247.89:8080"

	// Resolve the UDP address
	udpAddr, err := net.ResolveUDPAddr("udp4", serverAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a channel to wait for all goroutines to finish
	done := make(chan struct{}, 10)

	// Create a UDP connection to the server
	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer func() { done <- struct{}{} }()
			// Dial UDP connection to the server within the goroutine
			conn, err := net.DialUDP("udp", nil, udpAddr)
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			defer conn.Close()

			SendMessage(strconv.Itoa(i), conn)
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < 100000; i++ {
		<-done
	}
}

func SendMessage(msg string, conn *net.UDPConn) {
	message := []byte(msg)
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Set a timeout for receiving the server's response
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	// Receive response from the server
	buffer := make([]byte, 1024)
	_, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		// If there's a timeout or other error, print the message as not received
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Printf("Message not received by server: %s\n", msg)
		} else {
			fmt.Println("Error receiving response:", err)
		}
		return
	}

	// Print the server's response
	// fmt.Printf("Received response: %s\n", string(buffer[:n]))
}
