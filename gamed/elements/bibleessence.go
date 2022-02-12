package elements

import (
	"pw_server/utils/structEx"
)

type BibleEssence struct {
	ID         int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name       string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon   string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	IdAddons1  int    `elements:"name:id_addons_1;type:int;size:4;text:未知;" gorm:"column:id_addons_1;type:integer;default:0;not null;" json:"id_addons_1"`
	IdAddons2  int    `elements:"name:id_addons_2;type:int;size:4;text:未知;" gorm:"column:id_addons_2;type:integer;default:0;not null;" json:"id_addons_2"`
	IdAddons3  int    `elements:"name:id_addons_3;type:int;size:4;text:未知;" gorm:"column:id_addons_3;type:integer;default:0;not null;" json:"id_addons_3"`
	IdAddons4  int    `elements:"name:id_addons_4;type:int;size:4;text:未知;" gorm:"column:id_addons_4;type:integer;default:0;not null;" json:"id_addons_4"`
	IdAddons5  int    `elements:"name:id_addons_5;type:int;size:4;text:未知;" gorm:"column:id_addons_5;type:integer;default:0;not null;" json:"id_addons_5"`
	IdAddons6  int    `elements:"name:id_addons_6;type:int;size:4;text:未知;" gorm:"column:id_addons_6;type:integer;default:0;not null;" json:"id_addons_6"`
	IdAddons7  int    `elements:"name:id_addons_7;type:int;size:4;text:未知;" gorm:"column:id_addons_7;type:integer;default:0;not null;" json:"id_addons_7"`
	IdAddons8  int    `elements:"name:id_addons_8;type:int;size:4;text:未知;" gorm:"column:id_addons_8;type:integer;default:0;not null;" json:"id_addons_8"`
	IdAddons9  int    `elements:"name:id_addons_9;type:int;size:4;text:未知;" gorm:"column:id_addons_9;type:integer;default:0;not null;" json:"id_addons_9"`
	IdAddons10 int    `elements:"name:id_addons_10;type:int;size:4;text:未知;" gorm:"column:id_addons_10;type:integer;default:0;not null;" json:"id_addons_10"`
	Price      int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice  int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	PileNumMax int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid    int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType   int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (BibleEssence) TableName() string {
	return "el_bibleessence"
}

func (e *BibleEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *BibleEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
