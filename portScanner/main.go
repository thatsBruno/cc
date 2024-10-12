package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port", 0, "port")

	flag.Parse()

	scanPort(*port)
}

func scanPort(port int) {
	addr := fmt.Sprintf("localhost:%d", port)

	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}
}
