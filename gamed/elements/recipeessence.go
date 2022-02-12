package elements

import (
	"pw_server/utils/structEx"
)

type RecipeEssence struct {
	ID                  int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdMajorType         int     `elements:"name:id_major_type;type:int;size:4;text:未知;" gorm:"column:id_major_type;type:integer;default:0;not null;" json:"id_major_type"`
	IdSubType           int     `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name                string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	RecipeLevel         int     `elements:"name:recipe_level;type:int;size:4;text:未知;" gorm:"column:recipe_level;type:integer;default:0;not null;" json:"recipe_level"`
	IdSkill             int     `elements:"name:id_skill;type:int;size:4;text:未知;" gorm:"column:id_skill;type:integer;default:0;not null;" json:"id_skill"`
	SkillLevel          int     `elements:"name:skill_level;type:int;size:4;text:未知;" gorm:"column:skill_level;type:integer;default:0;not null;" json:"skill_level"`
	Targets1IdToMake    int     `elements:"name:targets_1_id_to_make;type:int;size:4;text:未知;" gorm:"column:targets_1_id_to_make;type:integer;default:0;not null;" json:"targets_1_id_to_make"`
	Targets1Probability float32 `elements:"name:targets_1_probability;type:float;size:4;text:未知;" gorm:"column:targets_1_probability;type:float;default:0;not null;" json:"targets_1_probability"`
	Targets2IdToMake    int     `elements:"name:targets_2_id_to_make;type:int;size:4;text:未知;" gorm:"column:targets_2_id_to_make;type:integer;default:0;not null;" json:"targets_2_id_to_make"`
	Targets2Probability float32 `elements:"name:targets_2_probability;type:float;size:4;text:未知;" gorm:"column:targets_2_probability;type:float;default:0;not null;" json:"targets_2_probability"`
	Targets3IdToMake    int     `elements:"name:targets_3_id_to_make;type:int;size:4;text:未知;" gorm:"column:targets_3_id_to_make;type:integer;default:0;not null;" json:"targets_3_id_to_make"`
	Targets3Probability float32 `elements:"name:targets_3_probability;type:float;size:4;text:未知;" gorm:"column:targets_3_probability;type:float;default:0;not null;" json:"targets_3_probability"`
	Targets4IdToMake    int     `elements:"name:targets_4_id_to_make;type:int;size:4;text:未知;" gorm:"column:targets_4_id_to_make;type:integer;default:0;not null;" json:"targets_4_id_to_make"`
	Targets4Probability float32 `elements:"name:targets_4_probability;type:float;size:4;text:未知;" gorm:"column:targets_4_probability;type:float;default:0;not null;" json:"targets_4_probability"`
	FailProbability     float32 `elements:"name:fail_probability;type:float;size:4;text:未知;" gorm:"column:fail_probability;type:float;default:0;not null;" json:"fail_probability"`
	NumToMake           int     `elements:"name:num_to_make;type:int;size:4;text:未知;" gorm:"column:num_to_make;type:integer;default:0;not null;" json:"num_to_make"`
	Price               int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	Duration            float32 `elements:"name:duration;type:float;size:4;text:未知;" gorm:"column:duration;type:float;default:0;not null;" json:"duration"`
	Exp                 int     `elements:"name:exp;type:int;size:4;text:未知;" gorm:"column:exp;type:integer;default:0;not null;" json:"exp"`
	Skillpoint          int     `elements:"name:skillpoint;type:int;size:4;text:未知;" gorm:"column:skillpoint;type:integer;default:0;not null;" json:"skillpoint"`
	Materials1Id        int     `elements:"name:materials_1_id;type:int;size:4;text:未知;" gorm:"column:materials_1_id;type:integer;default:0;not null;" json:"materials_1_id"`
	Materials1Num       int     `elements:"name:materials_1_num;type:int;size:4;text:未知;" gorm:"column:materials_1_num;type:integer;default:0;not null;" json:"materials_1_num"`
	Materials2Id        int     `elements:"name:materials_2_id;type:int;size:4;text:未知;" gorm:"column:materials_2_id;type:integer;default:0;not null;" json:"materials_2_id"`
	Materials2Num       int     `elements:"name:materials_2_num;type:int;size:4;text:未知;" gorm:"column:materials_2_num;type:integer;default:0;not null;" json:"materials_2_num"`
	Materials3Id        int     `elements:"name:materials_3_id;type:int;size:4;text:未知;" gorm:"column:materials_3_id;type:integer;default:0;not null;" json:"materials_3_id"`
	Materials3Num       int     `elements:"name:materials_3_num;type:int;size:4;text:未知;" gorm:"column:materials_3_num;type:integer;default:0;not null;" json:"materials_3_num"`
	Materials4Id        int     `elements:"name:materials_4_id;type:int;size:4;text:未知;" gorm:"column:materials_4_id;type:integer;default:0;not null;" json:"materials_4_id"`
	Materials4Num       int     `elements:"name:materials_4_num;type:int;size:4;text:未知;" gorm:"column:materials_4_num;type:integer;default:0;not null;" json:"materials_4_num"`
	Materials5Id        int     `elements:"name:materials_5_id;type:int;size:4;text:未知;" gorm:"column:materials_5_id;type:integer;default:0;not null;" json:"materials_5_id"`
	Materials5Num       int     `elements:"name:materials_5_num;type:int;size:4;text:未知;" gorm:"column:materials_5_num;type:integer;default:0;not null;" json:"materials_5_num"`
	Materials6Id        int     `elements:"name:materials_6_id;type:int;size:4;text:未知;" gorm:"column:materials_6_id;type:integer;default:0;not null;" json:"materials_6_id"`
	Materials6Num       int     `elements:"name:materials_6_num;type:int;size:4;text:未知;" gorm:"column:materials_6_num;type:integer;default:0;not null;" json:"materials_6_num"`
	Materials7Id        int     `elements:"name:materials_7_id;type:int;size:4;text:未知;" gorm:"column:materials_7_id;type:integer;default:0;not null;" json:"materials_7_id"`
	Materials7Num       int     `elements:"name:materials_7_num;type:int;size:4;text:未知;" gorm:"column:materials_7_num;type:integer;default:0;not null;" json:"materials_7_num"`
	Materials8Id        int     `elements:"name:materials_8_id;type:int;size:4;text:未知;" gorm:"column:materials_8_id;type:integer;default:0;not null;" json:"materials_8_id"`
	Materials8Num       int     `elements:"name:materials_8_num;type:int;size:4;text:未知;" gorm:"column:materials_8_num;type:integer;default:0;not null;" json:"materials_8_num"`
	Materials9Id        int     `elements:"name:materials_9_id;type:int;size:4;text:未知;" gorm:"column:materials_9_id;type:integer;default:0;not null;" json:"materials_9_id"`
	Materials9Num       int     `elements:"name:materials_9_num;type:int;size:4;text:未知;" gorm:"column:materials_9_num;type:integer;default:0;not null;" json:"materials_9_num"`
	Materials10Id       int     `elements:"name:materials_10_id;type:int;size:4;text:未知;" gorm:"column:materials_10_id;type:integer;default:0;not null;" json:"materials_10_id"`
	Materials10Num      int     `elements:"name:materials_10_num;type:int;size:4;text:未知;" gorm:"column:materials_10_num;type:integer;default:0;not null;" json:"materials_10_num"`
	Materials11Id       int     `elements:"name:materials_11_id;type:int;size:4;text:未知;" gorm:"column:materials_11_id;type:integer;default:0;not null;" json:"materials_11_id"`
	Materials11Num      int     `elements:"name:materials_11_num;type:int;size:4;text:未知;" gorm:"column:materials_11_num;type:integer;default:0;not null;" json:"materials_11_num"`
	Materials12Id       int     `elements:"name:materials_12_id;type:int;size:4;text:未知;" gorm:"column:materials_12_id;type:integer;default:0;not null;" json:"materials_12_id"`
	Materials12Num      int     `elements:"name:materials_12_num;type:int;size:4;text:未知;" gorm:"column:materials_12_num;type:integer;default:0;not null;" json:"materials_12_num"`
	Materials13Id       int     `elements:"name:materials_13_id;type:int;size:4;text:未知;" gorm:"column:materials_13_id;type:integer;default:0;not null;" json:"materials_13_id"`
	Materials13Num      int     `elements:"name:materials_13_num;type:int;size:4;text:未知;" gorm:"column:materials_13_num;type:integer;default:0;not null;" json:"materials_13_num"`
	Materials14Id       int     `elements:"name:materials_14_id;type:int;size:4;text:未知;" gorm:"column:materials_14_id;type:integer;default:0;not null;" json:"materials_14_id"`
	Materials14Num      int     `elements:"name:materials_14_num;type:int;size:4;text:未知;" gorm:"column:materials_14_num;type:integer;default:0;not null;" json:"materials_14_num"`
	Materials15Id       int     `elements:"name:materials_15_id;type:int;size:4;text:未知;" gorm:"column:materials_15_id;type:integer;default:0;not null;" json:"materials_15_id"`
	Materials15Num      int     `elements:"name:materials_15_num;type:int;size:4;text:未知;" gorm:"column:materials_15_num;type:integer;default:0;not null;" json:"materials_15_num"`
	Materials16Id       int     `elements:"name:materials_16_id;type:int;size:4;text:未知;" gorm:"column:materials_16_id;type:integer;default:0;not null;" json:"materials_16_id"`
	Materials16Num      int     `elements:"name:materials_16_num;type:int;size:4;text:未知;" gorm:"column:materials_16_num;type:integer;default:0;not null;" json:"materials_16_num"`
	Materials17Id       int     `elements:"name:materials_17_id;type:int;size:4;text:未知;" gorm:"column:materials_17_id;type:integer;default:0;not null;" json:"materials_17_id"`
	Materials17Num      int     `elements:"name:materials_17_num;type:int;size:4;text:未知;" gorm:"column:materials_17_num;type:integer;default:0;not null;" json:"materials_17_num"`
	Materials18Id       int     `elements:"name:materials_18_id;type:int;size:4;text:未知;" gorm:"column:materials_18_id;type:integer;default:0;not null;" json:"materials_18_id"`
	Materials18Num      int     `elements:"name:materials_18_num;type:int;size:4;text:未知;" gorm:"column:materials_18_num;type:integer;default:0;not null;" json:"materials_18_num"`
	Materials19Id       int     `elements:"name:materials_19_id;type:int;size:4;text:未知;" gorm:"column:materials_19_id;type:integer;default:0;not null;" json:"materials_19_id"`
	Materials19Num      int     `elements:"name:materials_19_num;type:int;size:4;text:未知;" gorm:"column:materials_19_num;type:integer;default:0;not null;" json:"materials_19_num"`
	Materials20Id       int     `elements:"name:materials_20_id;type:int;size:4;text:未知;" gorm:"column:materials_20_id;type:integer;default:0;not null;" json:"materials_20_id"`
	Materials20Num      int     `elements:"name:materials_20_num;type:int;size:4;text:未知;" gorm:"column:materials_20_num;type:integer;default:0;not null;" json:"materials_20_num"`
	Materials21Id       int     `elements:"name:materials_21_id;type:int;size:4;text:未知;" gorm:"column:materials_21_id;type:integer;default:0;not null;" json:"materials_21_id"`
	Materials21Num      int     `elements:"name:materials_21_num;type:int;size:4;text:未知;" gorm:"column:materials_21_num;type:integer;default:0;not null;" json:"materials_21_num"`
	Materials22Id       int     `elements:"name:materials_22_id;type:int;size:4;text:未知;" gorm:"column:materials_22_id;type:integer;default:0;not null;" json:"materials_22_id"`
	Materials22Num      int     `elements:"name:materials_22_num;type:int;size:4;text:未知;" gorm:"column:materials_22_num;type:integer;default:0;not null;" json:"materials_22_num"`
	Materials23Id       int     `elements:"name:materials_23_id;type:int;size:4;text:未知;" gorm:"column:materials_23_id;type:integer;default:0;not null;" json:"materials_23_id"`
	Materials23Num      int     `elements:"name:materials_23_num;type:int;size:4;text:未知;" gorm:"column:materials_23_num;type:integer;default:0;not null;" json:"materials_23_num"`
	Materials24Id       int     `elements:"name:materials_24_id;type:int;size:4;text:未知;" gorm:"column:materials_24_id;type:integer;default:0;not null;" json:"materials_24_id"`
	Materials24Num      int     `elements:"name:materials_24_num;type:int;size:4;text:未知;" gorm:"column:materials_24_num;type:integer;default:0;not null;" json:"materials_24_num"`
	Materials25Id       int     `elements:"name:materials_25_id;type:int;size:4;text:未知;" gorm:"column:materials_25_id;type:integer;default:0;not null;" json:"materials_25_id"`
	Materials25Num      int     `elements:"name:materials_25_num;type:int;size:4;text:未知;" gorm:"column:materials_25_num;type:integer;default:0;not null;" json:"materials_25_num"`
	Materials26Id       int     `elements:"name:materials_26_id;type:int;size:4;text:未知;" gorm:"column:materials_26_id;type:integer;default:0;not null;" json:"materials_26_id"`
	Materials26Num      int     `elements:"name:materials_26_num;type:int;size:4;text:未知;" gorm:"column:materials_26_num;type:integer;default:0;not null;" json:"materials_26_num"`
	Materials27Id       int     `elements:"name:materials_27_id;type:int;size:4;text:未知;" gorm:"column:materials_27_id;type:integer;default:0;not null;" json:"materials_27_id"`
	Materials27Num      int     `elements:"name:materials_27_num;type:int;size:4;text:未知;" gorm:"column:materials_27_num;type:integer;default:0;not null;" json:"materials_27_num"`
	Materials28Id       int     `elements:"name:materials_28_id;type:int;size:4;text:未知;" gorm:"column:materials_28_id;type:integer;default:0;not null;" json:"materials_28_id"`
	Materials28Num      int     `elements:"name:materials_28_num;type:int;size:4;text:未知;" gorm:"column:materials_28_num;type:integer;default:0;not null;" json:"materials_28_num"`
	Materials29Id       int     `elements:"name:materials_29_id;type:int;size:4;text:未知;" gorm:"column:materials_29_id;type:integer;default:0;not null;" json:"materials_29_id"`
	Materials29Num      int     `elements:"name:materials_29_num;type:int;size:4;text:未知;" gorm:"column:materials_29_num;type:integer;default:0;not null;" json:"materials_29_num"`
	Materials30Id       int     `elements:"name:materials_30_id;type:int;size:4;text:未知;" gorm:"column:materials_30_id;type:integer;default:0;not null;" json:"materials_30_id"`
	Materials30Num      int     `elements:"name:materials_30_num;type:int;size:4;text:未知;" gorm:"column:materials_30_num;type:integer;default:0;not null;" json:"materials_30_num"`
	Materials31Id       int     `elements:"name:materials_31_id;type:int;size:4;text:未知;" gorm:"column:materials_31_id;type:integer;default:0;not null;" json:"materials_31_id"`
	Materials31Num      int     `elements:"name:materials_31_num;type:int;size:4;text:未知;" gorm:"column:materials_31_num;type:integer;default:0;not null;" json:"materials_31_num"`
	Materials32Id       int     `elements:"name:materials_32_id;type:int;size:4;text:未知;" gorm:"column:materials_32_id;type:integer;default:0;not null;" json:"materials_32_id"`
	Materials32Num      int     `elements:"name:materials_32_num;type:int;size:4;text:未知;" gorm:"column:materials_32_num;type:integer;default:0;not null;" json:"materials_32_num"`
}

func (RecipeEssence) TableName() string {
	return "el_recipeessence"
}

func (e *RecipeEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *RecipeEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
