package gochat

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	nickname string   //昵称
	inputStr string   //用户输入的聊天内容
	data     MesgData //传输的数据
)

//设置昵称
func setNickName() {
	fmt.Print("\nPlease set nickname : ")
	fmt.Scanln(&nickname)
}

//连接服务端
func Client() {

	setNickName()

	conn, err := net.Dial(CONFIG_SERVER_PROTOCOL, CONFIG_SERVER_IP+CONFIG_SERVER_PORT)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	//客户端发送数据到服务端
	go inputToServer(conn)
	receiveFromSever(conn)
	//time.Sleep(1 * time.Hour)
}

func inputToServer(conn net.Conn) {
	for {
		fmt.Print("\n")
		_, err := fmt.Scanln(&inputStr)
		if err != nil {
			log.Print(err)
			continue
		}

		curTime := time.Now().Unix()

		data = MesgData{
			Uid:      nickname + strconv.FormatInt(curTime, 10),
			Nickname: nickname,
			Text:     inputStr,
			InfoType: MessageInfoType,
			Platform: ClientPlatform,
			CurTime:  curTime,
		}

		conn.Write(data.ToJson())
	}
}

func receiveFromSever(conn net.Conn) {
	var d MesgData
	var clientInput = make([]byte, 2048)
	for {
		num, err := conn.Read(clientInput)
		if err != nil {
			log.Print(err)
			//客户端连接断开
			if err == io.EOF {
				closeConnect(conn)
				break
			}
			continue
		}
		if json.Unmarshal(clientInput[:num], &d) != nil {
			log.Print(err)
		} else {
			fmt.Printf("\n[%s : %s]\n", d.Nickname, d.Text)
		}

	}

}
