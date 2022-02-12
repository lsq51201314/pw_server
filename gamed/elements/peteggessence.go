package elements

import (
	"pw_server/utils/structEx"
)

type PetEggEssence struct {
	ID              int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name            string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter      string `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon        string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	IdPet           int    `elements:"name:id_pet;type:int;size:4;text:未知;" gorm:"column:id_pet;type:integer;default:0;not null;" json:"id_pet"`
	MoneyHatched    int    `elements:"name:money_hatched;type:int;size:4;text:未知;" gorm:"column:money_hatched;type:integer;default:0;not null;" json:"money_hatched"`
	MoneyRestored   int    `elements:"name:money_restored;type:int;size:4;text:未知;" gorm:"column:money_restored;type:integer;default:0;not null;" json:"money_restored"`
	HonorPoint      int    `elements:"name:honor_point;type:int;size:4;text:未知;" gorm:"column:honor_point;type:integer;default:0;not null;" json:"honor_point"`
	Level           int    `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	Exp             int    `elements:"name:exp;type:int;size:4;text:未知;" gorm:"column:exp;type:integer;default:0;not null;" json:"exp"`
	SkillPoint      int    `elements:"name:skill_point;type:int;size:4;text:未知;" gorm:"column:skill_point;type:integer;default:0;not null;" json:"skill_point"`
	Skills1IdSkill  int    `elements:"name:skills_1_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_1_id_skill;type:integer;default:0;not null;" json:"skills_1_id_skill"`
	Skills1Level    int    `elements:"name:skills_1_level;type:int;size:4;text:未知;" gorm:"column:skills_1_level;type:integer;default:0;not null;" json:"skills_1_level"`
	Skills2IdSkill  int    `elements:"name:skills_2_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_2_id_skill;type:integer;default:0;not null;" json:"skills_2_id_skill"`
	Skills2Level    int    `elements:"name:skills_2_level;type:int;size:4;text:未知;" gorm:"column:skills_2_level;type:integer;default:0;not null;" json:"skills_2_level"`
	Skills3IdSkill  int    `elements:"name:skills_3_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_3_id_skill;type:integer;default:0;not null;" json:"skills_3_id_skill"`
	Skills3Level    int    `elements:"name:skills_3_level;type:int;size:4;text:未知;" gorm:"column:skills_3_level;type:integer;default:0;not null;" json:"skills_3_level"`
	Skills4IdSkill  int    `elements:"name:skills_4_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_4_id_skill;type:integer;default:0;not null;" json:"skills_4_id_skill"`
	Skills4Level    int    `elements:"name:skills_4_level;type:int;size:4;text:未知;" gorm:"column:skills_4_level;type:integer;default:0;not null;" json:"skills_4_level"`
	Skills5IdSkill  int    `elements:"name:skills_5_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_5_id_skill;type:integer;default:0;not null;" json:"skills_5_id_skill"`
	Skills5Level    int    `elements:"name:skills_5_level;type:int;size:4;text:未知;" gorm:"column:skills_5_level;type:integer;default:0;not null;" json:"skills_5_level"`
	Skills6IdSkill  int    `elements:"name:skills_6_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_6_id_skill;type:integer;default:0;not null;" json:"skills_6_id_skill"`
	Skills6Level    int    `elements:"name:skills_6_level;type:int;size:4;text:未知;" gorm:"column:skills_6_level;type:integer;default:0;not null;" json:"skills_6_level"`
	Skills7IdSkill  int    `elements:"name:skills_7_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_7_id_skill;type:integer;default:0;not null;" json:"skills_7_id_skill"`
	Skills7Level    int    `elements:"name:skills_7_level;type:int;size:4;text:未知;" gorm:"column:skills_7_level;type:integer;default:0;not null;" json:"skills_7_level"`
	Skills8IdSkill  int    `elements:"name:skills_8_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_8_id_skill;type:integer;default:0;not null;" json:"skills_8_id_skill"`
	Skills8Level    int    `elements:"name:skills_8_level;type:int;size:4;text:未知;" gorm:"column:skills_8_level;type:integer;default:0;not null;" json:"skills_8_level"`
	Skills9IdSkill  int    `elements:"name:skills_9_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_9_id_skill;type:integer;default:0;not null;" json:"skills_9_id_skill"`
	Skills9Level    int    `elements:"name:skills_9_level;type:int;size:4;text:未知;" gorm:"column:skills_9_level;type:integer;default:0;not null;" json:"skills_9_level"`
	Skills10IdSkill int    `elements:"name:skills_10_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_10_id_skill;type:integer;default:0;not null;" json:"skills_10_id_skill"`
	Skills10Level   int    `elements:"name:skills_10_level;type:int;size:4;text:未知;" gorm:"column:skills_10_level;type:integer;default:0;not null;" json:"skills_10_level"`
	Skills11IdSkill int    `elements:"name:skills_11_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_11_id_skill;type:integer;default:0;not null;" json:"skills_11_id_skill"`
	Skills11Level   int    `elements:"name:skills_11_level;type:int;size:4;text:未知;" gorm:"column:skills_11_level;type:integer;default:0;not null;" json:"skills_11_level"`
	Skills12IdSkill int    `elements:"name:skills_12_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_12_id_skill;type:integer;default:0;not null;" json:"skills_12_id_skill"`
	Skills12Level   int    `elements:"name:skills_12_level;type:int;size:4;text:未知;" gorm:"column:skills_12_level;type:integer;default:0;not null;" json:"skills_12_level"`
	Skills13IdSkill int    `elements:"name:skills_13_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_13_id_skill;type:integer;default:0;not null;" json:"skills_13_id_skill"`
	Skills13Level   int    `elements:"name:skills_13_level;type:int;size:4;text:未知;" gorm:"column:skills_13_level;type:integer;default:0;not null;" json:"skills_13_level"`
	Skills14IdSkill int    `elements:"name:skills_14_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_14_id_skill;type:integer;default:0;not null;" json:"skills_14_id_skill"`
	Skills14Level   int    `elements:"name:skills_14_level;type:int;size:4;text:未知;" gorm:"column:skills_14_level;type:integer;default:0;not null;" json:"skills_14_level"`
	Skills15IdSkill int    `elements:"name:skills_15_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_15_id_skill;type:integer;default:0;not null;" json:"skills_15_id_skill"`
	Skills15Level   int    `elements:"name:skills_15_level;type:int;size:4;text:未知;" gorm:"column:skills_15_level;type:integer;default:0;not null;" json:"skills_15_level"`
	Skills16IdSkill int    `elements:"name:skills_16_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_16_id_skill;type:integer;default:0;not null;" json:"skills_16_id_skill"`
	Skills16Level   int    `elements:"name:skills_16_level;type:int;size:4;text:未知;" gorm:"column:skills_16_level;type:integer;default:0;not null;" json:"skills_16_level"`
	Skills17IdSkill int    `elements:"name:skills_17_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_17_id_skill;type:integer;default:0;not null;" json:"skills_17_id_skill"`
	Skills17Level   int    `elements:"name:skills_17_level;type:int;size:4;text:未知;" gorm:"column:skills_17_level;type:integer;default:0;not null;" json:"skills_17_level"`
	Skills18IdSkill int    `elements:"name:skills_18_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_18_id_skill;type:integer;default:0;not null;" json:"skills_18_id_skill"`
	Skills18Level   int    `elements:"name:skills_18_level;type:int;size:4;text:未知;" gorm:"column:skills_18_level;type:integer;default:0;not null;" json:"skills_18_level"`
	Skills19IdSkill int    `elements:"name:skills_19_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_19_id_skill;type:integer;default:0;not null;" json:"skills_19_id_skill"`
	Skills19Level   int    `elements:"name:skills_19_level;type:int;size:4;text:未知;" gorm:"column:skills_19_level;type:integer;default:0;not null;" json:"skills_19_level"`
	Skills20IdSkill int    `elements:"name:skills_20_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_20_id_skill;type:integer;default:0;not null;" json:"skills_20_id_skill"`
	Skills20Level   int    `elements:"name:skills_20_level;type:int;size:4;text:未知;" gorm:"column:skills_20_level;type:integer;default:0;not null;" json:"skills_20_level"`
	Skills21IdSkill int    `elements:"name:skills_21_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_21_id_skill;type:integer;default:0;not null;" json:"skills_21_id_skill"`
	Skills21Level   int    `elements:"name:skills_21_level;type:int;size:4;text:未知;" gorm:"column:skills_21_level;type:integer;default:0;not null;" json:"skills_21_level"`
	Skills22IdSkill int    `elements:"name:skills_22_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_22_id_skill;type:integer;default:0;not null;" json:"skills_22_id_skill"`
	Skills22Level   int    `elements:"name:skills_22_level;type:int;size:4;text:未知;" gorm:"column:skills_22_level;type:integer;default:0;not null;" json:"skills_22_level"`
	Skills23IdSkill int    `elements:"name:skills_23_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_23_id_skill;type:integer;default:0;not null;" json:"skills_23_id_skill"`
	Skills23Level   int    `elements:"name:skills_23_level;type:int;size:4;text:未知;" gorm:"column:skills_23_level;type:integer;default:0;not null;" json:"skills_23_level"`
	Skills24IdSkill int    `elements:"name:skills_24_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_24_id_skill;type:integer;default:0;not null;" json:"skills_24_id_skill"`
	Skills24Level   int    `elements:"name:skills_24_level;type:int;size:4;text:未知;" gorm:"column:skills_24_level;type:integer;default:0;not null;" json:"skills_24_level"`
	Skills25IdSkill int    `elements:"name:skills_25_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_25_id_skill;type:integer;default:0;not null;" json:"skills_25_id_skill"`
	Skills25Level   int    `elements:"name:skills_25_level;type:int;size:4;text:未知;" gorm:"column:skills_25_level;type:integer;default:0;not null;" json:"skills_25_level"`
	Skills26IdSkill int    `elements:"name:skills_26_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_26_id_skill;type:integer;default:0;not null;" json:"skills_26_id_skill"`
	Skills26Level   int    `elements:"name:skills_26_level;type:int;size:4;text:未知;" gorm:"column:skills_26_level;type:integer;default:0;not null;" json:"skills_26_level"`
	Skills27IdSkill int    `elements:"name:skills_27_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_27_id_skill;type:integer;default:0;not null;" json:"skills_27_id_skill"`
	Skills27Level   int    `elements:"name:skills_27_level;type:int;size:4;text:未知;" gorm:"column:skills_27_level;type:integer;default:0;not null;" json:"skills_27_level"`
	Skills28IdSkill int    `elements:"name:skills_28_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_28_id_skill;type:integer;default:0;not null;" json:"skills_28_id_skill"`
	Skills28Level   int    `elements:"name:skills_28_level;type:int;size:4;text:未知;" gorm:"column:skills_28_level;type:integer;default:0;not null;" json:"skills_28_level"`
	Skills29IdSkill int    `elements:"name:skills_29_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_29_id_skill;type:integer;default:0;not null;" json:"skills_29_id_skill"`
	Skills29Level   int    `elements:"name:skills_29_level;type:int;size:4;text:未知;" gorm:"column:skills_29_level;type:integer;default:0;not null;" json:"skills_29_level"`
	Skills30IdSkill int    `elements:"name:skills_30_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_30_id_skill;type:integer;default:0;not null;" json:"skills_30_id_skill"`
	Skills30Level   int    `elements:"name:skills_30_level;type:int;size:4;text:未知;" gorm:"column:skills_30_level;type:integer;default:0;not null;" json:"skills_30_level"`
	Skills31IdSkill int    `elements:"name:skills_31_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_31_id_skill;type:integer;default:0;not null;" json:"skills_31_id_skill"`
	Skills31Level   int    `elements:"name:skills_31_level;type:int;size:4;text:未知;" gorm:"column:skills_31_level;type:integer;default:0;not null;" json:"skills_31_level"`
	Skills32IdSkill int    `elements:"name:skills_32_id_skill;type:int;size:4;text:未知;" gorm:"column:skills_32_id_skill;type:integer;default:0;not null;" json:"skills_32_id_skill"`
	Skills32Level   int    `elements:"name:skills_32_level;type:int;size:4;text:未知;" gorm:"column:skills_32_level;type:integer;default:0;not null;" json:"skills_32_level"`
	Price           int    `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice       int    `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	PileNumMax      int    `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid         int    `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType        int    `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (PetEggEssence) TableName() string {
	return "el_peteggessence"
}

func (e *PetEggEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *PetEggEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
