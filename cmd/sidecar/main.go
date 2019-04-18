package main

import (
	"fmt"
	"log"
	"net"
)

func sidecarServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := c.Read(buf)
		if err != nil {
			return
		}
        fmt.Println("Received: " + string(buf))
		_, err = c.Write([]byte("Sidecar received your data"))
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	l, err := net.Listen("unix", "/tmp/sidecar.sock")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	fmt.Println("Start sidecar...")
	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go sidecarServer(fd)
	}
}
