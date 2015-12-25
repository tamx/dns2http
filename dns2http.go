package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

var serverurl string = ""

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "dns2http [port] [serverurl]\n")
		return
	}
	service := ":" + os.Args[1]
	serverurl = os.Args[2]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func accessDNS(message []byte) []byte {
	req, _ := http.NewRequest("POST",
		serverurl,
		bytes.NewReader(message),
	)
	req.Header.Set("Connection", "close")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return body
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	messageBuf := make([]byte, 1024*100)
	messageLen, _ := conn.Read(messageBuf)
	messageBuf = accessDNS(messageBuf[:messageLen])
	conn.Write(messageBuf)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s\n", err.Error())
		os.Exit(1)
	}
}
