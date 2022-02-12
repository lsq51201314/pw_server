package octets

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"pw_server/utils/byteEx"
	"strings"
)

type Octets struct {
	Data []byte
}

//AddByte 添加一个字节
func (e *Octets) AddByte(val byte) {
	e.Data = mergeBytes(e.Data, []byte{val})
}

//AddBytes 添加一组字节
//val:添加的字节组
//addLen:是否添加长度
//fill:填充最终大小，小于等于0则不填充
func (e *Octets) AddBytes(val []byte, addLen bool, fill int) {
	l := len(val)
	if fill > 0 && fill-l > 0 {
		f := make([]byte, fill-l)
		val = mergeBytes(val, f)
	}
	if addLen {
		e.Data = mergeBytes(e.Data, []byte{byte(len(val))}, val)
	} else {
		e.Data = mergeBytes(e.Data, val)
	}
}

//AddUInt 添加短整数
func (e *Octets) AddUInt(val int) {
	e.Data = mergeBytes(e.Data, byteEx.UIntToBytes(val))
}

//AddInt 添加整数
func (e *Octets) AddInt(val int) {
	e.Data = mergeBytes(e.Data, byteEx.IntToBytesLittle(val))
}

//AddFloat 添加小数
func (e *Octets) AddFloat(val float32) {
	e.Data = mergeBytes(e.Data, byteEx.FloatToByteLittle(val))
}

//AddString 添加字符串
//val:添加的字符串
//addLen:是否添加长度
//fill:填充最终大小，小于等于0则不填充
func (e *Octets) AddString(val string, addLen bool, fill int) {
	d, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(val)), simplifiedchinese.GBK.NewEncoder()))
	l := len(d)
	if fill > 0 && fill-l > 0 {
		f := make([]byte, fill-l)
		d = mergeBytes(d, f)
	}
	if addLen {
		e.Data = mergeBytes(e.Data, []byte{byte(len(d))}, d)
	} else {
		e.Data = mergeBytes(e.Data, d)
	}
}

//AddWString 转到Unicode添加字符串
//val:添加的字符串
//addLen:是否添加长度
//fill:填充最终大小，小于等于0则不填充
func (e *Octets) AddWString(val string, addLen bool, fill int) {
	d := fmt.Sprintf("%U", []rune(val))
	d = strings.Replace(d, "U+", "", -1)
	d = strings.Replace(d, "[", "", -1)
	d = strings.Replace(d, "]", "", -1)
	a := strings.Split(d, " ")
	var r []byte
	for i := 0; i < len(a); i++ {
		b, _ := hex.DecodeString(a[i])
		b = byteEx.BytesReverse(b)
		r = mergeBytes(r, b)
	}
	l := len(r)
	if fill > 0 && fill-l > 0 {
		f := make([]byte, fill-l)
		r = mergeBytes(r, f)
	}
	if addLen {
		e.Data = mergeBytes(e.Data, []byte{byte(len(r))}, r)
	} else {
		e.Data = mergeBytes(e.Data, r)
	}
}

//GetBytes 取字节组
func (e *Octets) GetBytes() []byte {
	return e.Data
}
