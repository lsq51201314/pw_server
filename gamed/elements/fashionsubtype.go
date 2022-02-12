package elements

import (
	"pw_server/utils/structEx"
)

type FashionSubType struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	EquipFashionMask int    `elements:"name:equip_fashion_mask;type:int;size:4;text:未知;" gorm:"column:equip_fashion_mask;type:integer;default:0;not null;" json:"equip_fashion_mask"`
}

func (FashionSubType) TableName() string {
	return "el_fashionsubtype"
}

func (e *FashionSubType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FashionSubType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
