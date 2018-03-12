package gochat

import (
	"encoding/json"
	"log"
	"net"
)

const ( //消息发送平台
	ClientPlatform = iota
	ServerPlatform
)

const ( //消息类型
	MessageInfoType = iota
	ServerInfoType
	ClientInfoType
)

const ( //其他自定义事件类型
	AddOrUpdateType = iota
	DeleteType
	NoticeType
)

type MesgData struct {
	Uid       string   `json:"uid"`
	Nickname  string   `json:"nickname"`
	Text      string   `json:"text"`
	InfoType  int      `json:"info_type"`
	Addr      net.Addr `json:"-"`
	Conn      net.Conn `json:"-"`
	Platform  int      `json:"playform"`
	CurTime   int64    `json:"cur_time"`
	EventType int      `json:"event_type,omitempty"`
}

func (m MesgData) ToJson() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
	}
	return bytes
}
