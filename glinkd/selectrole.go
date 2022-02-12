package glinkd

import (
	"net"
	"pw_server/utils/byteEx"
)

func SelectRole(conn net.Conn, data []byte) {
	SendEnc(conn, byteEx.HexToBytes("47E81AF000"))
}
