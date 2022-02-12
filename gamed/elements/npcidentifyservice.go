package elements

import (
	"pw_server/utils/structEx"
)

type NpcIdentifyService struct {
	ID   int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Fee  int    `elements:"name:fee;type:int;size:4;text:未知;" gorm:"column:fee;type:integer;default:0;not null;" json:"fee"`
}

func (NpcIdentifyService) TableName() string {
	return "el_npcidentifyservice"
}

func (e *NpcIdentifyService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcIdentifyService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
