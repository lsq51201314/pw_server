package elements

import (
	"pw_server/utils/structEx"
)

type NpcResetpropService struct {
	ID                       int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                     string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	PropEntry1IdObjectNeed   int    `elements:"name:prop_entry_1_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_1_id_object_need;type:integer;default:0;not null;" json:"prop_entry_1_id_object_need"`
	PropEntry1StrengthDelta  int    `elements:"name:prop_entry_1_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_1_strength_delta;type:integer;default:0;not null;" json:"prop_entry_1_strength_delta"`
	PropEntry1AgilityDelta   int    `elements:"name:prop_entry_1_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_1_agility_delta;type:integer;default:0;not null;" json:"prop_entry_1_agility_delta"`
	PropEntry1VitalDelta     int    `elements:"name:prop_entry_1_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_1_vital_delta;type:integer;default:0;not null;" json:"prop_entry_1_vital_delta"`
	PropEntry1EnergyDelta    int    `elements:"name:prop_entry_1_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_1_energy_delta;type:integer;default:0;not null;" json:"prop_entry_1_energy_delta"`
	PropEntry2IdObjectNeed   int    `elements:"name:prop_entry_2_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_2_id_object_need;type:integer;default:0;not null;" json:"prop_entry_2_id_object_need"`
	PropEntry2StrengthDelta  int    `elements:"name:prop_entry_2_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_2_strength_delta;type:integer;default:0;not null;" json:"prop_entry_2_strength_delta"`
	PropEntry2AgilityDelta   int    `elements:"name:prop_entry_2_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_2_agility_delta;type:integer;default:0;not null;" json:"prop_entry_2_agility_delta"`
	PropEntry2VitalDelta     int    `elements:"name:prop_entry_2_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_2_vital_delta;type:integer;default:0;not null;" json:"prop_entry_2_vital_delta"`
	PropEntry2EnergyDelta    int    `elements:"name:prop_entry_2_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_2_energy_delta;type:integer;default:0;not null;" json:"prop_entry_2_energy_delta"`
	PropEntry3IdObjectNeed   int    `elements:"name:prop_entry_3_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_3_id_object_need;type:integer;default:0;not null;" json:"prop_entry_3_id_object_need"`
	PropEntry3StrengthDelta  int    `elements:"name:prop_entry_3_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_3_strength_delta;type:integer;default:0;not null;" json:"prop_entry_3_strength_delta"`
	PropEntry3AgilityDelta   int    `elements:"name:prop_entry_3_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_3_agility_delta;type:integer;default:0;not null;" json:"prop_entry_3_agility_delta"`
	PropEntry3VitalDelta     int    `elements:"name:prop_entry_3_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_3_vital_delta;type:integer;default:0;not null;" json:"prop_entry_3_vital_delta"`
	PropEntry3EnergyDelta    int    `elements:"name:prop_entry_3_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_3_energy_delta;type:integer;default:0;not null;" json:"prop_entry_3_energy_delta"`
	PropEntry4IdObjectNeed   int    `elements:"name:prop_entry_4_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_4_id_object_need;type:integer;default:0;not null;" json:"prop_entry_4_id_object_need"`
	PropEntry4StrengthDelta  int    `elements:"name:prop_entry_4_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_4_strength_delta;type:integer;default:0;not null;" json:"prop_entry_4_strength_delta"`
	PropEntry4AgilityDelta   int    `elements:"name:prop_entry_4_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_4_agility_delta;type:integer;default:0;not null;" json:"prop_entry_4_agility_delta"`
	PropEntry4VitalDelta     int    `elements:"name:prop_entry_4_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_4_vital_delta;type:integer;default:0;not null;" json:"prop_entry_4_vital_delta"`
	PropEntry4EnergyDelta    int    `elements:"name:prop_entry_4_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_4_energy_delta;type:integer;default:0;not null;" json:"prop_entry_4_energy_delta"`
	PropEntry5IdObjectNeed   int    `elements:"name:prop_entry_5_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_5_id_object_need;type:integer;default:0;not null;" json:"prop_entry_5_id_object_need"`
	PropEntry5StrengthDelta  int    `elements:"name:prop_entry_5_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_5_strength_delta;type:integer;default:0;not null;" json:"prop_entry_5_strength_delta"`
	PropEntry5AgilityDelta   int    `elements:"name:prop_entry_5_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_5_agility_delta;type:integer;default:0;not null;" json:"prop_entry_5_agility_delta"`
	PropEntry5VitalDelta     int    `elements:"name:prop_entry_5_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_5_vital_delta;type:integer;default:0;not null;" json:"prop_entry_5_vital_delta"`
	PropEntry5EnergyDelta    int    `elements:"name:prop_entry_5_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_5_energy_delta;type:integer;default:0;not null;" json:"prop_entry_5_energy_delta"`
	PropEntry6IdObjectNeed   int    `elements:"name:prop_entry_6_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_6_id_object_need;type:integer;default:0;not null;" json:"prop_entry_6_id_object_need"`
	PropEntry6StrengthDelta  int    `elements:"name:prop_entry_6_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_6_strength_delta;type:integer;default:0;not null;" json:"prop_entry_6_strength_delta"`
	PropEntry6AgilityDelta   int    `elements:"name:prop_entry_6_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_6_agility_delta;type:integer;default:0;not null;" json:"prop_entry_6_agility_delta"`
	PropEntry6VitalDelta     int    `elements:"name:prop_entry_6_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_6_vital_delta;type:integer;default:0;not null;" json:"prop_entry_6_vital_delta"`
	PropEntry6EnergyDelta    int    `elements:"name:prop_entry_6_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_6_energy_delta;type:integer;default:0;not null;" json:"prop_entry_6_energy_delta"`
	PropEntry7IdObjectNeed   int    `elements:"name:prop_entry_7_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_7_id_object_need;type:integer;default:0;not null;" json:"prop_entry_7_id_object_need"`
	PropEntry7StrengthDelta  int    `elements:"name:prop_entry_7_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_7_strength_delta;type:integer;default:0;not null;" json:"prop_entry_7_strength_delta"`
	PropEntry7AgilityDelta   int    `elements:"name:prop_entry_7_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_7_agility_delta;type:integer;default:0;not null;" json:"prop_entry_7_agility_delta"`
	PropEntry7VitalDelta     int    `elements:"name:prop_entry_7_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_7_vital_delta;type:integer;default:0;not null;" json:"prop_entry_7_vital_delta"`
	PropEntry7EnergyDelta    int    `elements:"name:prop_entry_7_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_7_energy_delta;type:integer;default:0;not null;" json:"prop_entry_7_energy_delta"`
	PropEntry8IdObjectNeed   int    `elements:"name:prop_entry_8_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_8_id_object_need;type:integer;default:0;not null;" json:"prop_entry_8_id_object_need"`
	PropEntry8StrengthDelta  int    `elements:"name:prop_entry_8_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_8_strength_delta;type:integer;default:0;not null;" json:"prop_entry_8_strength_delta"`
	PropEntry8AgilityDelta   int    `elements:"name:prop_entry_8_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_8_agility_delta;type:integer;default:0;not null;" json:"prop_entry_8_agility_delta"`
	PropEntry8VitalDelta     int    `elements:"name:prop_entry_8_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_8_vital_delta;type:integer;default:0;not null;" json:"prop_entry_8_vital_delta"`
	PropEntry8EnergyDelta    int    `elements:"name:prop_entry_8_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_8_energy_delta;type:integer;default:0;not null;" json:"prop_entry_8_energy_delta"`
	PropEntry9IdObjectNeed   int    `elements:"name:prop_entry_9_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_9_id_object_need;type:integer;default:0;not null;" json:"prop_entry_9_id_object_need"`
	PropEntry9StrengthDelta  int    `elements:"name:prop_entry_9_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_9_strength_delta;type:integer;default:0;not null;" json:"prop_entry_9_strength_delta"`
	PropEntry9AgilityDelta   int    `elements:"name:prop_entry_9_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_9_agility_delta;type:integer;default:0;not null;" json:"prop_entry_9_agility_delta"`
	PropEntry9VitalDelta     int    `elements:"name:prop_entry_9_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_9_vital_delta;type:integer;default:0;not null;" json:"prop_entry_9_vital_delta"`
	PropEntry9EnergyDelta    int    `elements:"name:prop_entry_9_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_9_energy_delta;type:integer;default:0;not null;" json:"prop_entry_9_energy_delta"`
	PropEntry10IdObjectNeed  int    `elements:"name:prop_entry_10_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_10_id_object_need;type:integer;default:0;not null;" json:"prop_entry_10_id_object_need"`
	PropEntry10StrengthDelta int    `elements:"name:prop_entry_10_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_10_strength_delta;type:integer;default:0;not null;" json:"prop_entry_10_strength_delta"`
	PropEntry10AgilityDelta  int    `elements:"name:prop_entry_10_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_10_agility_delta;type:integer;default:0;not null;" json:"prop_entry_10_agility_delta"`
	PropEntry10VitalDelta    int    `elements:"name:prop_entry_10_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_10_vital_delta;type:integer;default:0;not null;" json:"prop_entry_10_vital_delta"`
	PropEntry10EnergyDelta   int    `elements:"name:prop_entry_10_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_10_energy_delta;type:integer;default:0;not null;" json:"prop_entry_10_energy_delta"`
	PropEntry11IdObjectNeed  int    `elements:"name:prop_entry_11_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_11_id_object_need;type:integer;default:0;not null;" json:"prop_entry_11_id_object_need"`
	PropEntry11StrengthDelta int    `elements:"name:prop_entry_11_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_11_strength_delta;type:integer;default:0;not null;" json:"prop_entry_11_strength_delta"`
	PropEntry11AgilityDelta  int    `elements:"name:prop_entry_11_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_11_agility_delta;type:integer;default:0;not null;" json:"prop_entry_11_agility_delta"`
	PropEntry11VitalDelta    int    `elements:"name:prop_entry_11_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_11_vital_delta;type:integer;default:0;not null;" json:"prop_entry_11_vital_delta"`
	PropEntry11EnergyDelta   int    `elements:"name:prop_entry_11_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_11_energy_delta;type:integer;default:0;not null;" json:"prop_entry_11_energy_delta"`
	PropEntry12IdObjectNeed  int    `elements:"name:prop_entry_12_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_12_id_object_need;type:integer;default:0;not null;" json:"prop_entry_12_id_object_need"`
	PropEntry12StrengthDelta int    `elements:"name:prop_entry_12_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_12_strength_delta;type:integer;default:0;not null;" json:"prop_entry_12_strength_delta"`
	PropEntry12AgilityDelta  int    `elements:"name:prop_entry_12_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_12_agility_delta;type:integer;default:0;not null;" json:"prop_entry_12_agility_delta"`
	PropEntry12VitalDelta    int    `elements:"name:prop_entry_12_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_12_vital_delta;type:integer;default:0;not null;" json:"prop_entry_12_vital_delta"`
	PropEntry12EnergyDelta   int    `elements:"name:prop_entry_12_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_12_energy_delta;type:integer;default:0;not null;" json:"prop_entry_12_energy_delta"`
	PropEntry13IdObjectNeed  int    `elements:"name:prop_entry_13_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_13_id_object_need;type:integer;default:0;not null;" json:"prop_entry_13_id_object_need"`
	PropEntry13StrengthDelta int    `elements:"name:prop_entry_13_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_13_strength_delta;type:integer;default:0;not null;" json:"prop_entry_13_strength_delta"`
	PropEntry13AgilityDelta  int    `elements:"name:prop_entry_13_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_13_agility_delta;type:integer;default:0;not null;" json:"prop_entry_13_agility_delta"`
	PropEntry13VitalDelta    int    `elements:"name:prop_entry_13_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_13_vital_delta;type:integer;default:0;not null;" json:"prop_entry_13_vital_delta"`
	PropEntry13EnergyDelta   int    `elements:"name:prop_entry_13_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_13_energy_delta;type:integer;default:0;not null;" json:"prop_entry_13_energy_delta"`
	PropEntry14IdObjectNeed  int    `elements:"name:prop_entry_14_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_14_id_object_need;type:integer;default:0;not null;" json:"prop_entry_14_id_object_need"`
	PropEntry14StrengthDelta int    `elements:"name:prop_entry_14_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_14_strength_delta;type:integer;default:0;not null;" json:"prop_entry_14_strength_delta"`
	PropEntry14AgilityDelta  int    `elements:"name:prop_entry_14_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_14_agility_delta;type:integer;default:0;not null;" json:"prop_entry_14_agility_delta"`
	PropEntry14VitalDelta    int    `elements:"name:prop_entry_14_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_14_vital_delta;type:integer;default:0;not null;" json:"prop_entry_14_vital_delta"`
	PropEntry14EnergyDelta   int    `elements:"name:prop_entry_14_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_14_energy_delta;type:integer;default:0;not null;" json:"prop_entry_14_energy_delta"`
	PropEntry15IdObjectNeed  int    `elements:"name:prop_entry_15_id_object_need;type:int;size:4;text:未知;" gorm:"column:prop_entry_15_id_object_need;type:integer;default:0;not null;" json:"prop_entry_15_id_object_need"`
	PropEntry15StrengthDelta int    `elements:"name:prop_entry_15_strength_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_15_strength_delta;type:integer;default:0;not null;" json:"prop_entry_15_strength_delta"`
	PropEntry15AgilityDelta  int    `elements:"name:prop_entry_15_agility_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_15_agility_delta;type:integer;default:0;not null;" json:"prop_entry_15_agility_delta"`
	PropEntry15VitalDelta    int    `elements:"name:prop_entry_15_vital_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_15_vital_delta;type:integer;default:0;not null;" json:"prop_entry_15_vital_delta"`
	PropEntry15EnergyDelta   int    `elements:"name:prop_entry_15_energy_delta;type:int;size:4;text:未知;" gorm:"column:prop_entry_15_energy_delta;type:integer;default:0;not null;" json:"prop_entry_15_energy_delta"`
}

func (NpcResetpropService) TableName() string {
	return "el_npcresetpropservice"
}

func (e *NpcResetpropService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcResetpropService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
