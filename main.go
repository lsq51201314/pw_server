package main

import (
	"pw_server/authd"
	"pw_server/glinkd"
	"pw_server/global"
	"pw_server/utils/snowflake"
)

func main() {
	global.ClientVer = []byte{1, 3, 8}
	global.RunCfg.Init()
	snowflake.Init()
	authd.Init()
	glinkd.Init()
}
