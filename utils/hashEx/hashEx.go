package hashEx

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

//GetStrMd5 取字符串MD5值
func GetStrMd5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

//GetBytesMd5 取字节组MD5值
func GetBytesMd5(data []byte) []byte {
	m := md5.New()
	m.Write(data)
	return m.Sum(nil)
}

//GetStrHmac 取字符串Hmac值
func GetStrHmac(data, key string) string {
	d := []byte(data)
	h := hmac.New(md5.New, []byte(key))
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}

//GetBytesHmac 取字节组Hmac值
func GetBytesHmac(d, key []byte) (data []byte, err error) {
	defer func() {
		if ok := recover(); ok != nil {
			err = ok.(error)
		}
	}()

	h := hmac.New(md5.New, key)
	h.Write(d)
	data = h.Sum(nil)
	return
}
