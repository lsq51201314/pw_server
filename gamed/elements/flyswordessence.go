package elements

import (
	"pw_server/utils/structEx"
)

type FlyswordEssence struct {
	ID                     int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                   string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel              string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileMatter             string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon               string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	Price                  int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice              int     `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	Level                  int     `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	RequirePlayerLevelMin  int     `elements:"name:require_player_level_min;type:int;size:4;text:未知;" gorm:"column:require_player_level_min;type:integer;default:0;not null;" json:"require_player_level_min"`
	SpeedIncreaseMin       float32 `elements:"name:speed_increase_min;type:float;size:4;text:未知;" gorm:"column:speed_increase_min;type:float;default:0;not null;" json:"speed_increase_min"`
	SpeedIncreaseMax       float32 `elements:"name:speed_increase_max;type:float;size:4;text:未知;" gorm:"column:speed_increase_max;type:float;default:0;not null;" json:"speed_increase_max"`
	SpeedRushIncreaseMin   float32 `elements:"name:speed_rush_increase_min;type:float;size:4;text:未知;" gorm:"column:speed_rush_increase_min;type:float;default:0;not null;" json:"speed_rush_increase_min"`
	SpeedRushIncreaseMax   float32 `elements:"name:speed_rush_increase_max;type:float;size:4;text:未知;" gorm:"column:speed_rush_increase_max;type:float;default:0;not null;" json:"speed_rush_increase_max"`
	TimeMaxMin             float32 `elements:"name:time_max_min;type:float;size:4;text:未知;" gorm:"column:time_max_min;type:float;default:0;not null;" json:"time_max_min"`
	TimeMaxMax             float32 `elements:"name:time_max_max;type:float;size:4;text:未知;" gorm:"column:time_max_max;type:float;default:0;not null;" json:"time_max_max"`
	TimeIncreasePerElement float32 `elements:"name:time_increase_per_element;type:float;size:4;text:未知;" gorm:"column:time_increase_per_element;type:float;default:0;not null;" json:"time_increase_per_element"`
	FlyMode                int     `elements:"name:fly_mode;type:int;size:4;text:未知;" gorm:"column:fly_mode;type:integer;default:0;not null;" json:"fly_mode"`
	CharacterComboId       int     `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	PileNumMax             int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid                int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType               int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (FlyswordEssence) TableName() string {
	return "el_flyswordessence"
}

func (e *FlyswordEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FlyswordEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
