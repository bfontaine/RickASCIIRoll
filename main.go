package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var resetCursor = []byte("\033[;f")

func main() {
	var bindAddress = flag.String("bind", "0.0.0.0", "bind address")
	var bindPort = flag.Int("port", 23, "port")

	flag.Parse()

	bind := fmt.Sprintf("%s:%d", *bindAddress, *bindPort)

	l, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("Cannot bind on %s: %v", bind, err)
	}

	defer l.Close()
	log.Printf("Listening on %s", bind)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Can't accept: %v", err)
		}

		log.Printf("Someone's connected: %v", conn.RemoteAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for _, frame := range frames {
		conn.Write(resetCursor)
		conn.Write([]byte(frame))
		time.Sleep(100 * time.Millisecond)
	}

	conn.Close()
}
