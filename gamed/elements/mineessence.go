package elements

import (
	"pw_server/utils/structEx"
)

type MineEssence struct {
	ID                     int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdType                 int     `elements:"name:id_type;type:int;size:4;text:未知;" gorm:"column:id_type;type:integer;default:0;not null;" json:"id_type"`
	Name                   string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Level                  int     `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	LevelRequired          int     `elements:"name:level_required;type:int;size:4;text:未知;" gorm:"column:level_required;type:integer;default:0;not null;" json:"level_required"`
	IdEquipmentRequired    int     `elements:"name:id_equipment_required;type:int;size:4;text:未知;" gorm:"column:id_equipment_required;type:integer;default:0;not null;" json:"id_equipment_required"`
	EliminateTool          int     `elements:"name:eliminate_tool;type:int;size:4;text:未知;" gorm:"column:eliminate_tool;type:integer;default:0;not null;" json:"eliminate_tool"`
	TimeMin                int     `elements:"name:time_min;type:int;size:4;text:未知;" gorm:"column:time_min;type:integer;default:0;not null;" json:"time_min"`
	TimeMax                int     `elements:"name:time_max;type:int;size:4;text:未知;" gorm:"column:time_max;type:integer;default:0;not null;" json:"time_max"`
	Exp                    int     `elements:"name:exp;type:int;size:4;text:未知;" gorm:"column:exp;type:integer;default:0;not null;" json:"exp"`
	Skillpoint             int     `elements:"name:skillpoint;type:int;size:4;text:未知;" gorm:"column:skillpoint;type:integer;default:0;not null;" json:"skillpoint"`
	FileModel              string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	Materials1Id           int     `elements:"name:materials_1_id;type:int;size:4;text:未知;" gorm:"column:materials_1_id;type:integer;default:0;not null;" json:"materials_1_id"`
	Materials1Probability  float32 `elements:"name:materials_1_probability;type:float;size:4;text:未知;" gorm:"column:materials_1_probability;type:float;default:0;not null;" json:"materials_1_probability"`
	Materials2Id           int     `elements:"name:materials_2_id;type:int;size:4;text:未知;" gorm:"column:materials_2_id;type:integer;default:0;not null;" json:"materials_2_id"`
	Materials2Probability  float32 `elements:"name:materials_2_probability;type:float;size:4;text:未知;" gorm:"column:materials_2_probability;type:float;default:0;not null;" json:"materials_2_probability"`
	Materials3Id           int     `elements:"name:materials_3_id;type:int;size:4;text:未知;" gorm:"column:materials_3_id;type:integer;default:0;not null;" json:"materials_3_id"`
	Materials3Probability  float32 `elements:"name:materials_3_probability;type:float;size:4;text:未知;" gorm:"column:materials_3_probability;type:float;default:0;not null;" json:"materials_3_probability"`
	Materials4Id           int     `elements:"name:materials_4_id;type:int;size:4;text:未知;" gorm:"column:materials_4_id;type:integer;default:0;not null;" json:"materials_4_id"`
	Materials4Probability  float32 `elements:"name:materials_4_probability;type:float;size:4;text:未知;" gorm:"column:materials_4_probability;type:float;default:0;not null;" json:"materials_4_probability"`
	Materials5Id           int     `elements:"name:materials_5_id;type:int;size:4;text:未知;" gorm:"column:materials_5_id;type:integer;default:0;not null;" json:"materials_5_id"`
	Materials5Probability  float32 `elements:"name:materials_5_probability;type:float;size:4;text:未知;" gorm:"column:materials_5_probability;type:float;default:0;not null;" json:"materials_5_probability"`
	Materials6Id           int     `elements:"name:materials_6_id;type:int;size:4;text:未知;" gorm:"column:materials_6_id;type:integer;default:0;not null;" json:"materials_6_id"`
	Materials6Probability  float32 `elements:"name:materials_6_probability;type:float;size:4;text:未知;" gorm:"column:materials_6_probability;type:float;default:0;not null;" json:"materials_6_probability"`
	Materials7Id           int     `elements:"name:materials_7_id;type:int;size:4;text:未知;" gorm:"column:materials_7_id;type:integer;default:0;not null;" json:"materials_7_id"`
	Materials7Probability  float32 `elements:"name:materials_7_probability;type:float;size:4;text:未知;" gorm:"column:materials_7_probability;type:float;default:0;not null;" json:"materials_7_probability"`
	Materials8Id           int     `elements:"name:materials_8_id;type:int;size:4;text:未知;" gorm:"column:materials_8_id;type:integer;default:0;not null;" json:"materials_8_id"`
	Materials8Probability  float32 `elements:"name:materials_8_probability;type:float;size:4;text:未知;" gorm:"column:materials_8_probability;type:float;default:0;not null;" json:"materials_8_probability"`
	Materials9Id           int     `elements:"name:materials_9_id;type:int;size:4;text:未知;" gorm:"column:materials_9_id;type:integer;default:0;not null;" json:"materials_9_id"`
	Materials9Probability  float32 `elements:"name:materials_9_probability;type:float;size:4;text:未知;" gorm:"column:materials_9_probability;type:float;default:0;not null;" json:"materials_9_probability"`
	Materials10Id          int     `elements:"name:materials_10_id;type:int;size:4;text:未知;" gorm:"column:materials_10_id;type:integer;default:0;not null;" json:"materials_10_id"`
	Materials10Probability float32 `elements:"name:materials_10_probability;type:float;size:4;text:未知;" gorm:"column:materials_10_probability;type:float;default:0;not null;" json:"materials_10_probability"`
	Materials11Id          int     `elements:"name:materials_11_id;type:int;size:4;text:未知;" gorm:"column:materials_11_id;type:integer;default:0;not null;" json:"materials_11_id"`
	Materials11Probability float32 `elements:"name:materials_11_probability;type:float;size:4;text:未知;" gorm:"column:materials_11_probability;type:float;default:0;not null;" json:"materials_11_probability"`
	Materials12Id          int     `elements:"name:materials_12_id;type:int;size:4;text:未知;" gorm:"column:materials_12_id;type:integer;default:0;not null;" json:"materials_12_id"`
	Materials12Probability float32 `elements:"name:materials_12_probability;type:float;size:4;text:未知;" gorm:"column:materials_12_probability;type:float;default:0;not null;" json:"materials_12_probability"`
	Materials13Id          int     `elements:"name:materials_13_id;type:int;size:4;text:未知;" gorm:"column:materials_13_id;type:integer;default:0;not null;" json:"materials_13_id"`
	Materials13Probability float32 `elements:"name:materials_13_probability;type:float;size:4;text:未知;" gorm:"column:materials_13_probability;type:float;default:0;not null;" json:"materials_13_probability"`
	Materials14Id          int     `elements:"name:materials_14_id;type:int;size:4;text:未知;" gorm:"column:materials_14_id;type:integer;default:0;not null;" json:"materials_14_id"`
	Materials14Probability float32 `elements:"name:materials_14_probability;type:float;size:4;text:未知;" gorm:"column:materials_14_probability;type:float;default:0;not null;" json:"materials_14_probability"`
	Materials15Id          int     `elements:"name:materials_15_id;type:int;size:4;text:未知;" gorm:"column:materials_15_id;type:integer;default:0;not null;" json:"materials_15_id"`
	Materials15Probability float32 `elements:"name:materials_15_probability;type:float;size:4;text:未知;" gorm:"column:materials_15_probability;type:float;default:0;not null;" json:"materials_15_probability"`
	Materials16Id          int     `elements:"name:materials_16_id;type:int;size:4;text:未知;" gorm:"column:materials_16_id;type:integer;default:0;not null;" json:"materials_16_id"`
	Materials16Probability float32 `elements:"name:materials_16_probability;type:float;size:4;text:未知;" gorm:"column:materials_16_probability;type:float;default:0;not null;" json:"materials_16_probability"`
	Num1                   int     `elements:"name:num1;type:int;size:4;text:未知;" gorm:"column:num1;type:integer;default:0;not null;" json:"num1"`
	Probability1           float32 `elements:"name:probability1;type:float;size:4;text:未知;" gorm:"column:probability1;type:float;default:0;not null;" json:"probability1"`
	Num2                   int     `elements:"name:num2;type:int;size:4;text:未知;" gorm:"column:num2;type:integer;default:0;not null;" json:"num2"`
	Probability2           float32 `elements:"name:probability2;type:float;size:4;text:未知;" gorm:"column:probability2;type:float;default:0;not null;" json:"probability2"`
	TaskIn                 int     `elements:"name:task_in;type:int;size:4;text:未知;" gorm:"column:task_in;type:integer;default:0;not null;" json:"task_in"`
	TaskOut                int     `elements:"name:task_out;type:int;size:4;text:未知;" gorm:"column:task_out;type:integer;default:0;not null;" json:"task_out"`
	Uninterruptable        int     `elements:"name:uninterruptable;type:int;size:4;text:未知;" gorm:"column:uninterruptable;type:integer;default:0;not null;" json:"uninterruptable"`
	Npcgen1IdMonster       int     `elements:"name:npcgen_1_id_monster;type:int;size:4;text:未知;" gorm:"column:npcgen_1_id_monster;type:integer;default:0;not null;" json:"npcgen_1_id_monster"`
	Npcgen1Num             int     `elements:"name:npcgen_1_num;type:int;size:4;text:未知;" gorm:"column:npcgen_1_num;type:integer;default:0;not null;" json:"npcgen_1_num"`
	Npcgen1Radius          float32 `elements:"name:npcgen_1_radius;type:float;size:4;text:未知;" gorm:"column:npcgen_1_radius;type:float;default:0;not null;" json:"npcgen_1_radius"`
	Npcgen2IdMonster       int     `elements:"name:npcgen_2_id_monster;type:int;size:4;text:未知;" gorm:"column:npcgen_2_id_monster;type:integer;default:0;not null;" json:"npcgen_2_id_monster"`
	Npcgen2Num             int     `elements:"name:npcgen_2_num;type:int;size:4;text:未知;" gorm:"column:npcgen_2_num;type:integer;default:0;not null;" json:"npcgen_2_num"`
	Npcgen2Radius          float32 `elements:"name:npcgen_2_radius;type:float;size:4;text:未知;" gorm:"column:npcgen_2_radius;type:float;default:0;not null;" json:"npcgen_2_radius"`
	Npcgen3IdMonster       int     `elements:"name:npcgen_3_id_monster;type:int;size:4;text:未知;" gorm:"column:npcgen_3_id_monster;type:integer;default:0;not null;" json:"npcgen_3_id_monster"`
	Npcgen3Num             int     `elements:"name:npcgen_3_num;type:int;size:4;text:未知;" gorm:"column:npcgen_3_num;type:integer;default:0;not null;" json:"npcgen_3_num"`
	Npcgen3Radius          float32 `elements:"name:npcgen_3_radius;type:float;size:4;text:未知;" gorm:"column:npcgen_3_radius;type:float;default:0;not null;" json:"npcgen_3_radius"`
	Npcgen4IdMonster       int     `elements:"name:npcgen_4_id_monster;type:int;size:4;text:未知;" gorm:"column:npcgen_4_id_monster;type:integer;default:0;not null;" json:"npcgen_4_id_monster"`
	Npcgen4Num             int     `elements:"name:npcgen_4_num;type:int;size:4;text:未知;" gorm:"column:npcgen_4_num;type:integer;default:0;not null;" json:"npcgen_4_num"`
	Npcgen4Radius          float32 `elements:"name:npcgen_4_radius;type:float;size:4;text:未知;" gorm:"column:npcgen_4_radius;type:float;default:0;not null;" json:"npcgen_4_radius"`
	Aggros1MonsterFaction  int     `elements:"name:aggros_1_monster_faction;type:int;size:4;text:未知;" gorm:"column:aggros_1_monster_faction;type:integer;default:0;not null;" json:"aggros_1_monster_faction"`
	Aggros1Radius          float32 `elements:"name:aggros_1_radius;type:float;size:4;text:未知;" gorm:"column:aggros_1_radius;type:float;default:0;not null;" json:"aggros_1_radius"`
	Aggros1Num             int     `elements:"name:aggros_1_num;type:int;size:4;text:未知;" gorm:"column:aggros_1_num;type:integer;default:0;not null;" json:"aggros_1_num"`
	Permenent              int     `elements:"name:permenent;type:int;size:4;text:未知;" gorm:"column:permenent;type:integer;default:0;not null;" json:"permenent"`
}

func (MineEssence) TableName() string {
	return "el_mineessence"
}

func (e *MineEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *MineEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
