package mzgochat

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	nickname string //昵称
	inputStr string //用户输入的聊天内容
	data     CSData //传输的数据
)

//设置昵称
func setNickName() {
	fmt.Print("\nPlease set nickname : ")
	fmt.Scanln(&nickname)
}

//连接服务端
func Client() {

	setNickName()

	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	//客户端发送数据到服务端
	go inputToServer(conn)
	go receiveFromSever(conn)
	time.Sleep(1 * time.Hour)
}

func inputToServer(conn net.Conn) {
	for {
		fmt.Print("\n")
		_, err := fmt.Scanln(&inputStr)
		if err != nil {
			log.Print(err)
			continue
		}

		data = CSData{
			NickName: nickname,
			Info:     inputStr,
			InfoType: MessageInfoType,
			Platform: ClientPlatform,
		}

		conn.Write(data.ToJson())
	}
}

func receiveFromSever(conn net.Conn) {

	var clientInput = make([]byte, 2048)
	for {
		_, err := conn.Read(clientInput)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("Server已经收到你发的消息 : %s\n", string(clientInput))
	}

}
