package elements

import (
	"pw_server/utils/structEx"
)

type DecorationEssence struct {
	ID                       int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdMajorType              int     `elements:"name:id_major_type;type:int;size:4;text:未知;" gorm:"column:id_major_type;type:integer;default:0;not null;" json:"id_major_type"`
	IdSubType                int     `elements:"name:id_sub_type;type:int;size:4;text:未知;" gorm:"column:id_sub_type;type:integer;default:0;not null;" json:"id_sub_type"`
	Name                     string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel                string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileMatter               string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon                 string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	Level                    int     `elements:"name:level;type:int;size:4;text:未知;" gorm:"column:level;type:integer;default:0;not null;" json:"level"`
	RequireStrength          int     `elements:"name:require_strength;type:int;size:4;text:未知;" gorm:"column:require_strength;type:integer;default:0;not null;" json:"require_strength"`
	RequireAgility           int     `elements:"name:require_agility;type:int;size:4;text:未知;" gorm:"column:require_agility;type:integer;default:0;not null;" json:"require_agility"`
	RequireEnergy            int     `elements:"name:require_energy;type:int;size:4;text:未知;" gorm:"column:require_energy;type:integer;default:0;not null;" json:"require_energy"`
	RequireTili              int     `elements:"name:require_tili;type:int;size:4;text:未知;" gorm:"column:require_tili;type:integer;default:0;not null;" json:"require_tili"`
	CharacterComboId         int     `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	RequireLevel             int     `elements:"name:require_level;type:int;size:4;text:未知;" gorm:"column:require_level;type:integer;default:0;not null;" json:"require_level"`
	FixedProps               int     `elements:"name:fixed_props;type:int;size:4;text:未知;" gorm:"column:fixed_props;type:integer;default:0;not null;" json:"fixed_props"`
	DamageLow                int     `elements:"name:damage_low;type:int;size:4;text:未知;" gorm:"column:damage_low;type:integer;default:0;not null;" json:"damage_low"`
	DamageHigh               int     `elements:"name:damage_high;type:int;size:4;text:未知;" gorm:"column:damage_high;type:integer;default:0;not null;" json:"damage_high"`
	MagicDamageLow           int     `elements:"name:magic_damage_low;type:int;size:4;text:未知;" gorm:"column:magic_damage_low;type:integer;default:0;not null;" json:"magic_damage_low"`
	MagicDamageHigh          int     `elements:"name:magic_damage_high;type:int;size:4;text:未知;" gorm:"column:magic_damage_high;type:integer;default:0;not null;" json:"magic_damage_high"`
	DefenceLow               int     `elements:"name:defence_low;type:int;size:4;text:未知;" gorm:"column:defence_low;type:integer;default:0;not null;" json:"defence_low"`
	DefenceHigh              int     `elements:"name:defence_high;type:int;size:4;text:未知;" gorm:"column:defence_high;type:integer;default:0;not null;" json:"defence_high"`
	MagicDefences1Low        int     `elements:"name:magic_defences_1_low;type:int;size:4;text:未知;" gorm:"column:magic_defences_1_low;type:integer;default:0;not null;" json:"magic_defences_1_low"`
	MagicDefences1High       int     `elements:"name:magic_defences_1_high;type:int;size:4;text:未知;" gorm:"column:magic_defences_1_high;type:integer;default:0;not null;" json:"magic_defences_1_high"`
	MagicDefences2Low        int     `elements:"name:magic_defences_2_low;type:int;size:4;text:未知;" gorm:"column:magic_defences_2_low;type:integer;default:0;not null;" json:"magic_defences_2_low"`
	MagicDefences2High       int     `elements:"name:magic_defences_2_high;type:int;size:4;text:未知;" gorm:"column:magic_defences_2_high;type:integer;default:0;not null;" json:"magic_defences_2_high"`
	MagicDefences3Low        int     `elements:"name:magic_defences_3_low;type:int;size:4;text:未知;" gorm:"column:magic_defences_3_low;type:integer;default:0;not null;" json:"magic_defences_3_low"`
	MagicDefences3High       int     `elements:"name:magic_defences_3_high;type:int;size:4;text:未知;" gorm:"column:magic_defences_3_high;type:integer;default:0;not null;" json:"magic_defences_3_high"`
	MagicDefences4Low        int     `elements:"name:magic_defences_4_low;type:int;size:4;text:未知;" gorm:"column:magic_defences_4_low;type:integer;default:0;not null;" json:"magic_defences_4_low"`
	MagicDefences4High       int     `elements:"name:magic_defences_4_high;type:int;size:4;text:未知;" gorm:"column:magic_defences_4_high;type:integer;default:0;not null;" json:"magic_defences_4_high"`
	MagicDefences5Low        int     `elements:"name:magic_defences_5_low;type:int;size:4;text:未知;" gorm:"column:magic_defences_5_low;type:integer;default:0;not null;" json:"magic_defences_5_low"`
	MagicDefences5High       int     `elements:"name:magic_defences_5_high;type:int;size:4;text:未知;" gorm:"column:magic_defences_5_high;type:integer;default:0;not null;" json:"magic_defences_5_high"`
	ArmorEnhanceLow          int     `elements:"name:armor_enhance_low;type:int;size:4;text:未知;" gorm:"column:armor_enhance_low;type:integer;default:0;not null;" json:"armor_enhance_low"`
	ArmorEnhanceHigh         int     `elements:"name:armor_enhance_high;type:int;size:4;text:未知;" gorm:"column:armor_enhance_high;type:integer;default:0;not null;" json:"armor_enhance_high"`
	DurabilityMin            int     `elements:"name:durability_min;type:int;size:4;text:未知;" gorm:"column:durability_min;type:integer;default:0;not null;" json:"durability_min"`
	DurabilityMax            int     `elements:"name:durability_max;type:int;size:4;text:未知;" gorm:"column:durability_max;type:integer;default:0;not null;" json:"durability_max"`
	LevelupAddon             int     `elements:"name:levelup_addon;type:int;size:4;text:未知;" gorm:"column:levelup_addon;type:integer;default:0;not null;" json:"levelup_addon"`
	MaterialNeed             int     `elements:"name:material_need;type:int;size:4;text:未知;" gorm:"column:material_need;type:integer;default:0;not null;" json:"material_need"`
	Price                    int     `elements:"name:price;type:int;size:4;text:未知;" gorm:"column:price;type:integer;default:0;not null;" json:"price"`
	ShopPrice                int     `elements:"name:shop_price;type:int;size:4;text:未知;" gorm:"column:shop_price;type:integer;default:0;not null;" json:"shop_price"`
	Repairfee                int     `elements:"name:repairfee;type:int;size:4;text:未知;" gorm:"column:repairfee;type:integer;default:0;not null;" json:"repairfee"`
	ProbabilityAddonNum0     float32 `elements:"name:probability_addon_num0;type:float;size:4;text:未知;" gorm:"column:probability_addon_num0;type:float;default:0;not null;" json:"probability_addon_num0"`
	ProbabilityAddonNum1     float32 `elements:"name:probability_addon_num1;type:float;size:4;text:未知;" gorm:"column:probability_addon_num1;type:float;default:0;not null;" json:"probability_addon_num1"`
	ProbabilityAddonNum2     float32 `elements:"name:probability_addon_num2;type:float;size:4;text:未知;" gorm:"column:probability_addon_num2;type:float;default:0;not null;" json:"probability_addon_num2"`
	ProbabilityAddonNum3     float32 `elements:"name:probability_addon_num3;type:float;size:4;text:未知;" gorm:"column:probability_addon_num3;type:float;default:0;not null;" json:"probability_addon_num3"`
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
	Addons17IdAddon          int     `elements:"name:addons_17_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_17_id_addon;type:integer;default:0;not null;" json:"addons_17_id_addon"`
	Addons17ProbabilityAddon float32 `elements:"name:addons_17_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_17_probability_addon;type:float;default:0;not null;" json:"addons_17_probability_addon"`
	Addons18IdAddon          int     `elements:"name:addons_18_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_18_id_addon;type:integer;default:0;not null;" json:"addons_18_id_addon"`
	Addons18ProbabilityAddon float32 `elements:"name:addons_18_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_18_probability_addon;type:float;default:0;not null;" json:"addons_18_probability_addon"`
	Addons19IdAddon          int     `elements:"name:addons_19_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_19_id_addon;type:integer;default:0;not null;" json:"addons_19_id_addon"`
	Addons19ProbabilityAddon float32 `elements:"name:addons_19_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_19_probability_addon;type:float;default:0;not null;" json:"addons_19_probability_addon"`
	Addons20IdAddon          int     `elements:"name:addons_20_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_20_id_addon;type:integer;default:0;not null;" json:"addons_20_id_addon"`
	Addons20ProbabilityAddon float32 `elements:"name:addons_20_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_20_probability_addon;type:float;default:0;not null;" json:"addons_20_probability_addon"`
	Addons21IdAddon          int     `elements:"name:addons_21_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_21_id_addon;type:integer;default:0;not null;" json:"addons_21_id_addon"`
	Addons21ProbabilityAddon float32 `elements:"name:addons_21_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_21_probability_addon;type:float;default:0;not null;" json:"addons_21_probability_addon"`
	Addons22IdAddon          int     `elements:"name:addons_22_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_22_id_addon;type:integer;default:0;not null;" json:"addons_22_id_addon"`
	Addons22ProbabilityAddon float32 `elements:"name:addons_22_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_22_probability_addon;type:float;default:0;not null;" json:"addons_22_probability_addon"`
	Addons23IdAddon          int     `elements:"name:addons_23_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_23_id_addon;type:integer;default:0;not null;" json:"addons_23_id_addon"`
	Addons23ProbabilityAddon float32 `elements:"name:addons_23_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_23_probability_addon;type:float;default:0;not null;" json:"addons_23_probability_addon"`
	Addons24IdAddon          int     `elements:"name:addons_24_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_24_id_addon;type:integer;default:0;not null;" json:"addons_24_id_addon"`
	Addons24ProbabilityAddon float32 `elements:"name:addons_24_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_24_probability_addon;type:float;default:0;not null;" json:"addons_24_probability_addon"`
	Addons25IdAddon          int     `elements:"name:addons_25_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_25_id_addon;type:integer;default:0;not null;" json:"addons_25_id_addon"`
	Addons25ProbabilityAddon float32 `elements:"name:addons_25_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_25_probability_addon;type:float;default:0;not null;" json:"addons_25_probability_addon"`
	Addons26IdAddon          int     `elements:"name:addons_26_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_26_id_addon;type:integer;default:0;not null;" json:"addons_26_id_addon"`
	Addons26ProbabilityAddon float32 `elements:"name:addons_26_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_26_probability_addon;type:float;default:0;not null;" json:"addons_26_probability_addon"`
	Addons27IdAddon          int     `elements:"name:addons_27_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_27_id_addon;type:integer;default:0;not null;" json:"addons_27_id_addon"`
	Addons27ProbabilityAddon float32 `elements:"name:addons_27_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_27_probability_addon;type:float;default:0;not null;" json:"addons_27_probability_addon"`
	Addons28IdAddon          int     `elements:"name:addons_28_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_28_id_addon;type:integer;default:0;not null;" json:"addons_28_id_addon"`
	Addons28ProbabilityAddon float32 `elements:"name:addons_28_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_28_probability_addon;type:float;default:0;not null;" json:"addons_28_probability_addon"`
	Addons29IdAddon          int     `elements:"name:addons_29_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_29_id_addon;type:integer;default:0;not null;" json:"addons_29_id_addon"`
	Addons29ProbabilityAddon float32 `elements:"name:addons_29_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_29_probability_addon;type:float;default:0;not null;" json:"addons_29_probability_addon"`
	Addons30IdAddon          int     `elements:"name:addons_30_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_30_id_addon;type:integer;default:0;not null;" json:"addons_30_id_addon"`
	Addons30ProbabilityAddon float32 `elements:"name:addons_30_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_30_probability_addon;type:float;default:0;not null;" json:"addons_30_probability_addon"`
	Addons31IdAddon          int     `elements:"name:addons_31_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_31_id_addon;type:integer;default:0;not null;" json:"addons_31_id_addon"`
	Addons31ProbabilityAddon float32 `elements:"name:addons_31_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_31_probability_addon;type:float;default:0;not null;" json:"addons_31_probability_addon"`
	Addons32IdAddon          int     `elements:"name:addons_32_id_addon;type:int;size:4;text:未知;" gorm:"column:addons_32_id_addon;type:integer;default:0;not null;" json:"addons_32_id_addon"`
	Addons32ProbabilityAddon float32 `elements:"name:addons_32_probability_addon;type:float;size:4;text:未知;" gorm:"column:addons_32_probability_addon;type:float;default:0;not null;" json:"addons_32_probability_addon"`
	Rands1IdRand             int     `elements:"name:rands_1_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_1_id_rand;type:integer;default:0;not null;" json:"rands_1_id_rand"`
	Rands1ProbabilityRand    float32 `elements:"name:rands_1_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_1_probability_rand;type:float;default:0;not null;" json:"rands_1_probability_rand"`
	Rands2IdRand             int     `elements:"name:rands_2_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_2_id_rand;type:integer;default:0;not null;" json:"rands_2_id_rand"`
	Rands2ProbabilityRand    float32 `elements:"name:rands_2_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_2_probability_rand;type:float;default:0;not null;" json:"rands_2_probability_rand"`
	Rands3IdRand             int     `elements:"name:rands_3_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_3_id_rand;type:integer;default:0;not null;" json:"rands_3_id_rand"`
	Rands3ProbabilityRand    float32 `elements:"name:rands_3_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_3_probability_rand;type:float;default:0;not null;" json:"rands_3_probability_rand"`
	Rands4IdRand             int     `elements:"name:rands_4_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_4_id_rand;type:integer;default:0;not null;" json:"rands_4_id_rand"`
	Rands4ProbabilityRand    float32 `elements:"name:rands_4_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_4_probability_rand;type:float;default:0;not null;" json:"rands_4_probability_rand"`
	Rands5IdRand             int     `elements:"name:rands_5_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_5_id_rand;type:integer;default:0;not null;" json:"rands_5_id_rand"`
	Rands5ProbabilityRand    float32 `elements:"name:rands_5_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_5_probability_rand;type:float;default:0;not null;" json:"rands_5_probability_rand"`
	Rands6IdRand             int     `elements:"name:rands_6_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_6_id_rand;type:integer;default:0;not null;" json:"rands_6_id_rand"`
	Rands6ProbabilityRand    float32 `elements:"name:rands_6_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_6_probability_rand;type:float;default:0;not null;" json:"rands_6_probability_rand"`
	Rands7IdRand             int     `elements:"name:rands_7_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_7_id_rand;type:integer;default:0;not null;" json:"rands_7_id_rand"`
	Rands7ProbabilityRand    float32 `elements:"name:rands_7_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_7_probability_rand;type:float;default:0;not null;" json:"rands_7_probability_rand"`
	Rands8IdRand             int     `elements:"name:rands_8_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_8_id_rand;type:integer;default:0;not null;" json:"rands_8_id_rand"`
	Rands8ProbabilityRand    float32 `elements:"name:rands_8_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_8_probability_rand;type:float;default:0;not null;" json:"rands_8_probability_rand"`
	Rands9IdRand             int     `elements:"name:rands_9_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_9_id_rand;type:integer;default:0;not null;" json:"rands_9_id_rand"`
	Rands9ProbabilityRand    float32 `elements:"name:rands_9_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_9_probability_rand;type:float;default:0;not null;" json:"rands_9_probability_rand"`
	Rands10IdRand            int     `elements:"name:rands_10_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_10_id_rand;type:integer;default:0;not null;" json:"rands_10_id_rand"`
	Rands10ProbabilityRand   float32 `elements:"name:rands_10_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_10_probability_rand;type:float;default:0;not null;" json:"rands_10_probability_rand"`
	Rands11IdRand            int     `elements:"name:rands_11_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_11_id_rand;type:integer;default:0;not null;" json:"rands_11_id_rand"`
	Rands11ProbabilityRand   float32 `elements:"name:rands_11_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_11_probability_rand;type:float;default:0;not null;" json:"rands_11_probability_rand"`
	Rands12IdRand            int     `elements:"name:rands_12_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_12_id_rand;type:integer;default:0;not null;" json:"rands_12_id_rand"`
	Rands12ProbabilityRand   float32 `elements:"name:rands_12_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_12_probability_rand;type:float;default:0;not null;" json:"rands_12_probability_rand"`
	Rands13IdRand            int     `elements:"name:rands_13_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_13_id_rand;type:integer;default:0;not null;" json:"rands_13_id_rand"`
	Rands13ProbabilityRand   float32 `elements:"name:rands_13_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_13_probability_rand;type:float;default:0;not null;" json:"rands_13_probability_rand"`
	Rands14IdRand            int     `elements:"name:rands_14_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_14_id_rand;type:integer;default:0;not null;" json:"rands_14_id_rand"`
	Rands14ProbabilityRand   float32 `elements:"name:rands_14_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_14_probability_rand;type:float;default:0;not null;" json:"rands_14_probability_rand"`
	Rands15IdRand            int     `elements:"name:rands_15_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_15_id_rand;type:integer;default:0;not null;" json:"rands_15_id_rand"`
	Rands15ProbabilityRand   float32 `elements:"name:rands_15_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_15_probability_rand;type:float;default:0;not null;" json:"rands_15_probability_rand"`
	Rands16IdRand            int     `elements:"name:rands_16_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_16_id_rand;type:integer;default:0;not null;" json:"rands_16_id_rand"`
	Rands16ProbabilityRand   float32 `elements:"name:rands_16_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_16_probability_rand;type:float;default:0;not null;" json:"rands_16_probability_rand"`
	Rands17IdRand            int     `elements:"name:rands_17_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_17_id_rand;type:integer;default:0;not null;" json:"rands_17_id_rand"`
	Rands17ProbabilityRand   float32 `elements:"name:rands_17_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_17_probability_rand;type:float;default:0;not null;" json:"rands_17_probability_rand"`
	Rands18IdRand            int     `elements:"name:rands_18_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_18_id_rand;type:integer;default:0;not null;" json:"rands_18_id_rand"`
	Rands18ProbabilityRand   float32 `elements:"name:rands_18_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_18_probability_rand;type:float;default:0;not null;" json:"rands_18_probability_rand"`
	Rands19IdRand            int     `elements:"name:rands_19_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_19_id_rand;type:integer;default:0;not null;" json:"rands_19_id_rand"`
	Rands19ProbabilityRand   float32 `elements:"name:rands_19_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_19_probability_rand;type:float;default:0;not null;" json:"rands_19_probability_rand"`
	Rands20IdRand            int     `elements:"name:rands_20_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_20_id_rand;type:integer;default:0;not null;" json:"rands_20_id_rand"`
	Rands20ProbabilityRand   float32 `elements:"name:rands_20_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_20_probability_rand;type:float;default:0;not null;" json:"rands_20_probability_rand"`
	Rands21IdRand            int     `elements:"name:rands_21_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_21_id_rand;type:integer;default:0;not null;" json:"rands_21_id_rand"`
	Rands21ProbabilityRand   float32 `elements:"name:rands_21_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_21_probability_rand;type:float;default:0;not null;" json:"rands_21_probability_rand"`
	Rands22IdRand            int     `elements:"name:rands_22_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_22_id_rand;type:integer;default:0;not null;" json:"rands_22_id_rand"`
	Rands22ProbabilityRand   float32 `elements:"name:rands_22_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_22_probability_rand;type:float;default:0;not null;" json:"rands_22_probability_rand"`
	Rands23IdRand            int     `elements:"name:rands_23_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_23_id_rand;type:integer;default:0;not null;" json:"rands_23_id_rand"`
	Rands23ProbabilityRand   float32 `elements:"name:rands_23_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_23_probability_rand;type:float;default:0;not null;" json:"rands_23_probability_rand"`
	Rands24IdRand            int     `elements:"name:rands_24_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_24_id_rand;type:integer;default:0;not null;" json:"rands_24_id_rand"`
	Rands24ProbabilityRand   float32 `elements:"name:rands_24_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_24_probability_rand;type:float;default:0;not null;" json:"rands_24_probability_rand"`
	Rands25IdRand            int     `elements:"name:rands_25_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_25_id_rand;type:integer;default:0;not null;" json:"rands_25_id_rand"`
	Rands25ProbabilityRand   float32 `elements:"name:rands_25_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_25_probability_rand;type:float;default:0;not null;" json:"rands_25_probability_rand"`
	Rands26IdRand            int     `elements:"name:rands_26_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_26_id_rand;type:integer;default:0;not null;" json:"rands_26_id_rand"`
	Rands26ProbabilityRand   float32 `elements:"name:rands_26_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_26_probability_rand;type:float;default:0;not null;" json:"rands_26_probability_rand"`
	Rands27IdRand            int     `elements:"name:rands_27_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_27_id_rand;type:integer;default:0;not null;" json:"rands_27_id_rand"`
	Rands27ProbabilityRand   float32 `elements:"name:rands_27_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_27_probability_rand;type:float;default:0;not null;" json:"rands_27_probability_rand"`
	Rands28IdRand            int     `elements:"name:rands_28_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_28_id_rand;type:integer;default:0;not null;" json:"rands_28_id_rand"`
	Rands28ProbabilityRand   float32 `elements:"name:rands_28_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_28_probability_rand;type:float;default:0;not null;" json:"rands_28_probability_rand"`
	Rands29IdRand            int     `elements:"name:rands_29_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_29_id_rand;type:integer;default:0;not null;" json:"rands_29_id_rand"`
	Rands29ProbabilityRand   float32 `elements:"name:rands_29_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_29_probability_rand;type:float;default:0;not null;" json:"rands_29_probability_rand"`
	Rands30IdRand            int     `elements:"name:rands_30_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_30_id_rand;type:integer;default:0;not null;" json:"rands_30_id_rand"`
	Rands30ProbabilityRand   float32 `elements:"name:rands_30_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_30_probability_rand;type:float;default:0;not null;" json:"rands_30_probability_rand"`
	Rands31IdRand            int     `elements:"name:rands_31_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_31_id_rand;type:integer;default:0;not null;" json:"rands_31_id_rand"`
	Rands31ProbabilityRand   float32 `elements:"name:rands_31_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_31_probability_rand;type:float;default:0;not null;" json:"rands_31_probability_rand"`
	Rands32IdRand            int     `elements:"name:rands_32_id_rand;type:int;size:4;text:未知;" gorm:"column:rands_32_id_rand;type:integer;default:0;not null;" json:"rands_32_id_rand"`
	Rands32ProbabilityRand   float32 `elements:"name:rands_32_probability_rand;type:float;size:4;text:未知;" gorm:"column:rands_32_probability_rand;type:float;default:0;not null;" json:"rands_32_probability_rand"`
	DurabilityDropMin        int     `elements:"name:durability_drop_min;type:int;size:4;text:未知;" gorm:"column:durability_drop_min;type:integer;default:0;not null;" json:"durability_drop_min"`
	DurabilityDropMax        int     `elements:"name:durability_drop_max;type:int;size:4;text:未知;" gorm:"column:durability_drop_max;type:integer;default:0;not null;" json:"durability_drop_max"`
	DecomposePrice           int     `elements:"name:decompose_price;type:int;size:4;text:未知;" gorm:"column:decompose_price;type:integer;default:0;not null;" json:"decompose_price"`
	DecomposeTime            int     `elements:"name:decompose_time;type:int;size:4;text:未知;" gorm:"column:decompose_time;type:integer;default:0;not null;" json:"decompose_time"`
	ElementId                int     `elements:"name:element_id;type:int;size:4;text:未知;" gorm:"column:element_id;type:integer;default:0;not null;" json:"element_id"`
	ElementNum               int     `elements:"name:element_num;type:int;size:4;text:未知;" gorm:"column:element_num;type:integer;default:0;not null;" json:"element_num"`
	PileNumMax               int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid                  int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType                 int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (DecorationEssence) TableName() string {
	return "el_decorationessence"
}

func (e *DecorationEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *DecorationEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
