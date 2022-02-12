package glinkd

import (
	"encoding/hex"
	"fmt"
	"net"
	"pw_server/utils/mppc"
	"strings"
)

//普通发送
func SendBytes(conn net.Conn, data []byte) {
	fmt.Println("***********************************************************************************************")
	fmt.Println(fmt.Sprintf("普通发送：%v => %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(data))))
	conn.Write(data)
}

//加密发送
func SendEnc(conn net.Conn, data []byte) {
	fmt.Println("***********************************************************************************************")
	fmt.Println(fmt.Sprintf("发送数据：%v => %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(data))))
	if online, ok := ConnList.Load(conn); ok {
		item := online.(ConnInfo)
		//加密数据
		d := item.Msg.Enc.Get(data)
		fmt.Println(fmt.Sprintf("加密数据：%v => %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(d))))
		conn.Write(data)
	}
}

//加密压缩
func SendData(conn net.Conn, data []byte) {
	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("原始数据：%v => %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(data))))
	//压缩数据
	mp := new(mppc.Packer)
	mp.Init()
	mc := mp.Compress(data)
	fmt.Println(fmt.Sprintf("压缩数据：%v => %v", conn.RemoteAddr(), strings.ToUpper(hex.EncodeToString(mc))))
	SendEnc(conn, mc)
}
