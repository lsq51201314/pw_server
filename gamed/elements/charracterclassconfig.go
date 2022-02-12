package elements

import (
	"pw_server/utils/structEx"
)

type CharracterClassConfig struct {
	ID                int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name              string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	CharacterClassId  int     `elements:"name:character_class_id;type:int;size:4;text:未知;" gorm:"column:character_class_id;type:integer;default:0;not null;" json:"character_class_id"`
	Faction           int     `elements:"name:faction;type:int;size:4;text:未知;" gorm:"column:faction;type:integer;default:0;not null;" json:"faction"`
	EnemyFaction      int     `elements:"name:enemy_faction;type:int;size:4;text:未知;" gorm:"column:enemy_faction;type:integer;default:0;not null;" json:"enemy_faction"`
	AttackSpeed       float32 `elements:"name:attack_speed;type:float;size:4;text:未知;" gorm:"column:attack_speed;type:float;default:0;not null;" json:"attack_speed"`
	AttackRange       float32 `elements:"name:attack_range;type:float;size:4;text:未知;" gorm:"column:attack_range;type:float;default:0;not null;" json:"attack_range"`
	HpGen             int     `elements:"name:hp_gen;type:int;size:4;text:未知;" gorm:"column:hp_gen;type:integer;default:0;not null;" json:"hp_gen"`
	MpGen             int     `elements:"name:mp_gen;type:int;size:4;text:未知;" gorm:"column:mp_gen;type:integer;default:0;not null;" json:"mp_gen"`
	WalkSpeed         float32 `elements:"name:walk_speed;type:float;size:4;text:未知;" gorm:"column:walk_speed;type:float;default:0;not null;" json:"walk_speed"`
	RunSpeed          float32 `elements:"name:run_speed;type:float;size:4;text:未知;" gorm:"column:run_speed;type:float;default:0;not null;" json:"run_speed"`
	SwimSpeed         float32 `elements:"name:swim_speed;type:float;size:4;text:未知;" gorm:"column:swim_speed;type:float;default:0;not null;" json:"swim_speed"`
	FlySpeed          float32 `elements:"name:fly_speed;type:float;size:4;text:未知;" gorm:"column:fly_speed;type:float;default:0;not null;" json:"fly_speed"`
	CritRate          int     `elements:"name:crit_rate;type:int;size:4;text:未知;" gorm:"column:crit_rate;type:integer;default:0;not null;" json:"crit_rate"`
	VitHp             int     `elements:"name:vit_hp;type:int;size:4;text:未知;" gorm:"column:vit_hp;type:integer;default:0;not null;" json:"vit_hp"`
	EngMp             int     `elements:"name:eng_mp;type:int;size:4;text:未知;" gorm:"column:eng_mp;type:integer;default:0;not null;" json:"eng_mp"`
	AgiAttack         int     `elements:"name:agi_attack;type:int;size:4;text:未知;" gorm:"column:agi_attack;type:integer;default:0;not null;" json:"agi_attack"`
	AgiArmor          int     `elements:"name:agi_armor;type:int;size:4;text:未知;" gorm:"column:agi_armor;type:integer;default:0;not null;" json:"agi_armor"`
	LvlupHp           int     `elements:"name:lvlup_hp;type:int;size:4;text:未知;" gorm:"column:lvlup_hp;type:integer;default:0;not null;" json:"lvlup_hp"`
	LvlupMp           int     `elements:"name:lvlup_mp;type:int;size:4;text:未知;" gorm:"column:lvlup_mp;type:integer;default:0;not null;" json:"lvlup_mp"`
	LvlupDmg          float32 `elements:"name:lvlup_dmg;type:float;size:4;text:未知;" gorm:"column:lvlup_dmg;type:float;default:0;not null;" json:"lvlup_dmg"`
	LvlupMagic        float32 `elements:"name:lvlup_magic;type:float;size:4;text:未知;" gorm:"column:lvlup_magic;type:float;default:0;not null;" json:"lvlup_magic"`
	LvlupDefense      float32 `elements:"name:lvlup_defense;type:float;size:4;text:未知;" gorm:"column:lvlup_defense;type:float;default:0;not null;" json:"lvlup_defense"`
	LvlupMagicdefence float32 `elements:"name:lvlup_magicdefence;type:float;size:4;text:未知;" gorm:"column:lvlup_magicdefence;type:float;default:0;not null;" json:"lvlup_magicdefence"`
	AngroIncrease     int     `elements:"name:angro_increase;type:int;size:4;text:未知;" gorm:"column:angro_increase;type:integer;default:0;not null;" json:"angro_increase"`
}

func (CharracterClassConfig) TableName() string {
	return "el_charracterclassconfig"
}

func (e *CharracterClassConfig) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *CharracterClassConfig) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
