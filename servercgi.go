package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Stdout.Write([]byte("Content-type: text/html\n\n"))
	buf := make([]byte, 1024*100)
	n, err := os.Stdin.Read(buf)
	checkError(err)
	buf = accessDNS(buf[:n])
	os.Stdout.Write(buf)
}

func accessDNS(message []byte) []byte {
	// Please set your DNS server below.
	serverIP := "8.8.8.8"
	serverPort := "53"

	conn, err := net.Dial("tcp", serverIP+":"+serverPort)
	if err != nil {
		return nil
	}
	defer conn.Close()

	conn.Write(message)
	readBuf := make([]byte, 1024*100)
	readlen, _ := conn.Read(readBuf)
	return readBuf[:readlen]
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
