package authd

import (
	"fmt"
	"log"
	"os"
	"pw_server/gamed/elements"
	"pw_server/gamedbd"
	"pw_server/global"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MDB *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.RunCfg.MySQLCfg.User,
		global.RunCfg.MySQLCfg.Passwd,
		global.RunCfg.MySQLCfg.Host,
		global.RunCfg.MySQLCfg.Port,
		global.RunCfg.MySQLCfg.DbName,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	if MDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		zap.L().Error("无法连接到MySQL数据库。", zap.Error(err))
		return
	}
	//连接池
	sqlDB, _ := MDB.DB()
	sqlDB.SetMaxIdleConns(global.RunCfg.MySQLCfg.MaxIdle)
	sqlDB.SetMaxOpenConns(global.RunCfg.MySQLCfg.MaxOpen)
	// 自动迁移
	autoMigrate()
	return
}

func autoMigrate() {
	MDB.AutoMigrate(
		&Users{},
		&gamedbd.RoleInfo{},
		&elements.EquipmentAddon{},
		&elements.WeaponMajorType{},
		&elements.WeaponSubType{},
		&elements.WeaponEssence{},
		&elements.ArmorMajorType{},
		&elements.ArmorSubType{},
		&elements.ArmorEssence{},
		&elements.DecorationMajorType{},
		&elements.DecorationSubType{},
		&elements.DecorationEssence{},
		&elements.MedicineMajorType{},
		&elements.MedicineSubType{},
		&elements.MedicineEssence{},
		&elements.MaterialMajorType{},
		&elements.MaterialSubType{},
		&elements.MaterialEssence{},
		&elements.DamageruneSubType{},
		&elements.DamageruneEssence{},
		&elements.ArmorruneSubType{},
		&elements.ArmorruneEssence{},
		&elements.SkilltomeSubType{},
		&elements.SkilltomeEssence{},
		&elements.FlyswordEssence{},
		&elements.WingmanwingEssence{},
		&elements.TownscrollEssence{},
		&elements.UnionscrollEssence{},
		&elements.RevivescrollEssence{},
		&elements.ElementEssence{},
		&elements.TaskmatterEssence{},
		&elements.TossmatterEssence{},
		&elements.ProjectileType{},
		&elements.ProjectileEssence{},
		&elements.QuiverSubType{},
		&elements.QuiverEssence{},
		&elements.StoneSubType{},
		&elements.StoneEssence{},
		&elements.MonsterAddon{},
		&elements.MonsterType{},
		&elements.MonsterEssence{},
		&elements.NpcTalkService{},
		&elements.NpcSellService{},
		&elements.NpcBuyService{},
		&elements.NpcRepairService{},
		&elements.NpcInstallService{},
		&elements.NpcUninstallService{},
		&elements.NpcTaskInService{},
		&elements.NpcTaskOutService{},
		&elements.NpcTaskMatterService{},
		&elements.NpcSkillService{},
		&elements.NpcHealService{},
		&elements.NpcTransmitService{},
		&elements.NpcTransportService{},
		&elements.NpcProxyService{},
		&elements.NpcStorageService{},
		&elements.NpcMakeService{},
		&elements.NpcDecomposeService{},
		&elements.NpcType{},
		&elements.NpcEssence{},
		&elements.TalkProc{},
		&elements.FaceTextureEssence{},
		&elements.FaceShapeEssence{},
		&elements.FaceEmotionType{},
		&elements.FaceExpressionEssence{},
		&elements.FaceHairEssence{},
		&elements.FaceMoustacheEssence{},
		&elements.ColorpickerEssence{},
		&elements.CustomizedataEssence{},
		&elements.RecipeMajorType{},
		&elements.RecipeSubType{},
		&elements.RecipeEssence{},
		&elements.EnemyFactionConfig{},
		&elements.CharracterClassConfig{},
		&elements.ParamAdjustConfig{},
		&elements.PlayerActionInfoConfig{},
		&elements.TaskdiceEssence{},
		&elements.TasknormalmatterEssence{},
		&elements.FaceFalingEssence{},
		&elements.PlayerLevelexpConfig{},
		&elements.MineType{},
		&elements.MineEssence{},
		&elements.NpcIdentifyService{},
		&elements.FashionMajorType{},
		&elements.FashionSubType{},
		&elements.FashionEssence{},
		&elements.FaceticketMajorType{},
		&elements.FaceticketSubType{},
		&elements.FaceticketEssence{},
		&elements.FacepillMajorType{},
		&elements.FacepillSubType{},
		&elements.FacepillEssence{},
		&elements.SuiteEssence{},
		&elements.GmGeneratorType{},
		&elements.GmGeneratorEssence{},
		&elements.PetType{},
		&elements.PetEssence{},
		&elements.PetEggEssence{},
		&elements.PetFoodEssence{},
		&elements.PetFaceticketEssence{},
		&elements.FireworksEssence{},
		&elements.WarTankcallinEssence{},
		&elements.NpcWarTowerbuildService{},
		&elements.PlayerSecondlevelConfig{},
		&elements.NpcResetpropService{},
		&elements.NpcPetnameService{},
		&elements.NpcPetlearnskillService{},
		&elements.NpcPetforgetskillService{},
		&elements.SkillmatterEssence{},
		&elements.RefineTicketEssence{},
		&elements.DestroyingEssence{},
		&elements.NpcEquipbindService{},
		&elements.NpcEquipdestroyService{},
		&elements.NpcEquipundestroyService{},
		&elements.BibleEssence{},
		&elements.SpeakerEssence{},
		&elements.AutohpEssence{},
		&elements.AutompEssence{},
		&elements.DoubleExpEssence{},
	)
}
