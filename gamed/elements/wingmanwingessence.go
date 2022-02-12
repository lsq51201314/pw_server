package elements

import (
	"pw_server/utils/structEx"
)

type WingmanwingEssence struct {
	ID                    int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                  string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel             string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileMatter            string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon              string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	Price                 int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice             int     `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	RequirePlayerLevelMin int     `elements:"name:require_player_level_min;type:int;size:4;text:未知;" gorm:"column:require_player_level_min;type:integer;default:0;not null;" json:"require_player_level_min"`
	SpeedIncrease         float32 `elements:"name:speed_increase;type:float;size:4;text:未知;" gorm:"column:speed_increase;type:float;default:0;not null;" json:"speed_increase"`
	MpLaunch              int     `elements:"name:mp_launch;type:int;size:4;text:未知;" gorm:"column:mp_launch;type:integer;default:0;not null;" json:"mp_launch"`
	MpPerSecond           int     `elements:"name:mp_per_second;type:int;size:4;text:未知;" gorm:"column:mp_per_second;type:integer;default:0;not null;" json:"mp_per_second"`
	PileNumMax            int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid               int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType              int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (WingmanwingEssence) TableName() string {
	return "el_wingmanwingessence"
}

func (e *WingmanwingEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *WingmanwingEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
