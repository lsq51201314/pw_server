package elements

import (
	"pw_server/utils/structEx"
)

type CustomizedataEssence struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileData         string `elements:"name:file_data;type:gbk;size:128;text:未知;" gorm:"column:file_data;type:varchar(128);not null;" json:"file_data"`
	CharacterComboId int    `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	GenderId         int    `elements:"name:gender_id;type:int;size:4;text:未知;" gorm:"column:gender_id;type:integer;default:0;not null;" json:"gender_id"`
}

func (CustomizedataEssence) TableName() string {
	return "el_customizedataessence"
}

func (e *CustomizedataEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *CustomizedataEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
