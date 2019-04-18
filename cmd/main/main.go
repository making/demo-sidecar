package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"net"
	"os/exec"
)

func main() {
	http.HandleFunc("/", respond)
	http.HandleFunc("/kill-sidecar", killSidecar)
	http.HandleFunc("/kill-main", killMain)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func reader(r io.Reader) []byte {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		log.Fatal("Read error:", err)
		return []byte{}
	}
	return buf[0:n]
}

func killSidecar(res http.ResponseWriter, req *http.Request) {
	exec.Command("pkill", "-KILL", "sidecar").Start()
}

func killMain(res http.ResponseWriter, req *http.Request) {
	exec.Command("pkill", "-KILL", "main").Start()
}

func respond(res http.ResponseWriter, req *http.Request) {
	c, err := net.Dial("unix", "/tmp/sidecar.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	msg := "hello sidecar"
	_, err = c.Write([]byte(msg))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	res.Write(reader(c))
}

