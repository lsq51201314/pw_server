package elements

import (
	"pw_server/utils/structEx"
)

type TaskdiceEssence struct {
	ID                    int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                  string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileMatter            string  `elements:"name:file_matter;type:gbk;size:128;text:未知;" gorm:"column:file_matter;type:varchar(128);not null;" json:"file_matter"`
	FileIcon              string  `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	TaskLists1Id          int     `elements:"name:task_lists_1_id;type:int;size:4;text:未知;" gorm:"column:task_lists_1_id;type:integer;default:0;not null;" json:"task_lists_1_id"`
	TaskLists1Probability float32 `elements:"name:task_lists_1_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_1_probability;type:float;default:0;not null;" json:"task_lists_1_probability"`
	TaskLists2Id          int     `elements:"name:task_lists_2_id;type:int;size:4;text:未知;" gorm:"column:task_lists_2_id;type:integer;default:0;not null;" json:"task_lists_2_id"`
	TaskLists2Probability float32 `elements:"name:task_lists_2_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_2_probability;type:float;default:0;not null;" json:"task_lists_2_probability"`
	TaskLists3Id          int     `elements:"name:task_lists_3_id;type:int;size:4;text:未知;" gorm:"column:task_lists_3_id;type:integer;default:0;not null;" json:"task_lists_3_id"`
	TaskLists3Probability float32 `elements:"name:task_lists_3_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_3_probability;type:float;default:0;not null;" json:"task_lists_3_probability"`
	TaskLists4Id          int     `elements:"name:task_lists_4_id;type:int;size:4;text:未知;" gorm:"column:task_lists_4_id;type:integer;default:0;not null;" json:"task_lists_4_id"`
	TaskLists4Probability float32 `elements:"name:task_lists_4_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_4_probability;type:float;default:0;not null;" json:"task_lists_4_probability"`
	TaskLists5Id          int     `elements:"name:task_lists_5_id;type:int;size:4;text:未知;" gorm:"column:task_lists_5_id;type:integer;default:0;not null;" json:"task_lists_5_id"`
	TaskLists5Probability float32 `elements:"name:task_lists_5_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_5_probability;type:float;default:0;not null;" json:"task_lists_5_probability"`
	TaskLists6Id          int     `elements:"name:task_lists_6_id;type:int;size:4;text:未知;" gorm:"column:task_lists_6_id;type:integer;default:0;not null;" json:"task_lists_6_id"`
	TaskLists6Probability float32 `elements:"name:task_lists_6_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_6_probability;type:float;default:0;not null;" json:"task_lists_6_probability"`
	TaskLists7Id          int     `elements:"name:task_lists_7_id;type:int;size:4;text:未知;" gorm:"column:task_lists_7_id;type:integer;default:0;not null;" json:"task_lists_7_id"`
	TaskLists7Probability float32 `elements:"name:task_lists_7_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_7_probability;type:float;default:0;not null;" json:"task_lists_7_probability"`
	TaskLists8Id          int     `elements:"name:task_lists_8_id;type:int;size:4;text:未知;" gorm:"column:task_lists_8_id;type:integer;default:0;not null;" json:"task_lists_8_id"`
	TaskLists8Probability float32 `elements:"name:task_lists_8_probability;type:float;size:4;text:未知;" gorm:"column:task_lists_8_probability;type:float;default:0;not null;" json:"task_lists_8_probability"`
	UseOnPick             int     `elements:"name:use_on_pick;type:int;size:4;text:未知;" gorm:"column:use_on_pick;type:integer;default:0;not null;" json:"use_on_pick"`
	PileNumMax            int     `elements:"name:pile_num_max;type:int;size:4;text:未知;" gorm:"column:pile_num_max;type:integer;default:0;not null;" json:"pile_num_max"`
	HasGuid               int     `elements:"name:has_guid;type:int;size:4;text:未知;" gorm:"column:has_guid;type:integer;default:0;not null;" json:"has_guid"`
	ProcType              int     `elements:"name:proc_type;type:int;size:4;text:未知;" gorm:"column:proc_type;type:integer;default:0;not null;" json:"proc_type"`
}

func (TaskdiceEssence) TableName() string {
	return "el_taskdiceessence"
}

func (e *TaskdiceEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *TaskdiceEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
