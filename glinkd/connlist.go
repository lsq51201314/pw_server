package glinkd

import (
	"pw_server/utils/rpenc"
	"sync"
)

type UserInfo struct {
	ID     int
	Name   string
	Passwd []byte
}

type KeyInfo struct {
	Token     []byte
	Client    []byte
	ClientEnc []byte
	Server    []byte
	ServerEnc []byte
}

type MsgInfo struct {
	Dec *rpenc.RPEnc
	Enc *rpenc.RPEnc
}

type ConnInfo struct {
	IsLogin  bool
	User     UserInfo
	Key      KeyInfo
	Msg      MsgInfo
	LastTime int64
}

var ConnList sync.Map
