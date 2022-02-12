package elements

import (
	"pw_server/utils/structEx"
)

type WeaponSubType struct {
	ID                 int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name               string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileHitgfx         string  `elements:"name:file_hitgfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitgfx;type:varchar(128);not null;" json:"file_hitgfx"`
	FileHitsfx         string  `elements:"name:file_hitsfx;type:gbk;size:128;text:未知;" gorm:"column:file_hitsfx;type:varchar(128);not null;" json:"file_hitsfx"`
	ProbabilityFastest float32 `elements:"name:probability_fastest;type:float;size:4;text:未知;" gorm:"column:probability_fastest;type:float;default:0;not null;" json:"probability_fastest"`
	ProbabilityFast    float32 `elements:"name:probability_fast;type:float;size:4;text:未知;" gorm:"column:probability_fast;type:float;default:0;not null;" json:"probability_fast"`
	ProbabilityNormal  float32 `elements:"name:probability_normal;type:float;size:4;text:未知;" gorm:"column:probability_normal;type:float;default:0;not null;" json:"probability_normal"`
	ProbabilitySlow    float32 `elements:"name:probability_slow;type:float;size:4;text:未知;" gorm:"column:probability_slow;type:float;default:0;not null;" json:"probability_slow"`
	ProbabilitySlowest float32 `elements:"name:probability_slowest;type:float;size:4;text:未知;" gorm:"column:probability_slowest;type:float;default:0;not null;" json:"probability_slowest"`
	AttackSpeed        float32 `elements:"name:attack_speed;type:float;size:4;text:未知;" gorm:"column:attack_speed;type:float;default:0;not null;" json:"attack_speed"`
	AttackShortRange   float32 `elements:"name:attack_short_range;type:float;size:4;text:未知;" gorm:"column:attack_short_range;type:float;default:0;not null;" json:"attack_short_range"`
	ActionType         int     `elements:"name:action_type;type:int;size:4;text:未知;" gorm:"column:action_type;type:integer;default:0;not null;" json:"action_type"`
}

func (WeaponSubType) TableName() string {
	return "el_weaponsubtype"
}

func (e *WeaponSubType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *WeaponSubType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
