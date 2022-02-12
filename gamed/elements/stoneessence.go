package elements

import (
	"pw_server/utils/structEx"
)

type StoneEssence struct {
	ID             int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdSubType      int    `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name           string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter     string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon       string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	Level          int    `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	Color          int    `elements:"name:color;type:int;size:4;text:未知;" gorm:"column:color;type:integer;default:0;not null;" json:"color"`
	Price          int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice      int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	InstallPrice   int    `elements:"name:install_price;type:int;size:4;text:未知;" gorm:"column:install_price;type:integer;default:0;not null;" json:"install_price"`
	UninstallPrice int    `elements:"name:uninstall_price;type:int;size:4;text:未知;" gorm:"column:uninstall_price;type:integer;default:0;not null;" json:"uninstall_price"`
	IdAddonDamage  int    `elements:"name:id_addon_damage;type:int;size:4;text:未知;" gorm:"column:id_addon_damage;type:integer;default:0;not null;" json:"id_addon_damage"`
	IdAddonDefence int    `elements:"name:id_addon_defence;type:int;size:4;text:未知;" gorm:"column:id_addon_defence;type:integer;default:0;not null;" json:"id_addon_defence"`
	WeaponDesc     string `elements:"name:weapon_desc;type:unicode;size:32;text:未知;" gorm:"column:weapon_desc;type:varchar(32);not null;" json:"weapon_desc"`
	ArmorDesc      string `elements:"name:armor_desc;type:unicode;size:32;text:未知;" gorm:"column:armor_desc;type:varchar(32);not null;" json:"armor_desc"`
	PileNumMax     int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid        int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType       int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (StoneEssence) TableName() string {
	return "el_stoneessence"
}

func (e *StoneEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *StoneEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
