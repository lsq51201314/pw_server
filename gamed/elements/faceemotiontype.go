package elements

import (
	"pw_server/utils/structEx"
)

type FaceEmotionType struct {
	ID       int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name     string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileIcon string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
}

func (FaceEmotionType) TableName() string {
	return "el_faceemotiontype"
}

func (e *FaceEmotionType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FaceEmotionType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
