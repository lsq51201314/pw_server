package glinkd

import (
	"net"
	"pw_server/utils/octets"
)

func SendErr(conn net.Conn, code int, msg string) {
	md := []byte(msg)
	ml := len(md)
	data := new(octets.Octets)
	data.AddByte(0x05)
	data.AddByte(byte(ml + 2))
	data.AddByte(byte(code))
	data.AddByte(byte(ml))
	data.AddBytes(md, false, 0)
	SendBytes(conn, data.GetBytes())
}
