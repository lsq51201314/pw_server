package elements

import (
	"pw_server/utils/structEx"
)

type DecorationSubType struct {
	ID        int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name      string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	EquipMask int    `elements:"name:equip_mask;type:int;size:4;text:未知;" gorm:"column:equip_mask;type:integer;default:0;not null;" json:"equip_mask"`
}

func (DecorationSubType) TableName() string {
	return "el_decorationsubtype"
}

func (e *DecorationSubType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *DecorationSubType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
