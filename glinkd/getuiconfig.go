package glinkd

import (
	"net"
	"pw_server/utils/byteEx"
)

func GetUIConfig(conn net.Conn, data []byte) {
	SendEnc(conn, byteEx.HexToBytes("6982AA77DD625054151BD7816F15F4C7B51E5B2D4B1C2C07A7000FBE33D8C03FCC787185C8F1E30E9195335F047A33E2B174E3D08E4DC3989E1EB9A0B1F38024FE2BD9D7CD60F0F07C6F0FF31E84FD1300D68D3030C0E020003E0E1EB47A6C0E2C7D08F13F813B523C48ECCC2D9E685BCDAEFBBC00480064F359ABCBFDFED60D137CA4D0E7FA07CA27CD34EF309049480007A24FDDDA2232600F121C2FC240039E2083C4007888A2E501DE10413E77C1E55FE9C013F77B3FE818FF8C9EE0BA8800EFE6479E00C28A8B3A73E807A1121507D00984424C014DF2001E43E0870F31EA46680E5CACCE398F1DCCD385E2E7F796F0DE4DE7F1B031A78023D3BFBA3CF79CF2A3C48F39F43C67A8F198DB31CD1F0BC18C30A8F12269DC8503B92004909A3C0145577D00F4230D2A3DC8F12152292600F22220B786E900040820001003F02E91C3825E1FCBF800"))
}