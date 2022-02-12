package glinkd

import (
	"net"
	"time"
)

func KeepAlive(conn net.Conn, data []byte) {
	if online, ok := ConnList.Load(conn); ok {
		item := online.(ConnInfo)
		item.LastTime = time.Now().Unix()
		ConnList.Store(conn, item)
		SendEnc(conn, []byte{0xF0, 0xDE, 0}) //心跳数据 F0DE00
	}
}
