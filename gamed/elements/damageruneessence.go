package elements

import (
	"pw_server/utils/structEx"
)

type DamageruneEssence struct {
	ID                    int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdSubType             int    `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name                  string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter            string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon              string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	DamageType            int    `elements:"name:damage_type;type:int;size:4;text:未知;" gorm:"column:damage_type;type:integer;default:0;not null;" json:"damage_type"`
	Price                 int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice             int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	RequireWeaponLevelMin int    `elements:"name:require_weapon_level_min;type:int;size:4;text:未知;" gorm:"column:require_weapon_level_min;type:integer;default:0;not null;" json:"require_weapon_level_min"`
	RequireWeaponLevelMax int    `elements:"name:require_weapon_level_max;type:int;size:4;text:未知;" gorm:"column:require_weapon_level_max;type:integer;default:0;not null;" json:"require_weapon_level_max"`
	DamageIncreased       int    `elements:"name:damage_increased;type:int;size:4;text:未知;" gorm:"column:damage_increased;type:integer;default:0;not null;" json:"damage_increased"`
	PileNumMax            int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid               int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType              int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (DamageruneEssence) TableName() string {
	return "el_damageruneessence"
}

func (e *DamageruneEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *DamageruneEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
