package glinkd

import (
	"fmt"
	"net"
)

func OnClose(conn net.Conn) {
	fmt.Println(fmt.Sprintf("客户离开：%v", conn.RemoteAddr()))
	ConnList.Delete(conn)
}
