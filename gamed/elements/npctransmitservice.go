package elements

import (
	"pw_server/utils/structEx"
)

type NpcTransmitService struct {
	ID                     int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                   string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	NumTargets             int    `elements:"name:num_targets;type:int;size:4;text:未知;" gorm:"column:num_targets;type:integer;default:0;not null;" json:"num_targets"`
	Targets1Idtarget       int    `elements:"name:targets_1_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_1_idtarget;type:integer;default:0;not null;" json:"targets_1_idtarget"`
	Targets1Fee            int    `elements:"name:targets_1_fee;type:int;size:4;text:未知;" gorm:"column:targets_1_fee;type:integer;default:0;not null;" json:"targets_1_fee"`
	Targets1RequiredLevel  int    `elements:"name:targets_1_required_level;type:int;size:4;text:未知;" gorm:"column:targets_1_required_level;type:integer;default:0;not null;" json:"targets_1_required_level"`
	Targets2Idtarget       int    `elements:"name:targets_2_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_2_idtarget;type:integer;default:0;not null;" json:"targets_2_idtarget"`
	Targets2Fee            int    `elements:"name:targets_2_fee;type:int;size:4;text:未知;" gorm:"column:targets_2_fee;type:integer;default:0;not null;" json:"targets_2_fee"`
	Targets2RequiredLevel  int    `elements:"name:targets_2_required_level;type:int;size:4;text:未知;" gorm:"column:targets_2_required_level;type:integer;default:0;not null;" json:"targets_2_required_level"`
	Targets3Idtarget       int    `elements:"name:targets_3_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_3_idtarget;type:integer;default:0;not null;" json:"targets_3_idtarget"`
	Targets3Fee            int    `elements:"name:targets_3_fee;type:int;size:4;text:未知;" gorm:"column:targets_3_fee;type:integer;default:0;not null;" json:"targets_3_fee"`
	Targets3RequiredLevel  int    `elements:"name:targets_3_required_level;type:int;size:4;text:未知;" gorm:"column:targets_3_required_level;type:integer;default:0;not null;" json:"targets_3_required_level"`
	Targets4Idtarget       int    `elements:"name:targets_4_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_4_idtarget;type:integer;default:0;not null;" json:"targets_4_idtarget"`
	Targets4Fee            int    `elements:"name:targets_4_fee;type:int;size:4;text:未知;" gorm:"column:targets_4_fee;type:integer;default:0;not null;" json:"targets_4_fee"`
	Targets4RequiredLevel  int    `elements:"name:targets_4_required_level;type:int;size:4;text:未知;" gorm:"column:targets_4_required_level;type:integer;default:0;not null;" json:"targets_4_required_level"`
	Targets5Idtarget       int    `elements:"name:targets_5_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_5_idtarget;type:integer;default:0;not null;" json:"targets_5_idtarget"`
	Targets5Fee            int    `elements:"name:targets_5_fee;type:int;size:4;text:未知;" gorm:"column:targets_5_fee;type:integer;default:0;not null;" json:"targets_5_fee"`
	Targets5RequiredLevel  int    `elements:"name:targets_5_required_level;type:int;size:4;text:未知;" gorm:"column:targets_5_required_level;type:integer;default:0;not null;" json:"targets_5_required_level"`
	Targets6Idtarget       int    `elements:"name:targets_6_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_6_idtarget;type:integer;default:0;not null;" json:"targets_6_idtarget"`
	Targets6Fee            int    `elements:"name:targets_6_fee;type:int;size:4;text:未知;" gorm:"column:targets_6_fee;type:integer;default:0;not null;" json:"targets_6_fee"`
	Targets6RequiredLevel  int    `elements:"name:targets_6_required_level;type:int;size:4;text:未知;" gorm:"column:targets_6_required_level;type:integer;default:0;not null;" json:"targets_6_required_level"`
	Targets7Idtarget       int    `elements:"name:targets_7_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_7_idtarget;type:integer;default:0;not null;" json:"targets_7_idtarget"`
	Targets7Fee            int    `elements:"name:targets_7_fee;type:int;size:4;text:未知;" gorm:"column:targets_7_fee;type:integer;default:0;not null;" json:"targets_7_fee"`
	Targets7RequiredLevel  int    `elements:"name:targets_7_required_level;type:int;size:4;text:未知;" gorm:"column:targets_7_required_level;type:integer;default:0;not null;" json:"targets_7_required_level"`
	Targets8Idtarget       int    `elements:"name:targets_8_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_8_idtarget;type:integer;default:0;not null;" json:"targets_8_idtarget"`
	Targets8Fee            int    `elements:"name:targets_8_fee;type:int;size:4;text:未知;" gorm:"column:targets_8_fee;type:integer;default:0;not null;" json:"targets_8_fee"`
	Targets8RequiredLevel  int    `elements:"name:targets_8_required_level;type:int;size:4;text:未知;" gorm:"column:targets_8_required_level;type:integer;default:0;not null;" json:"targets_8_required_level"`
	Targets9Idtarget       int    `elements:"name:targets_9_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_9_idtarget;type:integer;default:0;not null;" json:"targets_9_idtarget"`
	Targets9Fee            int    `elements:"name:targets_9_fee;type:int;size:4;text:未知;" gorm:"column:targets_9_fee;type:integer;default:0;not null;" json:"targets_9_fee"`
	Targets9RequiredLevel  int    `elements:"name:targets_9_required_level;type:int;size:4;text:未知;" gorm:"column:targets_9_required_level;type:integer;default:0;not null;" json:"targets_9_required_level"`
	Targets10Idtarget      int    `elements:"name:targets_10_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_10_idtarget;type:integer;default:0;not null;" json:"targets_10_idtarget"`
	Targets10Fee           int    `elements:"name:targets_10_fee;type:int;size:4;text:未知;" gorm:"column:targets_10_fee;type:integer;default:0;not null;" json:"targets_10_fee"`
	Targets10RequiredLevel int    `elements:"name:targets_10_required_level;type:int;size:4;text:未知;" gorm:"column:targets_10_required_level;type:integer;default:0;not null;" json:"targets_10_required_level"`
	Targets11Idtarget      int    `elements:"name:targets_11_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_11_idtarget;type:integer;default:0;not null;" json:"targets_11_idtarget"`
	Targets11Fee           int    `elements:"name:targets_11_fee;type:int;size:4;text:未知;" gorm:"column:targets_11_fee;type:integer;default:0;not null;" json:"targets_11_fee"`
	Targets11RequiredLevel int    `elements:"name:targets_11_required_level;type:int;size:4;text:未知;" gorm:"column:targets_11_required_level;type:integer;default:0;not null;" json:"targets_11_required_level"`
	Targets12Idtarget      int    `elements:"name:targets_12_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_12_idtarget;type:integer;default:0;not null;" json:"targets_12_idtarget"`
	Targets12Fee           int    `elements:"name:targets_12_fee;type:int;size:4;text:未知;" gorm:"column:targets_12_fee;type:integer;default:0;not null;" json:"targets_12_fee"`
	Targets12RequiredLevel int    `elements:"name:targets_12_required_level;type:int;size:4;text:未知;" gorm:"column:targets_12_required_level;type:integer;default:0;not null;" json:"targets_12_required_level"`
	Targets13Idtarget      int    `elements:"name:targets_13_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_13_idtarget;type:integer;default:0;not null;" json:"targets_13_idtarget"`
	Targets13Fee           int    `elements:"name:targets_13_fee;type:int;size:4;text:未知;" gorm:"column:targets_13_fee;type:integer;default:0;not null;" json:"targets_13_fee"`
	Targets13RequiredLevel int    `elements:"name:targets_13_required_level;type:int;size:4;text:未知;" gorm:"column:targets_13_required_level;type:integer;default:0;not null;" json:"targets_13_required_level"`
	Targets14Idtarget      int    `elements:"name:targets_14_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_14_idtarget;type:integer;default:0;not null;" json:"targets_14_idtarget"`
	Targets14Fee           int    `elements:"name:targets_14_fee;type:int;size:4;text:未知;" gorm:"column:targets_14_fee;type:integer;default:0;not null;" json:"targets_14_fee"`
	Targets14RequiredLevel int    `elements:"name:targets_14_required_level;type:int;size:4;text:未知;" gorm:"column:targets_14_required_level;type:integer;default:0;not null;" json:"targets_14_required_level"`
	Targets15Idtarget      int    `elements:"name:targets_15_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_15_idtarget;type:integer;default:0;not null;" json:"targets_15_idtarget"`
	Targets15Fee           int    `elements:"name:targets_15_fee;type:int;size:4;text:未知;" gorm:"column:targets_15_fee;type:integer;default:0;not null;" json:"targets_15_fee"`
	Targets15RequiredLevel int    `elements:"name:targets_15_required_level;type:int;size:4;text:未知;" gorm:"column:targets_15_required_level;type:integer;default:0;not null;" json:"targets_15_required_level"`
	Targets16Idtarget      int    `elements:"name:targets_16_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_16_idtarget;type:integer;default:0;not null;" json:"targets_16_idtarget"`
	Targets16Fee           int    `elements:"name:targets_16_fee;type:int;size:4;text:未知;" gorm:"column:targets_16_fee;type:integer;default:0;not null;" json:"targets_16_fee"`
	Targets16RequiredLevel int    `elements:"name:targets_16_required_level;type:int;size:4;text:未知;" gorm:"column:targets_16_required_level;type:integer;default:0;not null;" json:"targets_16_required_level"`
	Targets17Idtarget      int    `elements:"name:targets_17_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_17_idtarget;type:integer;default:0;not null;" json:"targets_17_idtarget"`
	Targets17Fee           int    `elements:"name:targets_17_fee;type:int;size:4;text:未知;" gorm:"column:targets_17_fee;type:integer;default:0;not null;" json:"targets_17_fee"`
	Targets17RequiredLevel int    `elements:"name:targets_17_required_level;type:int;size:4;text:未知;" gorm:"column:targets_17_required_level;type:integer;default:0;not null;" json:"targets_17_required_level"`
	Targets18Idtarget      int    `elements:"name:targets_18_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_18_idtarget;type:integer;default:0;not null;" json:"targets_18_idtarget"`
	Targets18Fee           int    `elements:"name:targets_18_fee;type:int;size:4;text:未知;" gorm:"column:targets_18_fee;type:integer;default:0;not null;" json:"targets_18_fee"`
	Targets18RequiredLevel int    `elements:"name:targets_18_required_level;type:int;size:4;text:未知;" gorm:"column:targets_18_required_level;type:integer;default:0;not null;" json:"targets_18_required_level"`
	Targets19Idtarget      int    `elements:"name:targets_19_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_19_idtarget;type:integer;default:0;not null;" json:"targets_19_idtarget"`
	Targets19Fee           int    `elements:"name:targets_19_fee;type:int;size:4;text:未知;" gorm:"column:targets_19_fee;type:integer;default:0;not null;" json:"targets_19_fee"`
	Targets19RequiredLevel int    `elements:"name:targets_19_required_level;type:int;size:4;text:未知;" gorm:"column:targets_19_required_level;type:integer;default:0;not null;" json:"targets_19_required_level"`
	Targets20Idtarget      int    `elements:"name:targets_20_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_20_idtarget;type:integer;default:0;not null;" json:"targets_20_idtarget"`
	Targets20Fee           int    `elements:"name:targets_20_fee;type:int;size:4;text:未知;" gorm:"column:targets_20_fee;type:integer;default:0;not null;" json:"targets_20_fee"`
	Targets20RequiredLevel int    `elements:"name:targets_20_required_level;type:int;size:4;text:未知;" gorm:"column:targets_20_required_level;type:integer;default:0;not null;" json:"targets_20_required_level"`
	Targets21Idtarget      int    `elements:"name:targets_21_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_21_idtarget;type:integer;default:0;not null;" json:"targets_21_idtarget"`
	Targets21Fee           int    `elements:"name:targets_21_fee;type:int;size:4;text:未知;" gorm:"column:targets_21_fee;type:integer;default:0;not null;" json:"targets_21_fee"`
	Targets21RequiredLevel int    `elements:"name:targets_21_required_level;type:int;size:4;text:未知;" gorm:"column:targets_21_required_level;type:integer;default:0;not null;" json:"targets_21_required_level"`
	Targets22Idtarget      int    `elements:"name:targets_22_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_22_idtarget;type:integer;default:0;not null;" json:"targets_22_idtarget"`
	Targets22Fee           int    `elements:"name:targets_22_fee;type:int;size:4;text:未知;" gorm:"column:targets_22_fee;type:integer;default:0;not null;" json:"targets_22_fee"`
	Targets22RequiredLevel int    `elements:"name:targets_22_required_level;type:int;size:4;text:未知;" gorm:"column:targets_22_required_level;type:integer;default:0;not null;" json:"targets_22_required_level"`
	Targets23Idtarget      int    `elements:"name:targets_23_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_23_idtarget;type:integer;default:0;not null;" json:"targets_23_idtarget"`
	Targets23Fee           int    `elements:"name:targets_23_fee;type:int;size:4;text:未知;" gorm:"column:targets_23_fee;type:integer;default:0;not null;" json:"targets_23_fee"`
	Targets23RequiredLevel int    `elements:"name:targets_23_required_level;type:int;size:4;text:未知;" gorm:"column:targets_23_required_level;type:integer;default:0;not null;" json:"targets_23_required_level"`
	Targets24Idtarget      int    `elements:"name:targets_24_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_24_idtarget;type:integer;default:0;not null;" json:"targets_24_idtarget"`
	Targets24Fee           int    `elements:"name:targets_24_fee;type:int;size:4;text:未知;" gorm:"column:targets_24_fee;type:integer;default:0;not null;" json:"targets_24_fee"`
	Targets24RequiredLevel int    `elements:"name:targets_24_required_level;type:int;size:4;text:未知;" gorm:"column:targets_24_required_level;type:integer;default:0;not null;" json:"targets_24_required_level"`
	Targets25Idtarget      int    `elements:"name:targets_25_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_25_idtarget;type:integer;default:0;not null;" json:"targets_25_idtarget"`
	Targets25Fee           int    `elements:"name:targets_25_fee;type:int;size:4;text:未知;" gorm:"column:targets_25_fee;type:integer;default:0;not null;" json:"targets_25_fee"`
	Targets25RequiredLevel int    `elements:"name:targets_25_required_level;type:int;size:4;text:未知;" gorm:"column:targets_25_required_level;type:integer;default:0;not null;" json:"targets_25_required_level"`
	Targets26Idtarget      int    `elements:"name:targets_26_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_26_idtarget;type:integer;default:0;not null;" json:"targets_26_idtarget"`
	Targets26Fee           int    `elements:"name:targets_26_fee;type:int;size:4;text:未知;" gorm:"column:targets_26_fee;type:integer;default:0;not null;" json:"targets_26_fee"`
	Targets26RequiredLevel int    `elements:"name:targets_26_required_level;type:int;size:4;text:未知;" gorm:"column:targets_26_required_level;type:integer;default:0;not null;" json:"targets_26_required_level"`
	Targets27Idtarget      int    `elements:"name:targets_27_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_27_idtarget;type:integer;default:0;not null;" json:"targets_27_idtarget"`
	Targets27Fee           int    `elements:"name:targets_27_fee;type:int;size:4;text:未知;" gorm:"column:targets_27_fee;type:integer;default:0;not null;" json:"targets_27_fee"`
	Targets27RequiredLevel int    `elements:"name:targets_27_required_level;type:int;size:4;text:未知;" gorm:"column:targets_27_required_level;type:integer;default:0;not null;" json:"targets_27_required_level"`
	Targets28Idtarget      int    `elements:"name:targets_28_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_28_idtarget;type:integer;default:0;not null;" json:"targets_28_idtarget"`
	Targets28Fee           int    `elements:"name:targets_28_fee;type:int;size:4;text:未知;" gorm:"column:targets_28_fee;type:integer;default:0;not null;" json:"targets_28_fee"`
	Targets28RequiredLevel int    `elements:"name:targets_28_required_level;type:int;size:4;text:未知;" gorm:"column:targets_28_required_level;type:integer;default:0;not null;" json:"targets_28_required_level"`
	Targets29Idtarget      int    `elements:"name:targets_29_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_29_idtarget;type:integer;default:0;not null;" json:"targets_29_idtarget"`
	Targets29Fee           int    `elements:"name:targets_29_fee;type:int;size:4;text:未知;" gorm:"column:targets_29_fee;type:integer;default:0;not null;" json:"targets_29_fee"`
	Targets29RequiredLevel int    `elements:"name:targets_29_required_level;type:int;size:4;text:未知;" gorm:"column:targets_29_required_level;type:integer;default:0;not null;" json:"targets_29_required_level"`
	Targets30Idtarget      int    `elements:"name:targets_30_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_30_idtarget;type:integer;default:0;not null;" json:"targets_30_idtarget"`
	Targets30Fee           int    `elements:"name:targets_30_fee;type:int;size:4;text:未知;" gorm:"column:targets_30_fee;type:integer;default:0;not null;" json:"targets_30_fee"`
	Targets30RequiredLevel int    `elements:"name:targets_30_required_level;type:int;size:4;text:未知;" gorm:"column:targets_30_required_level;type:integer;default:0;not null;" json:"targets_30_required_level"`
	Targets31Idtarget      int    `elements:"name:targets_31_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_31_idtarget;type:integer;default:0;not null;" json:"targets_31_idtarget"`
	Targets31Fee           int    `elements:"name:targets_31_fee;type:int;size:4;text:未知;" gorm:"column:targets_31_fee;type:integer;default:0;not null;" json:"targets_31_fee"`
	Targets31RequiredLevel int    `elements:"name:targets_31_required_level;type:int;size:4;text:未知;" gorm:"column:targets_31_required_level;type:integer;default:0;not null;" json:"targets_31_required_level"`
	Targets32Idtarget      int    `elements:"name:targets_32_idtarget;type:int;size:4;text:未知;" gorm:"column:targets_32_idtarget;type:integer;default:0;not null;" json:"targets_32_idtarget"`
	Targets32Fee           int    `elements:"name:targets_32_fee;type:int;size:4;text:未知;" gorm:"column:targets_32_fee;type:integer;default:0;not null;" json:"targets_32_fee"`
	Targets32RequiredLevel int    `elements:"name:targets_32_required_level;type:int;size:4;text:未知;" gorm:"column:targets_32_required_level;type:integer;default:0;not null;" json:"targets_32_required_level"`
	IdDialog               int    `elements:"name:id_dialog;type:int;size:4;text:未知;" gorm:"column:id_dialog;type:integer;default:0;not null;" json:"id_dialog"`
}

func (NpcTransmitService) TableName() string {
	return "el_npctransmitservice"
}

func (e *NpcTransmitService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcTransmitService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
