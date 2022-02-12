package elements

import (
	"pw_server/utils/structEx"
)

type GmGeneratorEssence struct {
	ID         int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdType     int    `elements:"name:id_type;type:int;size:4;text:未知;" gorm:"column:id_type;type:integer;default:0;not null;" json:"id_type"`
	Name       string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon   string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	IdObject   int    `elements:"name:id_object;type:int;size:4;text:未知;" gorm:"column:id_object;type:integer;default:0;not null;" json:"id_object"`
	PileNumMax int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid    int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType   int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (GmGeneratorEssence) TableName() string {
	return "el_gmgeneratoressence"
}

func (e *GmGeneratorEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *GmGeneratorEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
