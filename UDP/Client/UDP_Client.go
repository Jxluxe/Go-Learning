/*
 UDP Client to interact with UDP_Server.go
 Will send UDP packets(bytes) to the server,
 which will be stringified and printed by the Server into the console.
 */

package main

import (
	"net"
	"log"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080") // resolving the UDP Address returns a UDP Address and an error, resolving converts UDP string into a UDP Address(8080)

	if err != nil {
		log.Fatal(err)

	}

	conn, err := net.DialUDP("udp", nil, addr) // establishes connection to UDP Server

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close() // defers closing the connection until the function exits

	if _, err := conn.Write([]byte("UDP MESSAGE")); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write(buf[:n]); err != nil {
		log.Fatal(err)
	}
}
