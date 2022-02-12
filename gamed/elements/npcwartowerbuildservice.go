package elements

import (
	"pw_server/utils/structEx"
)

type NpcWarTowerbuildService struct {
	ID                     int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                   string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	BuildInfo1IdInBuild    int    `elements:"name:build_info_1_id_in_build;type:int;size:4;text:未知;" gorm:"column:build_info_1_id_in_build;type:integer;default:0;not null;" json:"build_info_1_id_in_build"`
	BuildInfo1IdBuildup    int    `elements:"name:build_info_1_id_buildup;type:int;size:4;text:未知;" gorm:"column:build_info_1_id_buildup;type:integer;default:0;not null;" json:"build_info_1_id_buildup"`
	BuildInfo1IdObjectNeed int    `elements:"name:build_info_1_id_object_need;type:int;size:4;text:未知;" gorm:"column:build_info_1_id_object_need;type:integer;default:0;not null;" json:"build_info_1_id_object_need"`
	BuildInfo1TimeUse      int    `elements:"name:build_info_1_time_use;type:int;size:4;text:未知;" gorm:"column:build_info_1_time_use;type:integer;default:0;not null;" json:"build_info_1_time_use"`
	BuildInfo1Fee          int    `elements:"name:build_info_1_fee;type:int;size:4;text:未知;" gorm:"column:build_info_1_fee;type:integer;default:0;not null;" json:"build_info_1_fee"`
	BuildInfo2IdInBuild    int    `elements:"name:build_info_2_id_in_build;type:int;size:4;text:未知;" gorm:"column:build_info_2_id_in_build;type:integer;default:0;not null;" json:"build_info_2_id_in_build"`
	BuildInfo2IdBuildup    int    `elements:"name:build_info_2_id_buildup;type:int;size:4;text:未知;" gorm:"column:build_info_2_id_buildup;type:integer;default:0;not null;" json:"build_info_2_id_buildup"`
	BuildInfo2IdObjectNeed int    `elements:"name:build_info_2_id_object_need;type:int;size:4;text:未知;" gorm:"column:build_info_2_id_object_need;type:integer;default:0;not null;" json:"build_info_2_id_object_need"`
	BuildInfo2TimeUse      int    `elements:"name:build_info_2_time_use;type:int;size:4;text:未知;" gorm:"column:build_info_2_time_use;type:integer;default:0;not null;" json:"build_info_2_time_use"`
	BuildInfo2Fee          int    `elements:"name:build_info_2_fee;type:int;size:4;text:未知;" gorm:"column:build_info_2_fee;type:integer;default:0;not null;" json:"build_info_2_fee"`
	BuildInfo3IdInBuild    int    `elements:"name:build_info_3_id_in_build;type:int;size:4;text:未知;" gorm:"column:build_info_3_id_in_build;type:integer;default:0;not null;" json:"build_info_3_id_in_build"`
	BuildInfo3IdBuildup    int    `elements:"name:build_info_3_id_buildup;type:int;size:4;text:未知;" gorm:"column:build_info_3_id_buildup;type:integer;default:0;not null;" json:"build_info_3_id_buildup"`
	BuildInfo3IdObjectNeed int    `elements:"name:build_info_3_id_object_need;type:int;size:4;text:未知;" gorm:"column:build_info_3_id_object_need;type:integer;default:0;not null;" json:"build_info_3_id_object_need"`
	BuildInfo3TimeUse      int    `elements:"name:build_info_3_time_use;type:int;size:4;text:未知;" gorm:"column:build_info_3_time_use;type:integer;default:0;not null;" json:"build_info_3_time_use"`
	BuildInfo3Fee          int    `elements:"name:build_info_3_fee;type:int;size:4;text:未知;" gorm:"column:build_info_3_fee;type:integer;default:0;not null;" json:"build_info_3_fee"`
	BuildInfo4IdInBuild    int    `elements:"name:build_info_4_id_in_build;type:int;size:4;text:未知;" gorm:"column:build_info_4_id_in_build;type:integer;default:0;not null;" json:"build_info_4_id_in_build"`
	BuildInfo4IdBuildup    int    `elements:"name:build_info_4_id_buildup;type:int;size:4;text:未知;" gorm:"column:build_info_4_id_buildup;type:integer;default:0;not null;" json:"build_info_4_id_buildup"`
	BuildInfo4IdObjectNeed int    `elements:"name:build_info_4_id_object_need;type:int;size:4;text:未知;" gorm:"column:build_info_4_id_object_need;type:integer;default:0;not null;" json:"build_info_4_id_object_need"`
	BuildInfo4TimeUse      int    `elements:"name:build_info_4_time_use;type:int;size:4;text:未知;" gorm:"column:build_info_4_time_use;type:integer;default:0;not null;" json:"build_info_4_time_use"`
	BuildInfo4Fee          int    `elements:"name:build_info_4_fee;type:int;size:4;text:未知;" gorm:"column:build_info_4_fee;type:integer;default:0;not null;" json:"build_info_4_fee"`
}

func (NpcWarTowerbuildService) TableName() string {
	return "el_npcwartowerbuildservice"
}

func (e *NpcWarTowerbuildService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcWarTowerbuildService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
