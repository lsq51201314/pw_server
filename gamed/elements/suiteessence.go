package elements

import (
	"pw_server/utils/structEx"
)

type SuiteEssence struct {
	ID             int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name           string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	MaxEquips      int    `elements:"name:max_equips;type:int;size:4;text:未知;" gorm:"column:max_equips;type:integer;default:0;not null;" json:"max_equips"`
	Equipments1Id  int    `elements:"name:equipments_1_id;type:int;size:4;text:未知;" gorm:"column:equipments_1_id;type:integer;default:0;not null;" json:"equipments_1_id"`
	Equipments2Id  int    `elements:"name:equipments_2_id;type:int;size:4;text:未知;" gorm:"column:equipments_2_id;type:integer;default:0;not null;" json:"equipments_2_id"`
	Equipments3Id  int    `elements:"name:equipments_3_id;type:int;size:4;text:未知;" gorm:"column:equipments_3_id;type:integer;default:0;not null;" json:"equipments_3_id"`
	Equipments4Id  int    `elements:"name:equipments_4_id;type:int;size:4;text:未知;" gorm:"column:equipments_4_id;type:integer;default:0;not null;" json:"equipments_4_id"`
	Equipments5Id  int    `elements:"name:equipments_5_id;type:int;size:4;text:未知;" gorm:"column:equipments_5_id;type:integer;default:0;not null;" json:"equipments_5_id"`
	Equipments6Id  int    `elements:"name:equipments_6_id;type:int;size:4;text:未知;" gorm:"column:equipments_6_id;type:integer;default:0;not null;" json:"equipments_6_id"`
	Equipments7Id  int    `elements:"name:equipments_7_id;type:int;size:4;text:未知;" gorm:"column:equipments_7_id;type:integer;default:0;not null;" json:"equipments_7_id"`
	Equipments8Id  int    `elements:"name:equipments_8_id;type:int;size:4;text:未知;" gorm:"column:equipments_8_id;type:integer;default:0;not null;" json:"equipments_8_id"`
	Equipments9Id  int    `elements:"name:equipments_9_id;type:int;size:4;text:未知;" gorm:"column:equipments_9_id;type:integer;default:0;not null;" json:"equipments_9_id"`
	Equipments10Id int    `elements:"name:equipments_10_id;type:int;size:4;text:未知;" gorm:"column:equipments_10_id;type:integer;default:0;not null;" json:"equipments_10_id"`
	Equipments11Id int    `elements:"name:equipments_11_id;type:int;size:4;text:未知;" gorm:"column:equipments_11_id;type:integer;default:0;not null;" json:"equipments_11_id"`
	Equipments12Id int    `elements:"name:equipments_12_id;type:int;size:4;text:未知;" gorm:"column:equipments_12_id;type:integer;default:0;not null;" json:"equipments_12_id"`
	Addons1Id      int    `elements:"name:addons_1_id;type:int;size:4;text:未知;" gorm:"column:addons_1_id;type:integer;default:0;not null;" json:"addons_1_id"`
	Addons2Id      int    `elements:"name:addons_2_id;type:int;size:4;text:未知;" gorm:"column:addons_2_id;type:integer;default:0;not null;" json:"addons_2_id"`
	Addons3Id      int    `elements:"name:addons_3_id;type:int;size:4;text:未知;" gorm:"column:addons_3_id;type:integer;default:0;not null;" json:"addons_3_id"`
	Addons4Id      int    `elements:"name:addons_4_id;type:int;size:4;text:未知;" gorm:"column:addons_4_id;type:integer;default:0;not null;" json:"addons_4_id"`
	Addons5Id      int    `elements:"name:addons_5_id;type:int;size:4;text:未知;" gorm:"column:addons_5_id;type:integer;default:0;not null;" json:"addons_5_id"`
	Addons6Id      int    `elements:"name:addons_6_id;type:int;size:4;text:未知;" gorm:"column:addons_6_id;type:integer;default:0;not null;" json:"addons_6_id"`
	Addons7Id      int    `elements:"name:addons_7_id;type:int;size:4;text:未知;" gorm:"column:addons_7_id;type:integer;default:0;not null;" json:"addons_7_id"`
	Addons8Id      int    `elements:"name:addons_8_id;type:int;size:4;text:未知;" gorm:"column:addons_8_id;type:integer;default:0;not null;" json:"addons_8_id"`
	Addons9Id      int    `elements:"name:addons_9_id;type:int;size:4;text:未知;" gorm:"column:addons_9_id;type:integer;default:0;not null;" json:"addons_9_id"`
	Addons10Id     int    `elements:"name:addons_10_id;type:int;size:4;text:未知;" gorm:"column:addons_10_id;type:integer;default:0;not null;" json:"addons_10_id"`
	Addons11Id     int    `elements:"name:addons_11_id;type:int;size:4;text:未知;" gorm:"column:addons_11_id;type:integer;default:0;not null;" json:"addons_11_id"`
	FileGfx        string `elements:"name:file_gfx;type:gbk;size:128;text:未知;" gorm:"column:file_gfx;type:varchar(128);not null;" json:"file_gfx"`
}

func (SuiteEssence) TableName() string {
	return "el_suiteessence"
}

func (e *SuiteEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *SuiteEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
