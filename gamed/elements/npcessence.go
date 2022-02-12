package elements

import (
	"pw_server/utils/structEx"
)

type NpcEssence struct {
	ID                      int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                    string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdType                  int     `elements:"name:id_type;type:int;size:4;text:未知;" gorm:"column:id_type;type:integer;default:0;not null;" json:"id_type"`
	RefreshTime             float32 `elements:"name:refresh_time;type:float;size:4;text:未知;" gorm:"column:refresh_time;type:float;default:0;not null;" json:"refresh_time"`
	AttackRule              int     `elements:"name:attack_rule;type:int;size:4;text:未知;" gorm:"column:attack_rule;type:integer;default:0;not null;" json:"attack_rule"`
	FileModel               string  `elements:"name:file_model;type:gbk;size:128;text:未知;" gorm:"column:file_model;type:varchar(128);not null;" json:"file_model"`
	TaxRate                 float32 `elements:"name:tax_rate;type:float;size:4;text:未知;" gorm:"column:tax_rate;type:float;default:0;not null;" json:"tax_rate"`
	IdSrcMonster            int     `elements:"name:id_src_monster;type:int;size:4;text:未知;" gorm:"column:id_src_monster;type:integer;default:0;not null;" json:"id_src_monster"`
	HelloMsg                string  `elements:"name:hello_msg;type:unicode;size:512;text:未知;" gorm:"column:hello_msg;type:varchar(512);not null;" json:"hello_msg"`
	IdToDiscover            int     `elements:"name:id_to_discover;type:int;size:4;text:未知;" gorm:"column:id_to_discover;type:integer;default:0;not null;" json:"id_to_discover"`
	DomainRelated           int     `elements:"name:domain_related;type:int;size:4;text:未知;" gorm:"column:domain_related;type:integer;default:0;not null;" json:"domain_related"`
	IdTalkService           int     `elements:"name:id_talk_service;type:int;size:4;text:未知;" gorm:"column:id_talk_service;type:integer;default:0;not null;" json:"id_talk_service"`
	IdSellService           int     `elements:"name:id_sell_service;type:int;size:4;text:未知;" gorm:"column:id_sell_service;type:integer;default:0;not null;" json:"id_sell_service"`
	IdBuyService            int     `elements:"name:id_buy_service;type:int;size:4;text:未知;" gorm:"column:id_buy_service;type:integer;default:0;not null;" json:"id_buy_service"`
	IdRepairService         int     `elements:"name:id_repair_service;type:int;size:4;text:未知;" gorm:"column:id_repair_service;type:integer;default:0;not null;" json:"id_repair_service"`
	IdInstallService        int     `elements:"name:id_install_service;type:int;size:4;text:未知;" gorm:"column:id_install_service;type:integer;default:0;not null;" json:"id_install_service"`
	IdUninstallService      int     `elements:"name:id_uninstall_service;type:int;size:4;text:未知;" gorm:"column:id_uninstall_service;type:integer;default:0;not null;" json:"id_uninstall_service"`
	IdTaskOutService        int     `elements:"name:id_task_out_service;type:int;size:4;text:未知;" gorm:"column:id_task_out_service;type:integer;default:0;not null;" json:"id_task_out_service"`
	IdTaskInService         int     `elements:"name:id_task_in_service;type:int;size:4;text:未知;" gorm:"column:id_task_in_service;type:integer;default:0;not null;" json:"id_task_in_service"`
	IdTaskMatterService     int     `elements:"name:id_task_matter_service;type:int;size:4;text:未知;" gorm:"column:id_task_matter_service;type:integer;default:0;not null;" json:"id_task_matter_service"`
	IdSkillService          int     `elements:"name:id_skill_service;type:int;size:4;text:未知;" gorm:"column:id_skill_service;type:integer;default:0;not null;" json:"id_skill_service"`
	IdHealService           int     `elements:"name:id_heal_service;type:int;size:4;text:未知;" gorm:"column:id_heal_service;type:integer;default:0;not null;" json:"id_heal_service"`
	IdTransmitService       int     `elements:"name:id_transmit_service;type:int;size:4;text:未知;" gorm:"column:id_transmit_service;type:integer;default:0;not null;" json:"id_transmit_service"`
	IdTransportService      int     `elements:"name:id_transport_service;type:int;size:4;text:未知;" gorm:"column:id_transport_service;type:integer;default:0;not null;" json:"id_transport_service"`
	IdProxyService          int     `elements:"name:id_proxy_service;type:int;size:4;text:未知;" gorm:"column:id_proxy_service;type:integer;default:0;not null;" json:"id_proxy_service"`
	IdStorageService        int     `elements:"name:id_storage_service;type:int;size:4;text:未知;" gorm:"column:id_storage_service;type:integer;default:0;not null;" json:"id_storage_service"`
	IdMakeService           int     `elements:"name:id_make_service;type:int;size:4;text:未知;" gorm:"column:id_make_service;type:integer;default:0;not null;" json:"id_make_service"`
	IdDecomposeService      int     `elements:"name:id_decompose_service;type:int;size:4;text:未知;" gorm:"column:id_decompose_service;type:integer;default:0;not null;" json:"id_decompose_service"`
	IdIdentifyService       int     `elements:"name:id_identify_service;type:int;size:4;text:未知;" gorm:"column:id_identify_service;type:integer;default:0;not null;" json:"id_identify_service"`
	IdWarTowerbuildService  int     `elements:"name:id_war_towerbuild_service;type:int;size:4;text:未知;" gorm:"column:id_war_towerbuild_service;type:integer;default:0;not null;" json:"id_war_towerbuild_service"`
	IdResetpropService      int     `elements:"name:id_resetprop_service;type:int;size:4;text:未知;" gorm:"column:id_resetprop_service;type:integer;default:0;not null;" json:"id_resetprop_service"`
	IdPetnameService        int     `elements:"name:id_petname_service;type:int;size:4;text:未知;" gorm:"column:id_petname_service;type:integer;default:0;not null;" json:"id_petname_service"`
	IdPetlearnskillService  int     `elements:"name:id_petlearnskill_service;type:int;size:4;text:未知;" gorm:"column:id_petlearnskill_service;type:integer;default:0;not null;" json:"id_petlearnskill_service"`
	IdPetforgetskillService int     `elements:"name:id_petforgetskill_service;type:int;size:4;text:未知;" gorm:"column:id_petforgetskill_service;type:integer;default:0;not null;" json:"id_petforgetskill_service"`
	IdEquipbindService      int     `elements:"name:id_equipbind_service;type:int;size:4;text:未知;" gorm:"column:id_equipbind_service;type:integer;default:0;not null;" json:"id_equipbind_service"`
	IdEquipdestroyService   int     `elements:"name:id_equipdestroy_service;type:int;size:4;text:未知;" gorm:"column:id_equipdestroy_service;type:integer;default:0;not null;" json:"id_equipdestroy_service"`
	IdEquipundestroyService int     `elements:"name:id_equipundestroy_service;type:int;size:4;text:未知;" gorm:"column:id_equipundestroy_service;type:integer;default:0;not null;" json:"id_equipundestroy_service"`
	CombinedServices        int     `elements:"name:combined_services;type:int;size:4;text:未知;" gorm:"column:combined_services;type:integer;default:0;not null;" json:"combined_services"`
	IdMine                  int     `elements:"name:id_mine;type:int;size:4;text:未知;" gorm:"column:id_mine;type:integer;default:0;not null;" json:"id_mine"`
}

func (NpcEssence) TableName() string {
	return "el_npcessence"
}

func (e *NpcEssence) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcEssence) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
