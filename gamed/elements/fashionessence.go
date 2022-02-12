package elements

import (
	"pw_server/utils/structEx"
)

type FashionEssence struct {
	ID            int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdMajorType   int    `elements:"name:id_major_type;type:int;size:4;text:未知;" gorm:"column:id_major_type;type:integer;default:0;not null;" json:"id_major_type"`
	IdSubType     int    `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name          string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Realname      string `elements:"name:realname;type:gbk;size:32;text:未知;" gorm:"column:realname;type:varchar(32);not null;" json:"realname"`
	FileMatter    string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon      string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	EquipLocation int    `elements:"name:equip_location;type:int;size:4;text:未知;" gorm:"column:equip_location;type:integer;default:0;not null;" json:"equip_location"`
	Level         int    `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	RequireLevel  int    `elements:"name:require_level;type:int;size:4;text:未知;" gorm:"column:require_level;type:integer;default:0;not null;" json:"require_level"`
	Price         int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice     int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	Gender        int    `elements:"name:gender;type:int;size:4;text:未知;" gorm:"column:gender;type:integer;default:0;not null;" json:"gender"`
	PileNumMax    int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid       int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType      int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (FashionEssence) TableName() string {
	return "el_fashionessence"
}

func (e *FashionEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FashionEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
