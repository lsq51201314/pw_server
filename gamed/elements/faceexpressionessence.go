package elements

import (
	"pw_server/utils/structEx"
)

type FaceExpressionEssence struct {
	ID               int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name             string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	FileExpression   string `elements:"name:file_expression;type:gbk;size:128;text:未知;" gorm:"column:file_expression;type:varchar(128);not null;" json:"file_expression"`
	FileIcon         string `elements:"name:file_icon;type:gbk;size:128;text:未知;" gorm:"column:file_icon;type:varchar(128);not null;" json:"file_icon"`
	CharacterComboId int    `elements:"name:character_combo_id;type:int;size:4;text:未知;" gorm:"column:character_combo_id;type:integer;default:0;not null;" json:"character_combo_id"`
	GenderId         int    `elements:"name:gender_id;type:int;size:4;text:未知;" gorm:"column:gender_id;type:integer;default:0;not null;" json:"gender_id"`
	EmotionId        int    `elements:"name:emotion_id;type:int;size:4;text:未知;" gorm:"column:emotion_id;type:integer;default:0;not null;" json:"emotion_id"`
}

func (FaceExpressionEssence) TableName() string {
	return "el_faceexpressionessence"
}

func (e *FaceExpressionEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *FaceExpressionEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
