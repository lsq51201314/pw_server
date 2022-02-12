package elements

import (
	"pw_server/utils/structEx"
)

type ProjectileEssence struct {
	ID                    int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Type                  int    `elements:"name:type;type:int;size:4;text:未知;" gorm:"column:type;type:integer;default:0;not null;" json:"type"`
	Name                  string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel             string `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileMatter            string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon              string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	FileFiregfx           string `elements:"name:file_firegfx;type:gbk;size:128;text:未知;" gorm:"column:file_firegfx;type:varchar(128);not null;" json:"file_firegfx"`
	FileHitgfx            string `elements:"name:file_hitgfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitgfx;type:varchar(128);not null;" json:"file_hitgfx"`
	FileHitsfx            string `elements:"name:file_hitsfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitsfx;type:varchar(128);not null;" json:"file_hitsfx"`
	RequireWeaponLevelMin int    `elements:"name:require_weapon_level_min;type:int;size:4;text:未知;" gorm:"column:require_weapon_level_min;type:integer;default:0;not null;" json:"require_weapon_level_min"`
	RequireWeaponLevelMax int    `elements:"name:require_weapon_level_max;type:int;size:4;text:未知;" gorm:"column:require_weapon_level_max;type:integer;default:0;not null;" json:"require_weapon_level_max"`
	DamageEnhance         int    `elements:"name:damage_enhance;type:int;size:4;text:未知;" gorm:"column:damage_enhance;type:integer;default:0;not null;" json:"damage_enhance"`
	DamageScaleEnhance    int    `elements:"name:damage_scale_enhance;type:int;size:4;text:未知;" gorm:"column:damage_scale_enhance;type:integer;default:0;not null;" json:"damage_scale_enhance"`
	Price                 int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice             int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	IdAddon0              int    `elements:"name:id_addon0;type:int;size:4;text:未知;" gorm:"column:id_addon0;type:integer;default:0;not null;" json:"id_addon0"`
	IdAddon1              int    `elements:"name:id_addon1;type:int;size:4;text:未知;" gorm:"column:id_addon1;type:integer;default:0;not null;" json:"id_addon1"`
	IdAddon2              int    `elements:"name:id_addon2;type:int;size:4;text:未知;" gorm:"column:id_addon2;type:integer;default:0;not null;" json:"id_addon2"`
	IdAddon3              int    `elements:"name:id_addon3;type:int;size:4;text:未知;" gorm:"column:id_addon3;type:integer;default:0;not null;" json:"id_addon3"`
	PileNumMax            int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid               int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType              int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (ProjectileEssence) TableName() string {
	return "el_projectileessence"
}

func (e *ProjectileEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *ProjectileEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
