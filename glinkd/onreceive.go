package glinkd

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func OnReceive(conn net.Conn, data []byte) {
	if online, ok := ConnList.Load(conn); ok {
		item := online.(ConnInfo)
		var r error
		if item.IsLogin {
			data = item.Msg.Dec.Get(data)
		}
		fmt.Println("***********************************************************************************************")
		fmt.Println(fmt.Sprintf("数据到达：%v <= %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(data))))
		if r != nil {
			fmt.Println(fmt.Sprintf("解密失败：%v", r))
		} else {
			t := int(data[0])
			f := RPCalls[t]
			if f == nil {
				fmt.Println(fmt.Sprintf("无法识别：%v <= %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(data))))
			} else {
				f(conn, data)
			}
		}
	}
}
