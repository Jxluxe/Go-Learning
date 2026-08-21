/*
 A simple UDP Server written in Go,
 it listens for UDP packets and prints them to the console.
 */


package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8080") // this line resolves the UDP Addr to port 8080, returns an address of UDP endpoint.

	if err != nil { // if their is an error, print the error and return, return means the function exits.
		fmt.Println(err)
		return
	}



	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		buf := make([]byte, 1024) // creates a buffer of 1024 bytes to store incoming data
		n, err := conn.Read(buf) // n means the number of bytes read from the buffer
		if err != nil {
			fmt.Println(err)
			continue // continue means the loop will loop again
		}
		fmt.Println(string(buf[:n]))
		// we use : in front of the n specifically to get the first n bytes of the buffer, we must stringify it to print to console.

	}
}
