package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn) {
	fmt.Print("SERVIDOR:\n")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func CreateServer() {
	l, err := net.Listen("tcp4", ":50001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
