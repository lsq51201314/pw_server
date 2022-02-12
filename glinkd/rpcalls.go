package glinkd

import "net"

var RPCalls = map[int]func(conn net.Conn, data []byte){
	0x02: UserLogin,
	0x03: KeyExchange,
	0x52: RoleList,
	0x5A: KeepAlive,
	0x54: CreateRole,
	//0x46: SelectRole,
	//0x48: EnterWorld,
	//0x80: SetHelpStates,
	//0x83: GetHelpStates_Re,
	//0x68: GetUIConfig,
}
