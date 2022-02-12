package elements

import (
	"pw_server/utils/structEx"
)

type NpcTaskInService struct {
	ID        int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name      string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdTasks1  int    `elements:"name:id_tasks_1;type:int;size:4;text:未知;" gorm:"column:id_tasks_1;type:integer;default:0;not null;" json:"id_tasks_1"`
	IdTasks2  int    `elements:"name:id_tasks_2;type:int;size:4;text:未知;" gorm:"column:id_tasks_2;type:integer;default:0;not null;" json:"id_tasks_2"`
	IdTasks3  int    `elements:"name:id_tasks_3;type:int;size:4;text:未知;" gorm:"column:id_tasks_3;type:integer;default:0;not null;" json:"id_tasks_3"`
	IdTasks4  int    `elements:"name:id_tasks_4;type:int;size:4;text:未知;" gorm:"column:id_tasks_4;type:integer;default:0;not null;" json:"id_tasks_4"`
	IdTasks5  int    `elements:"name:id_tasks_5;type:int;size:4;text:未知;" gorm:"column:id_tasks_5;type:integer;default:0;not null;" json:"id_tasks_5"`
	IdTasks6  int    `elements:"name:id_tasks_6;type:int;size:4;text:未知;" gorm:"column:id_tasks_6;type:integer;default:0;not null;" json:"id_tasks_6"`
	IdTasks7  int    `elements:"name:id_tasks_7;type:int;size:4;text:未知;" gorm:"column:id_tasks_7;type:integer;default:0;not null;" json:"id_tasks_7"`
	IdTasks8  int    `elements:"name:id_tasks_8;type:int;size:4;text:未知;" gorm:"column:id_tasks_8;type:integer;default:0;not null;" json:"id_tasks_8"`
	IdTasks9  int    `elements:"name:id_tasks_9;type:int;size:4;text:未知;" gorm:"column:id_tasks_9;type:integer;default:0;not null;" json:"id_tasks_9"`
	IdTasks10 int    `elements:"name:id_tasks_10;type:int;size:4;text:未知;" gorm:"column:id_tasks_10;type:integer;default:0;not null;" json:"id_tasks_10"`
	IdTasks11 int    `elements:"name:id_tasks_11;type:int;size:4;text:未知;" gorm:"column:id_tasks_11;type:integer;default:0;not null;" json:"id_tasks_11"`
	IdTasks12 int    `elements:"name:id_tasks_12;type:int;size:4;text:未知;" gorm:"column:id_tasks_12;type:integer;default:0;not null;" json:"id_tasks_12"`
	IdTasks13 int    `elements:"name:id_tasks_13;type:int;size:4;text:未知;" gorm:"column:id_tasks_13;type:integer;default:0;not null;" json:"id_tasks_13"`
	IdTasks14 int    `elements:"name:id_tasks_14;type:int;size:4;text:未知;" gorm:"column:id_tasks_14;type:integer;default:0;not null;" json:"id_tasks_14"`
	IdTasks15 int    `elements:"name:id_tasks_15;type:int;size:4;text:未知;" gorm:"column:id_tasks_15;type:integer;default:0;not null;" json:"id_tasks_15"`
	IdTasks16 int    `elements:"name:id_tasks_16;type:int;size:4;text:未知;" gorm:"column:id_tasks_16;type:integer;default:0;not null;" json:"id_tasks_16"`
	IdTasks17 int    `elements:"name:id_tasks_17;type:int;size:4;text:未知;" gorm:"column:id_tasks_17;type:integer;default:0;not null;" json:"id_tasks_17"`
	IdTasks18 int    `elements:"name:id_tasks_18;type:int;size:4;text:未知;" gorm:"column:id_tasks_18;type:integer;default:0;not null;" json:"id_tasks_18"`
	IdTasks19 int    `elements:"name:id_tasks_19;type:int;size:4;text:未知;" gorm:"column:id_tasks_19;type:integer;default:0;not null;" json:"id_tasks_19"`
	IdTasks20 int    `elements:"name:id_tasks_20;type:int;size:4;text:未知;" gorm:"column:id_tasks_20;type:integer;default:0;not null;" json:"id_tasks_20"`
	IdTasks21 int    `elements:"name:id_tasks_21;type:int;size:4;text:未知;" gorm:"column:id_tasks_21;type:integer;default:0;not null;" json:"id_tasks_21"`
	IdTasks22 int    `elements:"name:id_tasks_22;type:int;size:4;text:未知;" gorm:"column:id_tasks_22;type:integer;default:0;not null;" json:"id_tasks_22"`
	IdTasks23 int    `elements:"name:id_tasks_23;type:int;size:4;text:未知;" gorm:"column:id_tasks_23;type:integer;default:0;not null;" json:"id_tasks_23"`
	IdTasks24 int    `elements:"name:id_tasks_24;type:int;size:4;text:未知;" gorm:"column:id_tasks_24;type:integer;default:0;not null;" json:"id_tasks_24"`
	IdTasks25 int    `elements:"name:id_tasks_25;type:int;size:4;text:未知;" gorm:"column:id_tasks_25;type:integer;default:0;not null;" json:"id_tasks_25"`
	IdTasks26 int    `elements:"name:id_tasks_26;type:int;size:4;text:未知;" gorm:"column:id_tasks_26;type:integer;default:0;not null;" json:"id_tasks_26"`
	IdTasks27 int    `elements:"name:id_tasks_27;type:int;size:4;text:未知;" gorm:"column:id_tasks_27;type:integer;default:0;not null;" json:"id_tasks_27"`
	IdTasks28 int    `elements:"name:id_tasks_28;type:int;size:4;text:未知;" gorm:"column:id_tasks_28;type:integer;default:0;not null;" json:"id_tasks_28"`
	IdTasks29 int    `elements:"name:id_tasks_29;type:int;size:4;text:未知;" gorm:"column:id_tasks_29;type:integer;default:0;not null;" json:"id_tasks_29"`
	IdTasks30 int    `elements:"name:id_tasks_30;type:int;size:4;text:未知;" gorm:"column:id_tasks_30;type:integer;default:0;not null;" json:"id_tasks_30"`
	IdTasks31 int    `elements:"name:id_tasks_31;type:int;size:4;text:未知;" gorm:"column:id_tasks_31;type:integer;default:0;not null;" json:"id_tasks_31"`
	IdTasks32 int    `elements:"name:id_tasks_32;type:int;size:4;text:未知;" gorm:"column:id_tasks_32;type:integer;default:0;not null;" json:"id_tasks_32"`
}

func (NpcTaskInService) TableName() string {
	return "el_npctaskinservice"
}

func (e *NpcTaskInService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcTaskInService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
