package authd

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         int       `gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;"`
	UserName   string    `gorm:"column:username;type:varchar(50);unique;not null;"`
	Password   string    `gorm:"column:password;type:varchar(50);not null;"`
	CreateTime time.Time `gorm:"autoCreateTime;"`
	UpdateTime time.Time `gorm:"autoUpdateTime;"`
	DeleteTime gorm.DeletedAt
}
