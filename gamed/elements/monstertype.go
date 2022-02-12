package elements

import (
	"pw_server/utils/structEx"
)

type MonsterType struct {
	ID                       int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                     string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Addons1IdAddon           int     `elements:"name:addons_1_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_1_id_addon;type:integer;default:0;not null;" json:"addons_1_id_addon"`
	Addons1ProbabilityAddon  float32 `elements:"name:addons_1_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_1_probability_addon;type:float;default:0;not null;" json:"addons_1_probability_addon"`
	Addons2IdAddon           int     `elements:"name:addons_2_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_2_id_addon;type:integer;default:0;not null;" json:"addons_2_id_addon"`
	Addons2ProbabilityAddon  float32 `elements:"name:addons_2_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_2_probability_addon;type:float;default:0;not null;" json:"addons_2_probability_addon"`
	Addons3IdAddon           int     `elements:"name:addons_3_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_3_id_addon;type:integer;default:0;not null;" json:"addons_3_id_addon"`
	Addons3ProbabilityAddon  float32 `elements:"name:addons_3_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_3_probability_addon;type:float;default:0;not null;" json:"addons_3_probability_addon"`
	Addons4IdAddon           int     `elements:"name:addons_4_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_4_id_addon;type:integer;default:0;not null;" json:"addons_4_id_addon"`
	Addons4ProbabilityAddon  float32 `elements:"name:addons_4_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_4_probability_addon;type:float;default:0;not null;" json:"addons_4_probability_addon"`
	Addons5IdAddon           int     `elements:"name:addons_5_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_5_id_addon;type:integer;default:0;not null;" json:"addons_5_id_addon"`
	Addons5ProbabilityAddon  float32 `elements:"name:addons_5_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_5_probability_addon;type:float;default:0;not null;" json:"addons_5_probability_addon"`
	Addons6IdAddon           int     `elements:"name:addons_6_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_6_id_addon;type:integer;default:0;not null;" json:"addons_6_id_addon"`
	Addons6ProbabilityAddon  float32 `elements:"name:addons_6_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_6_probability_addon;type:float;default:0;not null;" json:"addons_6_probability_addon"`
	Addons7IdAddon           int     `elements:"name:addons_7_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_7_id_addon;type:integer;default:0;not null;" json:"addons_7_id_addon"`
	Addons7ProbabilityAddon  float32 `elements:"name:addons_7_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_7_probability_addon;type:float;default:0;not null;" json:"addons_7_probability_addon"`
	Addons8IdAddon           int     `elements:"name:addons_8_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_8_id_addon;type:integer;default:0;not null;" json:"addons_8_id_addon"`
	Addons8ProbabilityAddon  float32 `elements:"name:addons_8_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_8_probability_addon;type:float;default:0;not null;" json:"addons_8_probability_addon"`
	Addons9IdAddon           int     `elements:"name:addons_9_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_9_id_addon;type:integer;default:0;not null;" json:"addons_9_id_addon"`
	Addons9ProbabilityAddon  float32 `elements:"name:addons_9_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_9_probability_addon;type:float;default:0;not null;" json:"addons_9_probability_addon"`
	Addons10IdAddon          int     `elements:"name:addons_10_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_10_id_addon;type:integer;default:0;not null;" json:"addons_10_id_addon"`
	Addons10ProbabilityAddon float32 `elements:"name:addons_10_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_10_probability_addon;type:float;default:0;not null;" json:"addons_10_probability_addon"`
	Addons11IdAddon          int     `elements:"name:addons_11_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_11_id_addon;type:integer;default:0;not null;" json:"addons_11_id_addon"`
	Addons11ProbabilityAddon float32 `elements:"name:addons_11_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_11_probability_addon;type:float;default:0;not null;" json:"addons_11_probability_addon"`
	Addons12IdAddon          int     `elements:"name:addons_12_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_12_id_addon;type:integer;default:0;not null;" json:"addons_12_id_addon"`
	Addons12ProbabilityAddon float32 `elements:"name:addons_12_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_12_probability_addon;type:float;default:0;not null;" json:"addons_12_probability_addon"`
	Addons13IdAddon          int     `elements:"name:addons_13_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_13_id_addon;type:integer;default:0;not null;" json:"addons_13_id_addon"`
	Addons13ProbabilityAddon float32 `elements:"name:addons_13_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_13_probability_addon;type:float;default:0;not null;" json:"addons_13_probability_addon"`
	Addons14IdAddon          int     `elements:"name:addons_14_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_14_id_addon;type:integer;default:0;not null;" json:"addons_14_id_addon"`
	Addons14ProbabilityAddon float32 `elements:"name:addons_14_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_14_probability_addon;type:float;default:0;not null;" json:"addons_14_probability_addon"`
	Addons15IdAddon          int     `elements:"name:addons_15_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_15_id_addon;type:integer;default:0;not null;" json:"addons_15_id_addon"`
	Addons15ProbabilityAddon float32 `elements:"name:addons_15_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_15_probability_addon;type:float;default:0;not null;" json:"addons_15_probability_addon"`
	Addons16IdAddon          int     `elements:"name:addons_16_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_16_id_addon;type:integer;default:0;not null;" json:"addons_16_id_addon"`
	Addons16ProbabilityAddon float32 `elements:"name:addons_16_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_16_probability_addon;type:float;default:0;not null;" json:"addons_16_probability_addon"`
}

func (MonsterType) TableName() string {
	return "el_monstertype"
}

func (e *MonsterType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *MonsterType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
