package elements

import (
	"pw_server/utils/structEx"
)

type MedicineEssence struct {
	ID           int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdMajorType  int    `elements:"name:id_major_type;type:int;size:4;text:未知;" gorm:"column:id_major_type;type:integer;default:0;not null;" json:"id_major_type"`
	IdSubType    int    `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name         string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter   string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon     string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	RequireLevel int    `elements:"name:require_level;type:int;size:4;text:未知;" gorm:"column:require_level;type:integer;default:0;not null;" json:"require_level"`
	CoolTime     int    `elements:"name:cool_time;type:int;size:4;text:未知;" gorm:"column:cool_time;type:integer;default:0;not null;" json:"cool_time"`
	HpAddTotal   int    `elements:"name:hp_add_total;type:int;size:4;text:未知;" gorm:"column:hp_add_total;type:integer;default:0;not null;" json:"hp_add_total"`
	HpAddTime    int    `elements:"name:hp_add_time;type:int;size:4;text:未知;" gorm:"column:hp_add_time;type:integer;default:0;not null;" json:"hp_add_time"`
	MpAddTotal   int    `elements:"name:mp_add_total;type:int;size:4;text:未知;" gorm:"column:mp_add_total;type:integer;default:0;not null;" json:"mp_add_total"`
	MpAddTime    int    `elements:"name:mp_add_time;type:int;size:4;text:未知;" gorm:"column:mp_add_time;type:integer;default:0;not null;" json:"mp_add_time"`
	Price        int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice    int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	PileNumMax   int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid      int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType     int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (MedicineEssence) TableName() string {
	return "el_medicineessence"
}

func (e *MedicineEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *MedicineEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
