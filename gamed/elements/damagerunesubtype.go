package elements

import (
	"pw_server/utils/structEx"
)

type DamageruneSubType struct {
	ID   int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
}

func (DamageruneSubType) TableName() string {
	return "el_damagerunesubtype"
}

func (e *DamageruneSubType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *DamageruneSubType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
