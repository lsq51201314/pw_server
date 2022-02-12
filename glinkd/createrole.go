package glinkd

import (
	"net"
	"pw_server/authd"
	"pw_server/gamedbd"
	"pw_server/utils/byteEx"
	"time"
)

func CreateRole(conn net.Conn, data []byte) {
	role := new(gamedbd.RoleInfo)
	role.ID = time.Now().UnixNano()
	role.UserID = byteEx.BytesToIntBig(byteEx.CutBytes(data, 3, 4))
	role.RoleID = byteEx.BytesToIntBig(byteEx.CutBytes(data, 7, 4))
	role.Gender = int(byteEx.CutBytes(data, 15, 1)[0]) //0:男 1:女
	role.Race = byteEx.BytesToUIntBig(byteEx.CutBytes(data, 16, 2))
	role.Level = byteEx.BytesToIntBig(byteEx.CutBytes(data, 18, 4))
	role.Level2 = byteEx.BytesToIntBig(byteEx.CutBytes(data, 22, 4))
	nlen := int(byteEx.CutBytes(data, 26, 1)[0])
	role.Name = byteEx.WStrByteToStr(byteEx.CutBytes(data, 27, nlen))
	role.CustomData = byteEx.BytesToHex(byteEx.CutBytes(data, 29+nlen, 176))
	authd.MDB.Create(&role)
}
