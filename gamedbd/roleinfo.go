package gamedbd

type RoleInfo struct {
	ID         int64  `gorm:"column:id;type:int;autoIncrement:false;uniqueIndex;not null;"`
	UserID     int    `gorm:"column:user_id;type:integer;default:0;not null;"`
	RoleID     int    `gorm:"column:role_id;type:integer;default:0;not null;"`
	Gender     int    `gorm:"column:gender;type:integer;default:0;not null;"`
	Race       int    `gorm:"column:race;type:integer;default:0;not null;"`
	Level      int    `gorm:"column:level;type:integer;default:0;not null;"`
	Level2     int    `gorm:"column:level2;type:integer;default:0;not null;"`
	Name       string `gorm:"column:name;type:varchar(50);not null;"`
	CustomData string `gorm:"column:custom_data;type:varchar(512);not null;"`
}

func (RoleInfo) TableName() string {
	return "db_roles"
}
