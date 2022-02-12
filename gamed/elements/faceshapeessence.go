package elements

import (
	"pw_server/utils/structEx"
)

type FaceShapeEssence struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileShape        string `elements:"name:file_shape;type:gbk;size:128;text:未知;" gorm:"column:file_shape;type:varchar(128);not null;" json:"file_shape"`
	FileIcon         string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	ShapePartId      int    `elements:"name:shape_part_id;type:int;size:4;text:未知;" gorm:"column:shape_part_id;type:integer;default:0;not null;" json:"shape_part_id"`
	CharacterComboId int    `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	GenderId         int    `elements:"name:gender_id;type:int;size:4;text:未知;" gorm:"column:gender_id;type:integer;default:0;not null;" json:"gender_id"`
	VisualizeId      int    `elements:"name:visualize_id;type:int;size:4;text:未知;" gorm:"column:visualize_id;type:integer;default:0;not null;" json:"visualize_id"`
	UserData         int    `elements:"name:user_data;type:int;size:4;text:未知;" gorm:"column:user_data;type:integer;default:0;not null;" json:"user_data"`
	FacepillOnly     int    `elements:"name:facepill_only;type:int;size:4;text:未知;" gorm:"column:facepill_only;type:integer;default:0;not null;" json:"facepill_only"`
}

func (FaceShapeEssence) TableName() string {
	return "el_faceshapeessence"
}

func (e *FaceShapeEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FaceShapeEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
