package glinkd

import (
	"bufio"
	"fmt"
	"net"
	"pw_server/global"
)

func Init() (err error) {
	var ln net.Listener
	if ln, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", global.RunCfg.ServerCfg.Port)); err != nil {
		return
	}
	defer ln.Close()
	for {
		if conn, ok := ln.Accept(); ok == nil {
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	OnConnect(conn)
	for {
		reader := bufio.NewReader(conn)
		var buf [8192]byte
		var n int
		var ok error
		if n, ok = reader.Read(buf[:]); ok != nil {
			break
		}
		go OnReceive(conn, buf[:n])
	}
	OnClose(conn)
}
