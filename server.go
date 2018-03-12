package gochat

import (
	"encoding/json"
	"io"
	"log"
	"net"
)

var (
	//存放客户端数据列表
	clientsList = make(map[net.Addr]MesgData)
	//chanels
	dataCh = make(chan MesgData)
)

//关闭连接
func closeConnect(conn net.Conn) {
	addr := conn.RemoteAddr()
	conn.Close()
	dataCh <- MesgData{
		Addr:      addr,
		EventType: DeleteType,
	}
}

//对客户端数据进行操作
func doWithData() {
	for data := range dataCh {
		addr := data.Addr
		switch data.EventType {
		case AddOrUpdateType:
			//判断是否是新用户
			if _, ok := clientsList[addr]; ok {
				//update
			} else {
				//add
			}
			clientsList[addr] = data
		case DeleteType:
			if _, ok := clientsList[addr]; ok {
				delete(clientsList, addr)
			}
		case NoticeType:
			noticeToAllClients(data)
		}
	}
}

//发数据给所有的客户端
func noticeToAllClients(data MesgData) {
	for _, v := range clientsList {
		//定时检测心跳
		//
		v.Conn.Write(data.ToJson())
	}
}

//开启服务端
func StartServer() {
	//定时检测心跳
	//处理数据
	go doWithData()
	//
	listen, err := net.Listen(CONFIG_SERVER_PROTOCOL, CONFIG_SERVER_PORT)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			//客户端连接断开
			if err == io.EOF {
				closeConnect(conn)
			}
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	var clientInput = make([]byte, 2048)
	var data MesgData
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
		//
		if json.Unmarshal(clientInput[:num], &data) != nil {
			log.Print(err)
		} else {
			//数据变更
			data.Addr = conn.RemoteAddr()
			data.Conn = conn
			data.EventType = AddOrUpdateType
			dataCh <- data
			//发送给所有的客户端
			dataCh <- MesgData{
				Uid:       data.Uid,
				Nickname:  data.Nickname,
				Text:      data.Text,
				InfoType:  data.InfoType,
				Platform:  data.Platform,
				EventType: NoticeType,
				CurTime:   data.CurTime,
			}
		}
	}

}
