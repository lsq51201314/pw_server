package structEx

import (
	"errors"
	"fmt"
	"pw_server/utils/byteEx"
	"pw_server/utils/stringEx"
	"reflect"
	"strconv"
	"strings"
)

const tagName = "elements"

func GetSize(t interface{}) (size int, err error) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get(tagName)
		str := strings.TrimSpace(stringEx.Between(tag, "size:", ";"))
		if str == "" {
			err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 未定义size标签。", val.Type(), field.Name))
			return
		}
		var num int
		if num, err = strconv.Atoi(str); err != nil {
			return
		}
		size += num
	}
	return
}

func SetData(t interface{}, data []byte) (err error) {
	val := reflect.ValueOf(t).Elem()
	offSet := 0
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get(tagName)
		tpe := strings.TrimSpace(stringEx.Between(tag, "type:", ";"))
		size := strings.TrimSpace(stringEx.Between(tag, "size:", ";"))
		if tpe == "" {
			err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 未定义type标签。", val.Type(), field.Name))
			return
		}
		if size == "" {
			err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 未定义size标签。", val.Type(), field.Name))
			return
		}
		var num int
		if num, err = strconv.Atoi(size); err != nil {
			return
		}
		station := offSet + num
		if station > len(data) {
			err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 数据长度错误(%v/%v)。", val.Type(), field.Name, len(data), station))
			return
		}
		d := data[offSet:station]
		switch tpe {
		case "int":
			tName := val.Field(i).Type().Name()
			switch tName {
			case "int":
				val.Field(i).SetInt(int64(byteEx.BytesToIntLittle(d)))
			case "float32":
				n := byteEx.BytesToIntLittle(d)
				if n < 1000000000 {
					val.Field(i).SetFloat(float64(n))
				} else {
					val.Field(i).SetFloat(float64(byteEx.ByteToFloatLittle(d)))
				}
			default:
				err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 无法转换到%v。", val.Type(), field.Name, tName))
				return
			}
		case "float":
			val.Field(i).SetFloat(float64(byteEx.ByteToFloatLittle(d)))
		case "gbk":
			val.Field(i).SetString(byteEx.ByteToStr(d))
		case "unicode":
			val.Field(i).SetString(byteEx.WStrByteToStr(d))
		default:
			err = errors.New(fmt.Sprintf("结构体:%v 字段:%v 不支持的类型%v。", val.Type(), field.Name, tpe))
			return
		}
		offSet += num
	}
	return
}
