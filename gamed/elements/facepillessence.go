package elements

import (
	"pw_server/utils/structEx"
)

type FacepillEssence struct {
	ID               int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdMajorType      int     `elements:"name:id_major_type;type:int;size:4;text:未知;" gorm:"column:id_major_type;type:integer;default:0;not null;" json:"id_major_type"`
	IdSubType        int     `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name             string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter       string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon         string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	Duration         int     `elements:"name:duration;type:int;size:4;text:未知;" gorm:"column:duration;type:integer;default:0;not null;" json:"duration"`
	CameraScale      float32 `elements:"name:camera_scale;type:float;size:4;text:未知;" gorm:"column:camera_scale;type:float;default:0;not null;" json:"camera_scale"`
	CharacterComboId int     `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	Pllfiles1File    string  `elements:"name:pllfiles_1_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_1_file;type:varchar(128);not null;" json:"pllfiles_1_file"`
	Pllfiles2File    string  `elements:"name:pllfiles_2_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_2_file;type:varchar(128);not null;" json:"pllfiles_2_file"`
	Pllfiles3File    string  `elements:"name:pllfiles_3_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_3_file;type:varchar(128);not null;" json:"pllfiles_3_file"`
	Pllfiles4File    string  `elements:"name:pllfiles_4_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_4_file;type:varchar(128);not null;" json:"pllfiles_4_file"`
	Pllfiles5File    string  `elements:"name:pllfiles_5_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_5_file;type:varchar(128);not null;" json:"pllfiles_5_file"`
	Pllfiles6File    string  `elements:"name:pllfiles_6_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_6_file;type:varchar(128);not null;" json:"pllfiles_6_file"`
	Pllfiles7File    string  `elements:"name:pllfiles_7_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_7_file;type:varchar(128);not null;" json:"pllfiles_7_file"`
	Pllfiles8File    string  `elements:"name:pllfiles_8_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_8_file;type:varchar(128);not null;" json:"pllfiles_8_file"`
	Pllfiles9File    string  `elements:"name:pllfiles_9_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_9_file;type:varchar(128);not null;" json:"pllfiles_9_file"`
	Pllfiles10File   string  `elements:"name:pllfiles_10_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_10_file;type:varchar(128);not null;" json:"pllfiles_10_file"`
	Pllfiles11File   string  `elements:"name:pllfiles_11_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_11_file;type:varchar(128);not null;" json:"pllfiles_11_file"`
	Pllfiles12File   string  `elements:"name:pllfiles_12_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_12_file;type:varchar(128);not null;" json:"pllfiles_12_file"`
	Pllfiles13File   string  `elements:"name:pllfiles_13_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_13_file;type:varchar(128);not null;" json:"pllfiles_13_file"`
	Pllfiles14File   string  `elements:"name:pllfiles_14_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_14_file;type:varchar(128);not null;" json:"pllfiles_14_file"`
	Pllfiles15File   string  `elements:"name:pllfiles_15_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_15_file;type:varchar(128);not null;" json:"pllfiles_15_file"`
	Pllfiles16File   string  `elements:"name:pllfiles_16_file;type:gbk;size:128;text:未知;" gorm:"column:pllfiles_16_file;type:varchar(128);not null;" json:"pllfiles_16_file"`
	Price            int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice        int     `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	PileNumMax       int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid          int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType         int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (FacepillEssence) TableName() string {
	return "el_facepillessence"
}

func (e *FacepillEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FacepillEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
