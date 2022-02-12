package glinkd

import (
	"net"
	"pw_server/authd"
	"pw_server/global"
	"pw_server/utils/byteEx"
	"pw_server/utils/hashEx"
	"pw_server/utils/octets"
	"pw_server/utils/rpenc"
	"time"

	"go.uber.org/zap"
)

//02 13 01 31 10 CD9BC8612B0419B28A65EA483BA1B16B
func UserLogin(conn net.Conn, data []byte) {
	if online, ok := ConnList.Load(conn); ok {
		item := online.(ConnInfo)
		//登录信息
		nLen := byteEx.CutBytes(data, 2, 1)
		name := byteEx.CutBytes(data, 3, int(nLen[0]))
		username := string(name)
		pLen := byteEx.CutBytes(data, 3+int(nLen[0]), 1)
		passwd := byteEx.CutBytes(data, 4+int(nLen[0]), int(pLen[0]))
		//查询用户
		user := new(authd.Users)
		db := authd.MDB.Where("username = ?", username).First(&user)
		if db.Error != nil {
			SendErr(conn, global.ERR_INVALID_PASSWORD, "Invalid username.")
			zap.L().Error("无法查询用户信息。", zap.String("host", conn.RemoteAddr().String()), zap.Error(db.Error), zap.String("username", username))
			return
		}
		//生成密码
		hPasswd, hpErr := hashEx.GetBytesHmac(item.Key.Token, byteEx.HexToBytes(user.Password))
		if hpErr != nil {
			SendErr(conn, global.ERR_LOGINFAIL, "Invalid login.")
			zap.L().Error("无法生成对比密码。", zap.String("host", conn.RemoteAddr().String()), zap.Error(hpErr), zap.String("username", username))
			return
		}
		//对比密码
		if byteEx.BytesToHex(passwd) != byteEx.BytesToHex(hPasswd) {
			SendErr(conn, global.ERR_INVALID_PASSWORD, "Invalid login.")
			zap.L().Warn("用户密码错误。", zap.String("host", conn.RemoteAddr().String()), zap.String("username", username))
			return
		}
		//登录信息
		item.User.ID = user.ID
		item.User.Name = user.UserName
		item.User.Passwd = passwd
		item.Key.Client = byteEx.RandBytes(16)
		k := new(octets.Octets)
		k.AddBytes(item.User.Passwd, false, 0)
		k.AddBytes(item.Key.Client, false, 0)
		var encErr error
		item.Key.ClientEnc, encErr = hashEx.GetBytesHmac(k.GetBytes(), []byte(item.User.Name))
		if encErr != nil {
			SendErr(conn, global.ERR_LOGINFAIL, "Invalid login.")
			zap.L().Error("无法生成解密密钥。", zap.String("host", conn.RemoteAddr().String()), zap.Error(encErr), zap.String("username", username))
			return
		}
		item.Msg.Dec = new(rpenc.RPEnc)
		item.Msg.Dec.Init(item.Key.ClientEnc)
		item.LastTime = time.Now().Unix()
		item.IsLogin = true
		ConnList.Store(conn, item)
		//03 12 10 9627EE526881CBB7AD8582030AC9358C 00
		d := new(octets.Octets)
		d.AddBytes([]byte{0x03, 0x12, 0x10}, false, 0)
		d.AddBytes(item.Key.Client, false, 0)
		d.AddByte(0)
		SendBytes(conn, d.GetBytes())
	}
}
