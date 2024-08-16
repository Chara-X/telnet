package main

import (
	"bufio"
	"io"
	"net"
	"os"
)

func main() {
	var conn, _ = net.Dial("tcp", os.Args[1])
	defer conn.Close()
	var scanner = bufio.NewScanner(os.Stdin)
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	for scanner.Scan() {
		conn.Write(scanner.Bytes())
		conn.Write([]byte("\n"))
	}
}
