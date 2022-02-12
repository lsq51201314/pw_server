package elements

import (
	"pw_server/utils/structEx"
)

type TossmatterEssence struct {
	ID              int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name            string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel       string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileMatter      string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon        string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	FileFiregfx     string  `elements:"name:file_firegfx;type:gbk;size:128;text:未知;" gorm:"column:file_firegfx;type:varchar(128);not null;" json:"file_firegfx"`
	FileHitgfx      string  `elements:"name:file_hitgfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitgfx;type:varchar(128);not null;" json:"file_hitgfx"`
	FileHitsfx      string  `elements:"name:file_hitsfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitsfx;type:varchar(128);not null;" json:"file_hitsfx"`
	RequireStrength int     `elements:"name:require_strength;type:int;size:4;text:未知;" gorm:"column:require_strength;type:integer;default:0;not null;" json:"require_strength"`
	RequireAgility  int     `elements:"name:require_agility;type:int;size:4;text:未知;" gorm:"column:require_agility;type:integer;default:0;not null;" json:"require_agility"`
	RequireLevel    int     `elements:"name:require_level;type:int;size:4;text:未知;" gorm:"column:require_level;type:integer;default:0;not null;" json:"require_level"`
	DamageLow       int     `elements:"name:damage_low;type:int;size:4;text:未知;" gorm:"column:damage_low;type:integer;default:0;not null;" json:"damage_low"`
	DamageHighMin   int     `elements:"name:damage_high_min;type:int;size:4;text:未知;" gorm:"column:damage_high_min;type:integer;default:0;not null;" json:"damage_high_min"`
	DamageHighMax   int     `elements:"name:damage_high_max;type:int;size:4;text:未知;" gorm:"column:damage_high_max;type:integer;default:0;not null;" json:"damage_high_max"`
	UseTime         float32 `elements:"name:use_time;type:float;size:4;text:未知;" gorm:"column:use_time;type:float;default:0;not null;" json:"use_time"`
	AttackRange     float32 `elements:"name:attack_range;type:float;size:4;text:未知;" gorm:"column:attack_range;type:float;default:0;not null;" json:"attack_range"`
	Price           int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice       int     `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	PileNumMax      int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid         int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType        int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (TossmatterEssence) TableName() string {
	return "el_tossmatteressence"
}

func (e *TossmatterEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *TossmatterEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
