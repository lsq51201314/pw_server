package elements

import (
	"pw_server/utils/structEx"
)

type MineType struct {
	ID   int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
}

func (MineType) TableName() string {
	return "el_minetype"
}

func (e *MineType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *MineType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
