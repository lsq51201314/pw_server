package elements

import (
	"pw_server/utils/structEx"
)

type NpcDecomposeService struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdDecomposeSkill int    `elements:"name:id_decompose_skill;type:int;size:4;text:未知;" gorm:"column:id_decompose_skill;type:integer;default:0;not null;" json:"id_decompose_skill"`
}

func (NpcDecomposeService) TableName() string {
	return "el_npcdecomposeservice"
}

func (e *NpcDecomposeService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcDecomposeService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
