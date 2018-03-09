package mzgochat

import (
	"fmt"
	"log"
	"net"
)

var (
	//存放客户端ip列表
	clientList = make(map[string]bool)
)

func Server() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(conn)

	}
}

func intoRoom(name, ip string) {
	log.Printf("\n有新用户 %s 进入聊天室 IP => %s", name, ip)
	clientList[ip] = true
}

func handleConnection(conn net.Conn) {
	//判断是否是新用户
	ip := conn.RemoteAddr().String()
	if _, have := clientList[ip]; !have {
		intoRoom("MZero", ip)
	}
	//
	var clientInput = make([]byte, 2048)
	for {
		_, err := conn.Read(clientInput)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("收到客户端的信息 : %s\n", string(clientInput))
		conn.Write([]byte(string(clientInput)))
	}

}
