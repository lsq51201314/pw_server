package byteEx

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"math"
	"math/rand"
	"pw_server/utils/snowflake"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//随机
func RandBytes(size int) (data []byte) {
	rand.Seed(snowflake.SFGenID())
	for i := 0; i < size; i++ {
		data = append(data, byte(rand.Intn(255)))
	}
	return
}

//倒序
func BytesReverse(data []byte) []byte {
	for from, to := 0, len(data)-1; from < to; from, to = from+1, to-1 {
		data[from], data[to] = data[to], data[from]
	}
	return data
}

//取值
func CutBytes(d []byte, offset, length int) (data []byte) {
	data = d[offset : offset+length]
	return
}

//小端
func UIntToBytes(n int) []byte {
	x := uint16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToUIntBig(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

//小端
func IntToBytesLittle(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//大端
func IntToBytesBig(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//小端
func BytesToIntLittle(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}

//大端
func BytesToIntBig(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

//小端
func FloatToByteLittle(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

//大端
func FloatToByteBig(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

//小端
func ByteToFloatLittle(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

//大端
func ByteToFloatBig(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

//转码
func WStrByteToStr(wsbyte []byte) string {
	c := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	d, ok := c.Bytes(wsbyte)
	if ok != nil {
		return ""
	}
	d = bytes.TrimRight(d, "\x00")
	return string(d)
}

//转码
func ByteToStr(sbyte []byte) string {
	d, ok := ioutil.ReadAll(transform.NewReader(bytes.NewReader(sbyte), simplifiedchinese.GBK.NewDecoder()))
	if ok != nil {
		return ""
	}
	d = bytes.TrimRight(d, "\x00")
	return string(d)
}

//拷贝
func BlockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("源缓冲区 索引超出范围")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("目标缓冲区 索引超出范围")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}

//HexToBytes 十六进制转字节组
func HexToBytes(hexStr string) (data []byte) {
	data, _ = hex.DecodeString(hexStr)
	return
}

//BytesToHex 字节组转十六进制
func BytesToHex(data []byte) (hexStr string) {
	hexStr = strings.ToUpper(hex.EncodeToString(data))
	return
}
