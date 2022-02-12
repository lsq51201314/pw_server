package elements

import (
	"pw_server/utils/structEx"
)

type PlayerActionInfoConfig struct {
	ID                         int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                       string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	ActionName                 string `elements:"name:action_name;type:gbk;size:32;text:未知;" gorm:"column:action_name;type:varchar(32);not null;" json:"action_name"`
	ActionPrefix               string `elements:"name:action_prefix;type:gbk;size:32;text:未知;" gorm:"column:action_prefix;type:varchar(32);not null;" json:"action_prefix"`
	ActionWeaponSuffix1Suffix  string `elements:"name:action_weapon_suffix_1_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_1_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_1_suffix"`
	ActionWeaponSuffix2Suffix  string `elements:"name:action_weapon_suffix_2_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_2_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_2_suffix"`
	ActionWeaponSuffix3Suffix  string `elements:"name:action_weapon_suffix_3_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_3_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_3_suffix"`
	ActionWeaponSuffix4Suffix  string `elements:"name:action_weapon_suffix_4_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_4_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_4_suffix"`
	ActionWeaponSuffix5Suffix  string `elements:"name:action_weapon_suffix_5_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_5_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_5_suffix"`
	ActionWeaponSuffix6Suffix  string `elements:"name:action_weapon_suffix_6_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_6_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_6_suffix"`
	ActionWeaponSuffix7Suffix  string `elements:"name:action_weapon_suffix_7_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_7_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_7_suffix"`
	ActionWeaponSuffix8Suffix  string `elements:"name:action_weapon_suffix_8_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_8_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_8_suffix"`
	ActionWeaponSuffix9Suffix  string `elements:"name:action_weapon_suffix_9_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_9_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_9_suffix"`
	ActionWeaponSuffix10Suffix string `elements:"name:action_weapon_suffix_10_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_10_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_10_suffix"`
	ActionWeaponSuffix11Suffix string `elements:"name:action_weapon_suffix_11_suffix;type:gbk;size:32;text:未知;" gorm:"column:action_weapon_suffix_11_suffix;type:varchar(32);not null;" json:"action_weapon_suffix_11_suffix"`
	HideWeapon                 int    `elements:"name:hide_weapon;type:int;size:4;text:未知;" gorm:"column:hide_weapon;type:integer;default:0;not null;" json:"hide_weapon"`
}

func (PlayerActionInfoConfig) TableName() string {
	return "el_playeractioninfoconfig"
}

func (e *PlayerActionInfoConfig) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *PlayerActionInfoConfig) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
