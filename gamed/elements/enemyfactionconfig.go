package elements

import (
	"pw_server/utils/structEx"
)

type EnemyFactionConfig struct {
	ID              int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name            string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	EnemyFactions1  int    `elements:"name:enemy_factions_1;type:int;size:4;text:未知;" gorm:"column:enemy_factions_1;type:integer;default:0;not null;" json:"enemy_factions_1"`
	EnemyFactions2  int    `elements:"name:enemy_factions_2;type:int;size:4;text:未知;" gorm:"column:enemy_factions_2;type:integer;default:0;not null;" json:"enemy_factions_2"`
	EnemyFactions3  int    `elements:"name:enemy_factions_3;type:int;size:4;text:未知;" gorm:"column:enemy_factions_3;type:integer;default:0;not null;" json:"enemy_factions_3"`
	EnemyFactions4  int    `elements:"name:enemy_factions_4;type:int;size:4;text:未知;" gorm:"column:enemy_factions_4;type:integer;default:0;not null;" json:"enemy_factions_4"`
	EnemyFactions5  int    `elements:"name:enemy_factions_5;type:int;size:4;text:未知;" gorm:"column:enemy_factions_5;type:integer;default:0;not null;" json:"enemy_factions_5"`
	EnemyFactions6  int    `elements:"name:enemy_factions_6;type:int;size:4;text:未知;" gorm:"column:enemy_factions_6;type:integer;default:0;not null;" json:"enemy_factions_6"`
	EnemyFactions7  int    `elements:"name:enemy_factions_7;type:int;size:4;text:未知;" gorm:"column:enemy_factions_7;type:integer;default:0;not null;" json:"enemy_factions_7"`
	EnemyFactions8  int    `elements:"name:enemy_factions_8;type:int;size:4;text:未知;" gorm:"column:enemy_factions_8;type:integer;default:0;not null;" json:"enemy_factions_8"`
	EnemyFactions9  int    `elements:"name:enemy_factions_9;type:int;size:4;text:未知;" gorm:"column:enemy_factions_9;type:integer;default:0;not null;" json:"enemy_factions_9"`
	EnemyFactions10 int    `elements:"name:enemy_factions_10;type:int;size:4;text:未知;" gorm:"column:enemy_factions_10;type:integer;default:0;not null;" json:"enemy_factions_10"`
	EnemyFactions11 int    `elements:"name:enemy_factions_11;type:int;size:4;text:未知;" gorm:"column:enemy_factions_11;type:integer;default:0;not null;" json:"enemy_factions_11"`
	EnemyFactions12 int    `elements:"name:enemy_factions_12;type:int;size:4;text:未知;" gorm:"column:enemy_factions_12;type:integer;default:0;not null;" json:"enemy_factions_12"`
	EnemyFactions13 int    `elements:"name:enemy_factions_13;type:int;size:4;text:未知;" gorm:"column:enemy_factions_13;type:integer;default:0;not null;" json:"enemy_factions_13"`
	EnemyFactions14 int    `elements:"name:enemy_factions_14;type:int;size:4;text:未知;" gorm:"column:enemy_factions_14;type:integer;default:0;not null;" json:"enemy_factions_14"`
	EnemyFactions15 int    `elements:"name:enemy_factions_15;type:int;size:4;text:未知;" gorm:"column:enemy_factions_15;type:integer;default:0;not null;" json:"enemy_factions_15"`
	EnemyFactions16 int    `elements:"name:enemy_factions_16;type:int;size:4;text:未知;" gorm:"column:enemy_factions_16;type:integer;default:0;not null;" json:"enemy_factions_16"`
	EnemyFactions17 int    `elements:"name:enemy_factions_17;type:int;size:4;text:未知;" gorm:"column:enemy_factions_17;type:integer;default:0;not null;" json:"enemy_factions_17"`
	EnemyFactions18 int    `elements:"name:enemy_factions_18;type:int;size:4;text:未知;" gorm:"column:enemy_factions_18;type:integer;default:0;not null;" json:"enemy_factions_18"`
	EnemyFactions19 int    `elements:"name:enemy_factions_19;type:int;size:4;text:未知;" gorm:"column:enemy_factions_19;type:integer;default:0;not null;" json:"enemy_factions_19"`
	EnemyFactions20 int    `elements:"name:enemy_factions_20;type:int;size:4;text:未知;" gorm:"column:enemy_factions_20;type:integer;default:0;not null;" json:"enemy_factions_20"`
	EnemyFactions21 int    `elements:"name:enemy_factions_21;type:int;size:4;text:未知;" gorm:"column:enemy_factions_21;type:integer;default:0;not null;" json:"enemy_factions_21"`
	EnemyFactions22 int    `elements:"name:enemy_factions_22;type:int;size:4;text:未知;" gorm:"column:enemy_factions_22;type:integer;default:0;not null;" json:"enemy_factions_22"`
	EnemyFactions23 int    `elements:"name:enemy_factions_23;type:int;size:4;text:未知;" gorm:"column:enemy_factions_23;type:integer;default:0;not null;" json:"enemy_factions_23"`
	EnemyFactions24 int    `elements:"name:enemy_factions_24;type:int;size:4;text:未知;" gorm:"column:enemy_factions_24;type:integer;default:0;not null;" json:"enemy_factions_24"`
	EnemyFactions25 int    `elements:"name:enemy_factions_25;type:int;size:4;text:未知;" gorm:"column:enemy_factions_25;type:integer;default:0;not null;" json:"enemy_factions_25"`
	EnemyFactions26 int    `elements:"name:enemy_factions_26;type:int;size:4;text:未知;" gorm:"column:enemy_factions_26;type:integer;default:0;not null;" json:"enemy_factions_26"`
	EnemyFactions27 int    `elements:"name:enemy_factions_27;type:int;size:4;text:未知;" gorm:"column:enemy_factions_27;type:integer;default:0;not null;" json:"enemy_factions_27"`
	EnemyFactions28 int    `elements:"name:enemy_factions_28;type:int;size:4;text:未知;" gorm:"column:enemy_factions_28;type:integer;default:0;not null;" json:"enemy_factions_28"`
	EnemyFactions29 int    `elements:"name:enemy_factions_29;type:int;size:4;text:未知;" gorm:"column:enemy_factions_29;type:integer;default:0;not null;" json:"enemy_factions_29"`
	EnemyFactions30 int    `elements:"name:enemy_factions_30;type:int;size:4;text:未知;" gorm:"column:enemy_factions_30;type:integer;default:0;not null;" json:"enemy_factions_30"`
	EnemyFactions31 int    `elements:"name:enemy_factions_31;type:int;size:4;text:未知;" gorm:"column:enemy_factions_31;type:integer;default:0;not null;" json:"enemy_factions_31"`
	EnemyFactions32 int    `elements:"name:enemy_factions_32;type:int;size:4;text:未知;" gorm:"column:enemy_factions_32;type:integer;default:0;not null;" json:"enemy_factions_32"`
}

func (EnemyFactionConfig) TableName() string {
	return "el_enemyfactionconfig"
}

func (e *EnemyFactionConfig) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *EnemyFactionConfig) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
