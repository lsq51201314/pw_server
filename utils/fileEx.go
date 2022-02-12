package utils

import "os"

func GetFileByte(filepath string, offset, length int64) (data []byte, err error) {
	var file *os.File
	if file, err = os.Open(filepath); err == nil {
		buf := make([]byte, length)
		var n int
		if n, err = file.ReadAt(buf, offset); err == nil {
			data = buf[:n]
		}
		file.Close()
	}
	return
}
