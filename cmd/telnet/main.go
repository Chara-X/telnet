package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	var conn, _ = net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	defer conn.Close()
	var req, res = bufio.NewScanner(os.Stdin), bufio.NewScanner(conn)
	res.Split(split)
	for res.Scan() {
		fmt.Print(res.Text())
		req.Scan()
		conn.Write([]byte(req.Text() + "\n"))
		if req.Text() == "exit" {
			break
		}
	}
}
func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return len(data), data, nil
	}
	if i := bytes.Index(data, []byte("> ")); i >= 0 {
		return i + 2, data[0 : i+2], nil
	}
	return 0, nil, nil
}
