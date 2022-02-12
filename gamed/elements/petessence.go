package elements

import (
	"pw_server/utils/structEx"
)

type PetEssence struct {
	ID               int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	IdType           int     `elements:"name:id_type;type:int;size:4;text:未知;" gorm:"column:id_type;type:integer;default:0;not null;" json:"id_type"`
	Name             string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileModel        string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	FileIcon         string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	CharacterComboId int     `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	LevelMax         int     `elements:"name:level_max;type:int;size:4;text:未知;" gorm:"column:level_max;type:integer;default:0;not null;" json:"level_max"`
	LevelRequire     int     `elements:"name:level_require;type:int;size:4;text:未知;" gorm:"column:level_require;type:integer;default:0;not null;" json:"level_require"`
	HpA              float32 `elements:"name:hp_a;type:float;size:4;text:未知;" gorm:"column:hp_a;type:float;default:0;not null;" json:"hp_a"`
	HpB              float32 `elements:"name:hp_b;type:float;size:4;text:未知;" gorm:"column:hp_b;type:float;default:0;not null;" json:"hp_b"`
	HpC              float32 `elements:"name:hp_c;type:float;size:4;text:未知;" gorm:"column:hp_c;type:float;default:0;not null;" json:"hp_c"`
	HpGenA           float32 `elements:"name:hp_gen_a;type:float;size:4;text:未知;" gorm:"column:hp_gen_a;type:float;default:0;not null;" json:"hp_gen_a"`
	HpGenB           float32 `elements:"name:hp_gen_b;type:float;size:4;text:未知;" gorm:"column:hp_gen_b;type:float;default:0;not null;" json:"hp_gen_b"`
	HpGenC           float32 `elements:"name:hp_gen_c;type:float;size:4;text:未知;" gorm:"column:hp_gen_c;type:float;default:0;not null;" json:"hp_gen_c"`
	DamageA          float32 `elements:"name:damage_a;type:float;size:4;text:未知;" gorm:"column:damage_a;type:float;default:0;not null;" json:"damage_a"`
	DamageB          float32 `elements:"name:damage_b;type:float;size:4;text:未知;" gorm:"column:damage_b;type:float;default:0;not null;" json:"damage_b"`
	DamageC          float32 `elements:"name:damage_c;type:float;size:4;text:未知;" gorm:"column:damage_c;type:float;default:0;not null;" json:"damage_c"`
	DamageD          float32 `elements:"name:damage_d;type:float;size:4;text:未知;" gorm:"column:damage_d;type:float;default:0;not null;" json:"damage_d"`
	SpeedA           float32 `elements:"name:speed_a;type:float;size:4;text:未知;" gorm:"column:speed_a;type:float;default:0;not null;" json:"speed_a"`
	SpeedB           float32 `elements:"name:speed_b;type:float;size:4;text:未知;" gorm:"column:speed_b;type:float;default:0;not null;" json:"speed_b"`
	AttackA          float32 `elements:"name:attack_a;type:float;size:4;text:未知;" gorm:"column:attack_a;type:float;default:0;not null;" json:"attack_a"`
	AttackB          float32 `elements:"name:attack_b;type:float;size:4;text:未知;" gorm:"column:attack_b;type:float;default:0;not null;" json:"attack_b"`
	AttackC          float32 `elements:"name:attack_c;type:float;size:4;text:未知;" gorm:"column:attack_c;type:float;default:0;not null;" json:"attack_c"`
	ArmorA           float32 `elements:"name:armor_a;type:float;size:4;text:未知;" gorm:"column:armor_a;type:float;default:0;not null;" json:"armor_a"`
	ArmorB           float32 `elements:"name:armor_b;type:float;size:4;text:未知;" gorm:"column:armor_b;type:float;default:0;not null;" json:"armor_b"`
	ArmorC           float32 `elements:"name:armor_c;type:float;size:4;text:未知;" gorm:"column:armor_c;type:float;default:0;not null;" json:"armor_c"`
	PhysicDefenceA   float32 `elements:"name:physic_defence_a;type:float;size:4;text:未知;" gorm:"column:physic_defence_a;type:float;default:0;not null;" json:"physic_defence_a"`
	PhysicDefenceB   float32 `elements:"name:physic_defence_b;type:float;size:4;text:未知;" gorm:"column:physic_defence_b;type:float;default:0;not null;" json:"physic_defence_b"`
	PhysicDefenceC   float32 `elements:"name:physic_defence_c;type:float;size:4;text:未知;" gorm:"column:physic_defence_c;type:float;default:0;not null;" json:"physic_defence_c"`
	PhysicDefenceD   float32 `elements:"name:physic_defence_d;type:float;size:4;text:未知;" gorm:"column:physic_defence_d;type:float;default:0;not null;" json:"physic_defence_d"`
	MagicDefenceA    float32 `elements:"name:magic_defence_a;type:float;size:4;text:未知;" gorm:"column:magic_defence_a;type:float;default:0;not null;" json:"magic_defence_a"`
	MagicDefenceB    float32 `elements:"name:magic_defence_b;type:float;size:4;text:未知;" gorm:"column:magic_defence_b;type:float;default:0;not null;" json:"magic_defence_b"`
	MagicDefenceC    float32 `elements:"name:magic_defence_c;type:float;size:4;text:未知;" gorm:"column:magic_defence_c;type:float;default:0;not null;" json:"magic_defence_c"`
	MagicDefenceD    float32 `elements:"name:magic_defence_d;type:float;size:4;text:未知;" gorm:"column:magic_defence_d;type:float;default:0;not null;" json:"magic_defence_d"`
	Size             float32 `elements:"name:size;type:float;size:4;text:未知;" gorm:"column:size;type:float;default:0;not null;" json:"size"`
	DamageDelay      float32 `elements:"name:damage_delay;type:float;size:4;text:未知;" gorm:"column:damage_delay;type:float;default:0;not null;" json:"damage_delay"`
	AttackRange      float32 `elements:"name:attack_range;type:float;size:4;text:未知;" gorm:"column:attack_range;type:float;default:0;not null;" json:"attack_range"`
	AttackSpeed      float32 `elements:"name:attack_speed;type:float;size:4;text:未知;" gorm:"column:attack_speed;type:float;default:0;not null;" json:"attack_speed"`
	SightRange       int     `elements:"name:sight_range;type:int;size:4;text:未知;" gorm:"column:sight_range;type:integer;default:0;not null;" json:"sight_range"`
	FoodMask         int     `elements:"name:food_mask;type:int;size:4;text:未知;" gorm:"column:food_mask;type:integer;default:0;not null;" json:"food_mask"`
	InhabitType      int     `elements:"name:inhabit_type;type:int;size:4;text:未知;" gorm:"column:inhabit_type;type:integer;default:0;not null;" json:"inhabit_type"`
	StandMode        int     `elements:"name:stand_mode;type:int;size:4;text:未知;" gorm:"column:stand_mode;type:integer;default:0;not null;" json:"stand_mode"`
}

func (PetEssence) TableName() string {
	return "el_petessence"
}

func (e *PetEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *PetEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
