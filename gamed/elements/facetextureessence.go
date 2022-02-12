package elements

import (
	"pw_server/utils/structEx"
)

type FaceTextureEssence struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileBaseTex      string `elements:"name:file_base_tex;type:gbk;size:128;text:未知;" gorm:"column:file_base_tex;type:varchar(128);not null;" json:"file_base_tex"`
	FileHighTex      string `elements:"name:file_high_tex;type:gbk;size:128;text:未知;" gorm:"column:file_high_tex;type:varchar(128);not null;" json:"file_high_tex"`
	FileIcon         string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	TexPartId        int    `elements:"name:tex_part_id;type:int;size:4;text:未知;" gorm:"column:tex_part_id;type:integer;default:0;not null;" json:"tex_part_id"`
	CharacterComboId int    `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	GenderId         int    `elements:"name:gender_id;type:int;size:4;text:未知;" gorm:"column:gender_id;type:integer;default:0;not null;" json:"gender_id"`
	VisualizeId      int    `elements:"name:visualize_id;type:int;size:4;text:未知;" gorm:"column:visualize_id;type:integer;default:0;not null;" json:"visualize_id"`
	UserData         int    `elements:"name:user_data;type:int;size:4;text:未知;" gorm:"column:user_data;type:integer;default:0;not null;" json:"user_data"`
	FacepillOnly     int    `elements:"name:facepill_only;type:int;size:4;text:未知;" gorm:"column:facepill_only;type:integer;default:0;not null;" json:"facepill_only"`
}

func (FaceTextureEssence) TableName() string {
	return "el_facetextureessence"
}

func (e *FaceTextureEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FaceTextureEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
