package elements

import (
	"pw_server/utils/structEx"
)

type QuiverEssence struct {
	ID           int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdSubType    int    `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name         string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter   string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon     string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	IdProjectile int    `elements:"name:id_projectile;type:int;size:4;text:未知;" gorm:"column:id_projectile;type:integer;default:0;not null;" json:"id_projectile"`
	NumMin       int    `elements:"name:num_min;type:int;size:4;text:未知;" gorm:"column:num_min;type:integer;default:0;not null;" json:"num_min"`
	NumMax       int    `elements:"name:num_max;type:int;size:4;text:未知;" gorm:"column:num_max;type:integer;default:0;not null;" json:"num_max"`
}

func (QuiverEssence) TableName() string {
	return "el_quiveressence"
}

func (e *QuiverEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *QuiverEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
