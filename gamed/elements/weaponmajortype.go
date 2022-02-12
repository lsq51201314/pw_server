package elements

import (
	"pw_server/utils/structEx"
)

type WeaponMajorType struct {
	ID   int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
}

func (WeaponMajorType) TableName() string {
	return "el_weaponmajortype"
}

func (e *WeaponMajorType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *WeaponMajorType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
