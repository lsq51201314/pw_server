package elements

import (
	"pw_server/utils/structEx"
)

type NpcUninstallService struct {
	ID        int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name      string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdGoods1  int    `elements:"name:id_goods_1;type:int;size:4;text:未知;" gorm:"column:id_goods_1;type:integer;default:0;not null;" json:"id_goods_1"`
	IdGoods2  int    `elements:"name:id_goods_2;type:int;size:4;text:未知;" gorm:"column:id_goods_2;type:integer;default:0;not null;" json:"id_goods_2"`
	IdGoods3  int    `elements:"name:id_goods_3;type:int;size:4;text:未知;" gorm:"column:id_goods_3;type:integer;default:0;not null;" json:"id_goods_3"`
	IdGoods4  int    `elements:"name:id_goods_4;type:int;size:4;text:未知;" gorm:"column:id_goods_4;type:integer;default:0;not null;" json:"id_goods_4"`
	IdGoods5  int    `elements:"name:id_goods_5;type:int;size:4;text:未知;" gorm:"column:id_goods_5;type:integer;default:0;not null;" json:"id_goods_5"`
	IdGoods6  int    `elements:"name:id_goods_6;type:int;size:4;text:未知;" gorm:"column:id_goods_6;type:integer;default:0;not null;" json:"id_goods_6"`
	IdGoods7  int    `elements:"name:id_goods_7;type:int;size:4;text:未知;" gorm:"column:id_goods_7;type:integer;default:0;not null;" json:"id_goods_7"`
	IdGoods8  int    `elements:"name:id_goods_8;type:int;size:4;text:未知;" gorm:"column:id_goods_8;type:integer;default:0;not null;" json:"id_goods_8"`
	IdGoods9  int    `elements:"name:id_goods_9;type:int;size:4;text:未知;" gorm:"column:id_goods_9;type:integer;default:0;not null;" json:"id_goods_9"`
	IdGoods10 int    `elements:"name:id_goods_10;type:int;size:4;text:未知;" gorm:"column:id_goods_10;type:integer;default:0;not null;" json:"id_goods_10"`
	IdGoods11 int    `elements:"name:id_goods_11;type:int;size:4;text:未知;" gorm:"column:id_goods_11;type:integer;default:0;not null;" json:"id_goods_11"`
	IdGoods12 int    `elements:"name:id_goods_12;type:int;size:4;text:未知;" gorm:"column:id_goods_12;type:integer;default:0;not null;" json:"id_goods_12"`
	IdGoods13 int    `elements:"name:id_goods_13;type:int;size:4;text:未知;" gorm:"column:id_goods_13;type:integer;default:0;not null;" json:"id_goods_13"`
	IdGoods14 int    `elements:"name:id_goods_14;type:int;size:4;text:未知;" gorm:"column:id_goods_14;type:integer;default:0;not null;" json:"id_goods_14"`
	IdGoods15 int    `elements:"name:id_goods_15;type:int;size:4;text:未知;" gorm:"column:id_goods_15;type:integer;default:0;not null;" json:"id_goods_15"`
	IdGoods16 int    `elements:"name:id_goods_16;type:int;size:4;text:未知;" gorm:"column:id_goods_16;type:integer;default:0;not null;" json:"id_goods_16"`
	IdGoods17 int    `elements:"name:id_goods_17;type:int;size:4;text:未知;" gorm:"column:id_goods_17;type:integer;default:0;not null;" json:"id_goods_17"`
	IdGoods18 int    `elements:"name:id_goods_18;type:int;size:4;text:未知;" gorm:"column:id_goods_18;type:integer;default:0;not null;" json:"id_goods_18"`
	IdGoods19 int    `elements:"name:id_goods_19;type:int;size:4;text:未知;" gorm:"column:id_goods_19;type:integer;default:0;not null;" json:"id_goods_19"`
	IdGoods20 int    `elements:"name:id_goods_20;type:int;size:4;text:未知;" gorm:"column:id_goods_20;type:integer;default:0;not null;" json:"id_goods_20"`
	IdGoods21 int    `elements:"name:id_goods_21;type:int;size:4;text:未知;" gorm:"column:id_goods_21;type:integer;default:0;not null;" json:"id_goods_21"`
	IdGoods22 int    `elements:"name:id_goods_22;type:int;size:4;text:未知;" gorm:"column:id_goods_22;type:integer;default:0;not null;" json:"id_goods_22"`
	IdGoods23 int    `elements:"name:id_goods_23;type:int;size:4;text:未知;" gorm:"column:id_goods_23;type:integer;default:0;not null;" json:"id_goods_23"`
	IdGoods24 int    `elements:"name:id_goods_24;type:int;size:4;text:未知;" gorm:"column:id_goods_24;type:integer;default:0;not null;" json:"id_goods_24"`
	IdGoods25 int    `elements:"name:id_goods_25;type:int;size:4;text:未知;" gorm:"column:id_goods_25;type:integer;default:0;not null;" json:"id_goods_25"`
	IdGoods26 int    `elements:"name:id_goods_26;type:int;size:4;text:未知;" gorm:"column:id_goods_26;type:integer;default:0;not null;" json:"id_goods_26"`
	IdGoods27 int    `elements:"name:id_goods_27;type:int;size:4;text:未知;" gorm:"column:id_goods_27;type:integer;default:0;not null;" json:"id_goods_27"`
	IdGoods28 int    `elements:"name:id_goods_28;type:int;size:4;text:未知;" gorm:"column:id_goods_28;type:integer;default:0;not null;" json:"id_goods_28"`
	IdGoods29 int    `elements:"name:id_goods_29;type:int;size:4;text:未知;" gorm:"column:id_goods_29;type:integer;default:0;not null;" json:"id_goods_29"`
	IdGoods30 int    `elements:"name:id_goods_30;type:int;size:4;text:未知;" gorm:"column:id_goods_30;type:integer;default:0;not null;" json:"id_goods_30"`
	IdGoods31 int    `elements:"name:id_goods_31;type:int;size:4;text:未知;" gorm:"column:id_goods_31;type:integer;default:0;not null;" json:"id_goods_31"`
	IdGoods32 int    `elements:"name:id_goods_32;type:int;size:4;text:未知;" gorm:"column:id_goods_32;type:integer;default:0;not null;" json:"id_goods_32"`
	IdDialog  int    `elements:"name:id_dialog;type:int;size:4;text:未知;" gorm:"column:id_dialog;type:integer;default:0;not null;" json:"id_dialog"`
}

func (NpcUninstallService) TableName() string {
	return "el_npcuninstallservice"
}

func (e *NpcUninstallService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcUninstallService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
