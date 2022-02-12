package glinkd

import (
	"fmt"
	"net"
	"pw_server/global"
	"pw_server/utils/byteEx"
	"pw_server/utils/octets"
	"time"
)

//01 16 10 000000E0 FFFFFFFF 1D5A1B3E5698C6F3 00 010308 00
func OnConnect(conn net.Conn) {
	fmt.Println(fmt.Sprintf("客户进入：%v", conn.RemoteAddr()))
	//双倍信息
	v := 0
	if global.DoubleExp {
		v += 1
	}
	if global.DoubleSP {
		v += 8
	}
	if global.DoubleObject {
		v += 4
	}
	if global.DoubleMoney {
		v += 2
	}
	//密钥信息
	k := new(octets.Octets)
	k.AddBytes([]byte{0, 0, 0, byte(224 + v)}, false, 0)
	k.AddBytes(byteEx.IntToBytesLittle(-1), false, 0)
	k.AddBytes(byteEx.RandBytes(8), false, 0)
	//配置信息
	l := new(octets.Octets)
	l.AddBytes([]byte{0x01, 0x16, 0x10}, false, 0)
	l.AddBytes(k.GetBytes(), false, 0)
	l.AddByte(0)
	l.AddBytes(global.ClientVer, false, 0)
	l.AddByte(0)
	//加入列表
	var item ConnInfo
	item.Key.Token = k.GetBytes()
	item.LastTime = time.Now().Unix()
	ConnList.LoadOrStore(conn, item)
	//发送配置
	SendBytes(conn, l.GetBytes())
	return
}
