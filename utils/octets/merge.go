package octets

import "bytes"

func mergeBytes(pBytes ...[]byte) (data []byte) {
	l := len(pBytes)
	s := make([][]byte, l)
	for index := 0; index < l; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	data = bytes.Join(s, sep)
	return
}
