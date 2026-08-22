/*
 A simple UDP Server written in Go,
 it listens for UDP packets and prints them to the console.

 Our Server currently recieves the message "UDP MESSAGE" in bytes from the Client, and stringifies it.
*/

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8080") // this line resolves the UDP Addr to port 8080, returns an address of UDP endpoint.

	if err != nil { // if their is an error, print the error and return, return means the function exits.
		log.Fatal(err)
		return
	}

	conn, err := net.ListenUDP("udp", addr) // listens for UDP packets at the address defined above

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024) // creates a buffer of 1024 bytes to store incoming data
		n, err := conn.Read(buf)  // n means the number of bytes read from the buffer
		if err != nil {
			log.Fatal(err)
			continue // continue means the loop will loop again
		}
		fmt.Println(string(buf[:n]))
		// we use : in front of the n specifically to get the first n bytes of the buffer, we must stringify it to print to console.

	}
}
