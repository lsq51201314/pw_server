package elements

import (
	"pw_server/utils/structEx"
)

type ColorpickerEssence struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileColorpicker  string `elements:"name:file_colorpicker;type:gbk;size:128;text:未知;" gorm:"column:file_colorpicker;type:varchar(128);not null;" json:"file_colorpicker"`
	ColorPartId      int    `elements:"name:color_part_id;type:int;size:4;text:未知;" gorm:"column:color_part_id;type:integer;default:0;not null;" json:"color_part_id"`
	CharacterComboId int    `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	GenderId         int    `elements:"name:gender_id;type:int;size:4;text:未知;" gorm:"column:gender_id;type:integer;default:0;not null;" json:"gender_id"`
}

func (ColorpickerEssence) TableName() string {
	return "el_colorpickeressence"
}

func (e *ColorpickerEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *ColorpickerEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
