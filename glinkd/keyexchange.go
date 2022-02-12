package glinkd

import (
	"net"
	"pw_server/global"
	"pw_server/utils/byteEx"
	"pw_server/utils/hashEx"
	"pw_server/utils/octets"
	"pw_server/utils/rpenc"
	"time"

	"go.uber.org/zap"
)

//03 12 10 FAE51FB49BA43E81F785F976E5977141 00
func KeyExchange(conn net.Conn, data []byte) {
	if online, ok := ConnList.Load(conn); ok {
		item := online.(ConnInfo)
		item.Key.Server = data[3:19]
		p := new(octets.Octets)
		p.AddBytes(item.User.Passwd, false, 0)
		p.AddBytes(item.Key.Server, false, 0)
		var encErr error
		item.Key.ServerEnc, encErr = hashEx.GetBytesHmac(p.GetBytes(), []byte(item.User.Name))
		if encErr != nil {
			SendErr(conn, global.ERR_LOGINFAIL, "Invalid login.")
			zap.L().Error("无法生成加密密钥。", zap.String("host", conn.RemoteAddr().String()), zap.Error(encErr), zap.String("username", item.User.Name))
			return
		}
		item.Msg.Enc = new(rpenc.RPEnc)
		item.Msg.Enc.Init(item.Key.ServerEnc)
		item.LastTime = time.Now().Unix()
		ConnList.Store(conn, item)
		//04 1B 00000020 00000012 0098967F 01 0098967F 00000000 610B5277 0000
		d := new(octets.Octets)
		d.AddBytes([]byte{0x04, 0x1B}, false, 0)
		d.AddBytes(byteEx.IntToBytesBig(item.User.ID), false, 0)
		d.AddBytes(byteEx.IntToBytesBig(0x12), false, 0)
		d.AddBytes(byteEx.IntToBytesBig(9999999), false, 0)
		d.AddByte(1)
		d.AddBytes(byteEx.IntToBytesBig(9999999), false, 0)
		d.AddBytes(byteEx.IntToBytesBig(0), false, 0)
		d.AddBytes(byteEx.IntToBytesBig(1628131959), false, 0)
		d.AddByte(0)
		d.AddByte(0)
		SendData(conn, d.GetBytes())
	}
}
