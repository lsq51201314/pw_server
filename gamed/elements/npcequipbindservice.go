package elements

import (
	"pw_server/utils/structEx"
)

type NpcEquipbindService struct {
	ID           int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name         string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdObjectNeed int    `elements:"name:id_object_need;type:int;size:4;text:未知;" gorm:"column:id_object_need;type:integer;default:0;not null;" json:"id_object_need"`
	Price        int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
}

func (NpcEquipbindService) TableName() string {
	return "el_npcequipbindservice"
}

func (e *NpcEquipbindService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcEquipbindService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
