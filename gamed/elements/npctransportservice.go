package elements

import (
	"pw_server/utils/structEx"
)

type NpcTransportService struct {
	ID          int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name        string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Routes1Id   int    `elements:"name:routes_1_id;type:int;size:4;text:未知;" gorm:"column:routes_1_id;type:integer;default:0;not null;" json:"routes_1_id"`
	Routes1Fee  int    `elements:"name:routes_1_fee;type:int;size:4;text:未知;" gorm:"column:routes_1_fee;type:integer;default:0;not null;" json:"routes_1_fee"`
	Routes2Id   int    `elements:"name:routes_2_id;type:int;size:4;text:未知;" gorm:"column:routes_2_id;type:integer;default:0;not null;" json:"routes_2_id"`
	Routes2Fee  int    `elements:"name:routes_2_fee;type:int;size:4;text:未知;" gorm:"column:routes_2_fee;type:integer;default:0;not null;" json:"routes_2_fee"`
	Routes3Id   int    `elements:"name:routes_3_id;type:int;size:4;text:未知;" gorm:"column:routes_3_id;type:integer;default:0;not null;" json:"routes_3_id"`
	Routes3Fee  int    `elements:"name:routes_3_fee;type:int;size:4;text:未知;" gorm:"column:routes_3_fee;type:integer;default:0;not null;" json:"routes_3_fee"`
	Routes4Id   int    `elements:"name:routes_4_id;type:int;size:4;text:未知;" gorm:"column:routes_4_id;type:integer;default:0;not null;" json:"routes_4_id"`
	Routes4Fee  int    `elements:"name:routes_4_fee;type:int;size:4;text:未知;" gorm:"column:routes_4_fee;type:integer;default:0;not null;" json:"routes_4_fee"`
	Routes5Id   int    `elements:"name:routes_5_id;type:int;size:4;text:未知;" gorm:"column:routes_5_id;type:integer;default:0;not null;" json:"routes_5_id"`
	Routes5Fee  int    `elements:"name:routes_5_fee;type:int;size:4;text:未知;" gorm:"column:routes_5_fee;type:integer;default:0;not null;" json:"routes_5_fee"`
	Routes6Id   int    `elements:"name:routes_6_id;type:int;size:4;text:未知;" gorm:"column:routes_6_id;type:integer;default:0;not null;" json:"routes_6_id"`
	Routes6Fee  int    `elements:"name:routes_6_fee;type:int;size:4;text:未知;" gorm:"column:routes_6_fee;type:integer;default:0;not null;" json:"routes_6_fee"`
	Routes7Id   int    `elements:"name:routes_7_id;type:int;size:4;text:未知;" gorm:"column:routes_7_id;type:integer;default:0;not null;" json:"routes_7_id"`
	Routes7Fee  int    `elements:"name:routes_7_fee;type:int;size:4;text:未知;" gorm:"column:routes_7_fee;type:integer;default:0;not null;" json:"routes_7_fee"`
	Routes8Id   int    `elements:"name:routes_8_id;type:int;size:4;text:未知;" gorm:"column:routes_8_id;type:integer;default:0;not null;" json:"routes_8_id"`
	Routes8Fee  int    `elements:"name:routes_8_fee;type:int;size:4;text:未知;" gorm:"column:routes_8_fee;type:integer;default:0;not null;" json:"routes_8_fee"`
	Routes9Id   int    `elements:"name:routes_9_id;type:int;size:4;text:未知;" gorm:"column:routes_9_id;type:integer;default:0;not null;" json:"routes_9_id"`
	Routes9Fee  int    `elements:"name:routes_9_fee;type:int;size:4;text:未知;" gorm:"column:routes_9_fee;type:integer;default:0;not null;" json:"routes_9_fee"`
	Routes10Id  int    `elements:"name:routes_10_id;type:int;size:4;text:未知;" gorm:"column:routes_10_id;type:integer;default:0;not null;" json:"routes_10_id"`
	Routes10Fee int    `elements:"name:routes_10_fee;type:int;size:4;text:未知;" gorm:"column:routes_10_fee;type:integer;default:0;not null;" json:"routes_10_fee"`
	Routes11Id  int    `elements:"name:routes_11_id;type:int;size:4;text:未知;" gorm:"column:routes_11_id;type:integer;default:0;not null;" json:"routes_11_id"`
	Routes11Fee int    `elements:"name:routes_11_fee;type:int;size:4;text:未知;" gorm:"column:routes_11_fee;type:integer;default:0;not null;" json:"routes_11_fee"`
	Routes12Id  int    `elements:"name:routes_12_id;type:int;size:4;text:未知;" gorm:"column:routes_12_id;type:integer;default:0;not null;" json:"routes_12_id"`
	Routes12Fee int    `elements:"name:routes_12_fee;type:int;size:4;text:未知;" gorm:"column:routes_12_fee;type:integer;default:0;not null;" json:"routes_12_fee"`
	Routes13Id  int    `elements:"name:routes_13_id;type:int;size:4;text:未知;" gorm:"column:routes_13_id;type:integer;default:0;not null;" json:"routes_13_id"`
	Routes13Fee int    `elements:"name:routes_13_fee;type:int;size:4;text:未知;" gorm:"column:routes_13_fee;type:integer;default:0;not null;" json:"routes_13_fee"`
	Routes14Id  int    `elements:"name:routes_14_id;type:int;size:4;text:未知;" gorm:"column:routes_14_id;type:integer;default:0;not null;" json:"routes_14_id"`
	Routes14Fee int    `elements:"name:routes_14_fee;type:int;size:4;text:未知;" gorm:"column:routes_14_fee;type:integer;default:0;not null;" json:"routes_14_fee"`
	Routes15Id  int    `elements:"name:routes_15_id;type:int;size:4;text:未知;" gorm:"column:routes_15_id;type:integer;default:0;not null;" json:"routes_15_id"`
	Routes15Fee int    `elements:"name:routes_15_fee;type:int;size:4;text:未知;" gorm:"column:routes_15_fee;type:integer;default:0;not null;" json:"routes_15_fee"`
	Routes16Id  int    `elements:"name:routes_16_id;type:int;size:4;text:未知;" gorm:"column:routes_16_id;type:integer;default:0;not null;" json:"routes_16_id"`
	Routes16Fee int    `elements:"name:routes_16_fee;type:int;size:4;text:未知;" gorm:"column:routes_16_fee;type:integer;default:0;not null;" json:"routes_16_fee"`
	Routes17Id  int    `elements:"name:routes_17_id;type:int;size:4;text:未知;" gorm:"column:routes_17_id;type:integer;default:0;not null;" json:"routes_17_id"`
	Routes17Fee int    `elements:"name:routes_17_fee;type:int;size:4;text:未知;" gorm:"column:routes_17_fee;type:integer;default:0;not null;" json:"routes_17_fee"`
	Routes18Id  int    `elements:"name:routes_18_id;type:int;size:4;text:未知;" gorm:"column:routes_18_id;type:integer;default:0;not null;" json:"routes_18_id"`
	Routes18Fee int    `elements:"name:routes_18_fee;type:int;size:4;text:未知;" gorm:"column:routes_18_fee;type:integer;default:0;not null;" json:"routes_18_fee"`
	Routes19Id  int    `elements:"name:routes_19_id;type:int;size:4;text:未知;" gorm:"column:routes_19_id;type:integer;default:0;not null;" json:"routes_19_id"`
	Routes19Fee int    `elements:"name:routes_19_fee;type:int;size:4;text:未知;" gorm:"column:routes_19_fee;type:integer;default:0;not null;" json:"routes_19_fee"`
	Routes20Id  int    `elements:"name:routes_20_id;type:int;size:4;text:未知;" gorm:"column:routes_20_id;type:integer;default:0;not null;" json:"routes_20_id"`
	Routes20Fee int    `elements:"name:routes_20_fee;type:int;size:4;text:未知;" gorm:"column:routes_20_fee;type:integer;default:0;not null;" json:"routes_20_fee"`
	Routes21Id  int    `elements:"name:routes_21_id;type:int;size:4;text:未知;" gorm:"column:routes_21_id;type:integer;default:0;not null;" json:"routes_21_id"`
	Routes21Fee int    `elements:"name:routes_21_fee;type:int;size:4;text:未知;" gorm:"column:routes_21_fee;type:integer;default:0;not null;" json:"routes_21_fee"`
	Routes22Id  int    `elements:"name:routes_22_id;type:int;size:4;text:未知;" gorm:"column:routes_22_id;type:integer;default:0;not null;" json:"routes_22_id"`
	Routes22Fee int    `elements:"name:routes_22_fee;type:int;size:4;text:未知;" gorm:"column:routes_22_fee;type:integer;default:0;not null;" json:"routes_22_fee"`
	Routes23Id  int    `elements:"name:routes_23_id;type:int;size:4;text:未知;" gorm:"column:routes_23_id;type:integer;default:0;not null;" json:"routes_23_id"`
	Routes23Fee int    `elements:"name:routes_23_fee;type:int;size:4;text:未知;" gorm:"column:routes_23_fee;type:integer;default:0;not null;" json:"routes_23_fee"`
	Routes24Id  int    `elements:"name:routes_24_id;type:int;size:4;text:未知;" gorm:"column:routes_24_id;type:integer;default:0;not null;" json:"routes_24_id"`
	Routes24Fee int    `elements:"name:routes_24_fee;type:int;size:4;text:未知;" gorm:"column:routes_24_fee;type:integer;default:0;not null;" json:"routes_24_fee"`
	Routes25Id  int    `elements:"name:routes_25_id;type:int;size:4;text:未知;" gorm:"column:routes_25_id;type:integer;default:0;not null;" json:"routes_25_id"`
	Routes25Fee int    `elements:"name:routes_25_fee;type:int;size:4;text:未知;" gorm:"column:routes_25_fee;type:integer;default:0;not null;" json:"routes_25_fee"`
	Routes26Id  int    `elements:"name:routes_26_id;type:int;size:4;text:未知;" gorm:"column:routes_26_id;type:integer;default:0;not null;" json:"routes_26_id"`
	Routes26Fee int    `elements:"name:routes_26_fee;type:int;size:4;text:未知;" gorm:"column:routes_26_fee;type:integer;default:0;not null;" json:"routes_26_fee"`
	Routes27Id  int    `elements:"name:routes_27_id;type:int;size:4;text:未知;" gorm:"column:routes_27_id;type:integer;default:0;not null;" json:"routes_27_id"`
	Routes27Fee int    `elements:"name:routes_27_fee;type:int;size:4;text:未知;" gorm:"column:routes_27_fee;type:integer;default:0;not null;" json:"routes_27_fee"`
	Routes28Id  int    `elements:"name:routes_28_id;type:int;size:4;text:未知;" gorm:"column:routes_28_id;type:integer;default:0;not null;" json:"routes_28_id"`
	Routes28Fee int    `elements:"name:routes_28_fee;type:int;size:4;text:未知;" gorm:"column:routes_28_fee;type:integer;default:0;not null;" json:"routes_28_fee"`
	Routes29Id  int    `elements:"name:routes_29_id;type:int;size:4;text:未知;" gorm:"column:routes_29_id;type:integer;default:0;not null;" json:"routes_29_id"`
	Routes29Fee int    `elements:"name:routes_29_fee;type:int;size:4;text:未知;" gorm:"column:routes_29_fee;type:integer;default:0;not null;" json:"routes_29_fee"`
	Routes30Id  int    `elements:"name:routes_30_id;type:int;size:4;text:未知;" gorm:"column:routes_30_id;type:integer;default:0;not null;" json:"routes_30_id"`
	Routes30Fee int    `elements:"name:routes_30_fee;type:int;size:4;text:未知;" gorm:"column:routes_30_fee;type:integer;default:0;not null;" json:"routes_30_fee"`
	Routes31Id  int    `elements:"name:routes_31_id;type:int;size:4;text:未知;" gorm:"column:routes_31_id;type:integer;default:0;not null;" json:"routes_31_id"`
	Routes31Fee int    `elements:"name:routes_31_fee;type:int;size:4;text:未知;" gorm:"column:routes_31_fee;type:integer;default:0;not null;" json:"routes_31_fee"`
	Routes32Id  int    `elements:"name:routes_32_id;type:int;size:4;text:未知;" gorm:"column:routes_32_id;type:integer;default:0;not null;" json:"routes_32_id"`
	Routes32Fee int    `elements:"name:routes_32_fee;type:int;size:4;text:未知;" gorm:"column:routes_32_fee;type:integer;default:0;not null;" json:"routes_32_fee"`
	IdDialog    int    `elements:"name:id_dialog;type:int;size:4;text:未知;" gorm:"column:id_dialog;type:integer;default:0;not null;" json:"id_dialog"`
}

func (NpcTransportService) TableName() string {
	return "el_npctransportservice"
}

func (e *NpcTransportService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcTransportService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
