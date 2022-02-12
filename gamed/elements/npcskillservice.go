package elements

import (
	"pw_server/utils/structEx"
)

type NpcSkillService struct {
	ID          int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name        string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdSkills1   int    `elements:"name:id_skills_1;type:int;size:4;text:未知;" gorm:"column:id_skills_1;type:integer;default:0;not null;" json:"id_skills_1"`
	IdSkills2   int    `elements:"name:id_skills_2;type:int;size:4;text:未知;" gorm:"column:id_skills_2;type:integer;default:0;not null;" json:"id_skills_2"`
	IdSkills3   int    `elements:"name:id_skills_3;type:int;size:4;text:未知;" gorm:"column:id_skills_3;type:integer;default:0;not null;" json:"id_skills_3"`
	IdSkills4   int    `elements:"name:id_skills_4;type:int;size:4;text:未知;" gorm:"column:id_skills_4;type:integer;default:0;not null;" json:"id_skills_4"`
	IdSkills5   int    `elements:"name:id_skills_5;type:int;size:4;text:未知;" gorm:"column:id_skills_5;type:integer;default:0;not null;" json:"id_skills_5"`
	IdSkills6   int    `elements:"name:id_skills_6;type:int;size:4;text:未知;" gorm:"column:id_skills_6;type:integer;default:0;not null;" json:"id_skills_6"`
	IdSkills7   int    `elements:"name:id_skills_7;type:int;size:4;text:未知;" gorm:"column:id_skills_7;type:integer;default:0;not null;" json:"id_skills_7"`
	IdSkills8   int    `elements:"name:id_skills_8;type:int;size:4;text:未知;" gorm:"column:id_skills_8;type:integer;default:0;not null;" json:"id_skills_8"`
	IdSkills9   int    `elements:"name:id_skills_9;type:int;size:4;text:未知;" gorm:"column:id_skills_9;type:integer;default:0;not null;" json:"id_skills_9"`
	IdSkills10  int    `elements:"name:id_skills_10;type:int;size:4;text:未知;" gorm:"column:id_skills_10;type:integer;default:0;not null;" json:"id_skills_10"`
	IdSkills11  int    `elements:"name:id_skills_11;type:int;size:4;text:未知;" gorm:"column:id_skills_11;type:integer;default:0;not null;" json:"id_skills_11"`
	IdSkills12  int    `elements:"name:id_skills_12;type:int;size:4;text:未知;" gorm:"column:id_skills_12;type:integer;default:0;not null;" json:"id_skills_12"`
	IdSkills13  int    `elements:"name:id_skills_13;type:int;size:4;text:未知;" gorm:"column:id_skills_13;type:integer;default:0;not null;" json:"id_skills_13"`
	IdSkills14  int    `elements:"name:id_skills_14;type:int;size:4;text:未知;" gorm:"column:id_skills_14;type:integer;default:0;not null;" json:"id_skills_14"`
	IdSkills15  int    `elements:"name:id_skills_15;type:int;size:4;text:未知;" gorm:"column:id_skills_15;type:integer;default:0;not null;" json:"id_skills_15"`
	IdSkills16  int    `elements:"name:id_skills_16;type:int;size:4;text:未知;" gorm:"column:id_skills_16;type:integer;default:0;not null;" json:"id_skills_16"`
	IdSkills17  int    `elements:"name:id_skills_17;type:int;size:4;text:未知;" gorm:"column:id_skills_17;type:integer;default:0;not null;" json:"id_skills_17"`
	IdSkills18  int    `elements:"name:id_skills_18;type:int;size:4;text:未知;" gorm:"column:id_skills_18;type:integer;default:0;not null;" json:"id_skills_18"`
	IdSkills19  int    `elements:"name:id_skills_19;type:int;size:4;text:未知;" gorm:"column:id_skills_19;type:integer;default:0;not null;" json:"id_skills_19"`
	IdSkills20  int    `elements:"name:id_skills_20;type:int;size:4;text:未知;" gorm:"column:id_skills_20;type:integer;default:0;not null;" json:"id_skills_20"`
	IdSkills21  int    `elements:"name:id_skills_21;type:int;size:4;text:未知;" gorm:"column:id_skills_21;type:integer;default:0;not null;" json:"id_skills_21"`
	IdSkills22  int    `elements:"name:id_skills_22;type:int;size:4;text:未知;" gorm:"column:id_skills_22;type:integer;default:0;not null;" json:"id_skills_22"`
	IdSkills23  int    `elements:"name:id_skills_23;type:int;size:4;text:未知;" gorm:"column:id_skills_23;type:integer;default:0;not null;" json:"id_skills_23"`
	IdSkills24  int    `elements:"name:id_skills_24;type:int;size:4;text:未知;" gorm:"column:id_skills_24;type:integer;default:0;not null;" json:"id_skills_24"`
	IdSkills25  int    `elements:"name:id_skills_25;type:int;size:4;text:未知;" gorm:"column:id_skills_25;type:integer;default:0;not null;" json:"id_skills_25"`
	IdSkills26  int    `elements:"name:id_skills_26;type:int;size:4;text:未知;" gorm:"column:id_skills_26;type:integer;default:0;not null;" json:"id_skills_26"`
	IdSkills27  int    `elements:"name:id_skills_27;type:int;size:4;text:未知;" gorm:"column:id_skills_27;type:integer;default:0;not null;" json:"id_skills_27"`
	IdSkills28  int    `elements:"name:id_skills_28;type:int;size:4;text:未知;" gorm:"column:id_skills_28;type:integer;default:0;not null;" json:"id_skills_28"`
	IdSkills29  int    `elements:"name:id_skills_29;type:int;size:4;text:未知;" gorm:"column:id_skills_29;type:integer;default:0;not null;" json:"id_skills_29"`
	IdSkills30  int    `elements:"name:id_skills_30;type:int;size:4;text:未知;" gorm:"column:id_skills_30;type:integer;default:0;not null;" json:"id_skills_30"`
	IdSkills31  int    `elements:"name:id_skills_31;type:int;size:4;text:未知;" gorm:"column:id_skills_31;type:integer;default:0;not null;" json:"id_skills_31"`
	IdSkills32  int    `elements:"name:id_skills_32;type:int;size:4;text:未知;" gorm:"column:id_skills_32;type:integer;default:0;not null;" json:"id_skills_32"`
	IdSkills33  int    `elements:"name:id_skills_33;type:int;size:4;text:未知;" gorm:"column:id_skills_33;type:integer;default:0;not null;" json:"id_skills_33"`
	IdSkills34  int    `elements:"name:id_skills_34;type:int;size:4;text:未知;" gorm:"column:id_skills_34;type:integer;default:0;not null;" json:"id_skills_34"`
	IdSkills35  int    `elements:"name:id_skills_35;type:int;size:4;text:未知;" gorm:"column:id_skills_35;type:integer;default:0;not null;" json:"id_skills_35"`
	IdSkills36  int    `elements:"name:id_skills_36;type:int;size:4;text:未知;" gorm:"column:id_skills_36;type:integer;default:0;not null;" json:"id_skills_36"`
	IdSkills37  int    `elements:"name:id_skills_37;type:int;size:4;text:未知;" gorm:"column:id_skills_37;type:integer;default:0;not null;" json:"id_skills_37"`
	IdSkills38  int    `elements:"name:id_skills_38;type:int;size:4;text:未知;" gorm:"column:id_skills_38;type:integer;default:0;not null;" json:"id_skills_38"`
	IdSkills39  int    `elements:"name:id_skills_39;type:int;size:4;text:未知;" gorm:"column:id_skills_39;type:integer;default:0;not null;" json:"id_skills_39"`
	IdSkills40  int    `elements:"name:id_skills_40;type:int;size:4;text:未知;" gorm:"column:id_skills_40;type:integer;default:0;not null;" json:"id_skills_40"`
	IdSkills41  int    `elements:"name:id_skills_41;type:int;size:4;text:未知;" gorm:"column:id_skills_41;type:integer;default:0;not null;" json:"id_skills_41"`
	IdSkills42  int    `elements:"name:id_skills_42;type:int;size:4;text:未知;" gorm:"column:id_skills_42;type:integer;default:0;not null;" json:"id_skills_42"`
	IdSkills43  int    `elements:"name:id_skills_43;type:int;size:4;text:未知;" gorm:"column:id_skills_43;type:integer;default:0;not null;" json:"id_skills_43"`
	IdSkills44  int    `elements:"name:id_skills_44;type:int;size:4;text:未知;" gorm:"column:id_skills_44;type:integer;default:0;not null;" json:"id_skills_44"`
	IdSkills45  int    `elements:"name:id_skills_45;type:int;size:4;text:未知;" gorm:"column:id_skills_45;type:integer;default:0;not null;" json:"id_skills_45"`
	IdSkills46  int    `elements:"name:id_skills_46;type:int;size:4;text:未知;" gorm:"column:id_skills_46;type:integer;default:0;not null;" json:"id_skills_46"`
	IdSkills47  int    `elements:"name:id_skills_47;type:int;size:4;text:未知;" gorm:"column:id_skills_47;type:integer;default:0;not null;" json:"id_skills_47"`
	IdSkills48  int    `elements:"name:id_skills_48;type:int;size:4;text:未知;" gorm:"column:id_skills_48;type:integer;default:0;not null;" json:"id_skills_48"`
	IdSkills49  int    `elements:"name:id_skills_49;type:int;size:4;text:未知;" gorm:"column:id_skills_49;type:integer;default:0;not null;" json:"id_skills_49"`
	IdSkills50  int    `elements:"name:id_skills_50;type:int;size:4;text:未知;" gorm:"column:id_skills_50;type:integer;default:0;not null;" json:"id_skills_50"`
	IdSkills51  int    `elements:"name:id_skills_51;type:int;size:4;text:未知;" gorm:"column:id_skills_51;type:integer;default:0;not null;" json:"id_skills_51"`
	IdSkills52  int    `elements:"name:id_skills_52;type:int;size:4;text:未知;" gorm:"column:id_skills_52;type:integer;default:0;not null;" json:"id_skills_52"`
	IdSkills53  int    `elements:"name:id_skills_53;type:int;size:4;text:未知;" gorm:"column:id_skills_53;type:integer;default:0;not null;" json:"id_skills_53"`
	IdSkills54  int    `elements:"name:id_skills_54;type:int;size:4;text:未知;" gorm:"column:id_skills_54;type:integer;default:0;not null;" json:"id_skills_54"`
	IdSkills55  int    `elements:"name:id_skills_55;type:int;size:4;text:未知;" gorm:"column:id_skills_55;type:integer;default:0;not null;" json:"id_skills_55"`
	IdSkills56  int    `elements:"name:id_skills_56;type:int;size:4;text:未知;" gorm:"column:id_skills_56;type:integer;default:0;not null;" json:"id_skills_56"`
	IdSkills57  int    `elements:"name:id_skills_57;type:int;size:4;text:未知;" gorm:"column:id_skills_57;type:integer;default:0;not null;" json:"id_skills_57"`
	IdSkills58  int    `elements:"name:id_skills_58;type:int;size:4;text:未知;" gorm:"column:id_skills_58;type:integer;default:0;not null;" json:"id_skills_58"`
	IdSkills59  int    `elements:"name:id_skills_59;type:int;size:4;text:未知;" gorm:"column:id_skills_59;type:integer;default:0;not null;" json:"id_skills_59"`
	IdSkills60  int    `elements:"name:id_skills_60;type:int;size:4;text:未知;" gorm:"column:id_skills_60;type:integer;default:0;not null;" json:"id_skills_60"`
	IdSkills61  int    `elements:"name:id_skills_61;type:int;size:4;text:未知;" gorm:"column:id_skills_61;type:integer;default:0;not null;" json:"id_skills_61"`
	IdSkills62  int    `elements:"name:id_skills_62;type:int;size:4;text:未知;" gorm:"column:id_skills_62;type:integer;default:0;not null;" json:"id_skills_62"`
	IdSkills63  int    `elements:"name:id_skills_63;type:int;size:4;text:未知;" gorm:"column:id_skills_63;type:integer;default:0;not null;" json:"id_skills_63"`
	IdSkills64  int    `elements:"name:id_skills_64;type:int;size:4;text:未知;" gorm:"column:id_skills_64;type:integer;default:0;not null;" json:"id_skills_64"`
	IdSkills65  int    `elements:"name:id_skills_65;type:int;size:4;text:未知;" gorm:"column:id_skills_65;type:integer;default:0;not null;" json:"id_skills_65"`
	IdSkills66  int    `elements:"name:id_skills_66;type:int;size:4;text:未知;" gorm:"column:id_skills_66;type:integer;default:0;not null;" json:"id_skills_66"`
	IdSkills67  int    `elements:"name:id_skills_67;type:int;size:4;text:未知;" gorm:"column:id_skills_67;type:integer;default:0;not null;" json:"id_skills_67"`
	IdSkills68  int    `elements:"name:id_skills_68;type:int;size:4;text:未知;" gorm:"column:id_skills_68;type:integer;default:0;not null;" json:"id_skills_68"`
	IdSkills69  int    `elements:"name:id_skills_69;type:int;size:4;text:未知;" gorm:"column:id_skills_69;type:integer;default:0;not null;" json:"id_skills_69"`
	IdSkills70  int    `elements:"name:id_skills_70;type:int;size:4;text:未知;" gorm:"column:id_skills_70;type:integer;default:0;not null;" json:"id_skills_70"`
	IdSkills71  int    `elements:"name:id_skills_71;type:int;size:4;text:未知;" gorm:"column:id_skills_71;type:integer;default:0;not null;" json:"id_skills_71"`
	IdSkills72  int    `elements:"name:id_skills_72;type:int;size:4;text:未知;" gorm:"column:id_skills_72;type:integer;default:0;not null;" json:"id_skills_72"`
	IdSkills73  int    `elements:"name:id_skills_73;type:int;size:4;text:未知;" gorm:"column:id_skills_73;type:integer;default:0;not null;" json:"id_skills_73"`
	IdSkills74  int    `elements:"name:id_skills_74;type:int;size:4;text:未知;" gorm:"column:id_skills_74;type:integer;default:0;not null;" json:"id_skills_74"`
	IdSkills75  int    `elements:"name:id_skills_75;type:int;size:4;text:未知;" gorm:"column:id_skills_75;type:integer;default:0;not null;" json:"id_skills_75"`
	IdSkills76  int    `elements:"name:id_skills_76;type:int;size:4;text:未知;" gorm:"column:id_skills_76;type:integer;default:0;not null;" json:"id_skills_76"`
	IdSkills77  int    `elements:"name:id_skills_77;type:int;size:4;text:未知;" gorm:"column:id_skills_77;type:integer;default:0;not null;" json:"id_skills_77"`
	IdSkills78  int    `elements:"name:id_skills_78;type:int;size:4;text:未知;" gorm:"column:id_skills_78;type:integer;default:0;not null;" json:"id_skills_78"`
	IdSkills79  int    `elements:"name:id_skills_79;type:int;size:4;text:未知;" gorm:"column:id_skills_79;type:integer;default:0;not null;" json:"id_skills_79"`
	IdSkills80  int    `elements:"name:id_skills_80;type:int;size:4;text:未知;" gorm:"column:id_skills_80;type:integer;default:0;not null;" json:"id_skills_80"`
	IdSkills81  int    `elements:"name:id_skills_81;type:int;size:4;text:未知;" gorm:"column:id_skills_81;type:integer;default:0;not null;" json:"id_skills_81"`
	IdSkills82  int    `elements:"name:id_skills_82;type:int;size:4;text:未知;" gorm:"column:id_skills_82;type:integer;default:0;not null;" json:"id_skills_82"`
	IdSkills83  int    `elements:"name:id_skills_83;type:int;size:4;text:未知;" gorm:"column:id_skills_83;type:integer;default:0;not null;" json:"id_skills_83"`
	IdSkills84  int    `elements:"name:id_skills_84;type:int;size:4;text:未知;" gorm:"column:id_skills_84;type:integer;default:0;not null;" json:"id_skills_84"`
	IdSkills85  int    `elements:"name:id_skills_85;type:int;size:4;text:未知;" gorm:"column:id_skills_85;type:integer;default:0;not null;" json:"id_skills_85"`
	IdSkills86  int    `elements:"name:id_skills_86;type:int;size:4;text:未知;" gorm:"column:id_skills_86;type:integer;default:0;not null;" json:"id_skills_86"`
	IdSkills87  int    `elements:"name:id_skills_87;type:int;size:4;text:未知;" gorm:"column:id_skills_87;type:integer;default:0;not null;" json:"id_skills_87"`
	IdSkills88  int    `elements:"name:id_skills_88;type:int;size:4;text:未知;" gorm:"column:id_skills_88;type:integer;default:0;not null;" json:"id_skills_88"`
	IdSkills89  int    `elements:"name:id_skills_89;type:int;size:4;text:未知;" gorm:"column:id_skills_89;type:integer;default:0;not null;" json:"id_skills_89"`
	IdSkills90  int    `elements:"name:id_skills_90;type:int;size:4;text:未知;" gorm:"column:id_skills_90;type:integer;default:0;not null;" json:"id_skills_90"`
	IdSkills91  int    `elements:"name:id_skills_91;type:int;size:4;text:未知;" gorm:"column:id_skills_91;type:integer;default:0;not null;" json:"id_skills_91"`
	IdSkills92  int    `elements:"name:id_skills_92;type:int;size:4;text:未知;" gorm:"column:id_skills_92;type:integer;default:0;not null;" json:"id_skills_92"`
	IdSkills93  int    `elements:"name:id_skills_93;type:int;size:4;text:未知;" gorm:"column:id_skills_93;type:integer;default:0;not null;" json:"id_skills_93"`
	IdSkills94  int    `elements:"name:id_skills_94;type:int;size:4;text:未知;" gorm:"column:id_skills_94;type:integer;default:0;not null;" json:"id_skills_94"`
	IdSkills95  int    `elements:"name:id_skills_95;type:int;size:4;text:未知;" gorm:"column:id_skills_95;type:integer;default:0;not null;" json:"id_skills_95"`
	IdSkills96  int    `elements:"name:id_skills_96;type:int;size:4;text:未知;" gorm:"column:id_skills_96;type:integer;default:0;not null;" json:"id_skills_96"`
	IdSkills97  int    `elements:"name:id_skills_97;type:int;size:4;text:未知;" gorm:"column:id_skills_97;type:integer;default:0;not null;" json:"id_skills_97"`
	IdSkills98  int    `elements:"name:id_skills_98;type:int;size:4;text:未知;" gorm:"column:id_skills_98;type:integer;default:0;not null;" json:"id_skills_98"`
	IdSkills99  int    `elements:"name:id_skills_99;type:int;size:4;text:未知;" gorm:"column:id_skills_99;type:integer;default:0;not null;" json:"id_skills_99"`
	IdSkills100 int    `elements:"name:id_skills_100;type:int;size:4;text:未知;" gorm:"column:id_skills_100;type:integer;default:0;not null;" json:"id_skills_100"`
	IdSkills101 int    `elements:"name:id_skills_101;type:int;size:4;text:未知;" gorm:"column:id_skills_101;type:integer;default:0;not null;" json:"id_skills_101"`
	IdSkills102 int    `elements:"name:id_skills_102;type:int;size:4;text:未知;" gorm:"column:id_skills_102;type:integer;default:0;not null;" json:"id_skills_102"`
	IdSkills103 int    `elements:"name:id_skills_103;type:int;size:4;text:未知;" gorm:"column:id_skills_103;type:integer;default:0;not null;" json:"id_skills_103"`
	IdSkills104 int    `elements:"name:id_skills_104;type:int;size:4;text:未知;" gorm:"column:id_skills_104;type:integer;default:0;not null;" json:"id_skills_104"`
	IdSkills105 int    `elements:"name:id_skills_105;type:int;size:4;text:未知;" gorm:"column:id_skills_105;type:integer;default:0;not null;" json:"id_skills_105"`
	IdSkills106 int    `elements:"name:id_skills_106;type:int;size:4;text:未知;" gorm:"column:id_skills_106;type:integer;default:0;not null;" json:"id_skills_106"`
	IdSkills107 int    `elements:"name:id_skills_107;type:int;size:4;text:未知;" gorm:"column:id_skills_107;type:integer;default:0;not null;" json:"id_skills_107"`
	IdSkills108 int    `elements:"name:id_skills_108;type:int;size:4;text:未知;" gorm:"column:id_skills_108;type:integer;default:0;not null;" json:"id_skills_108"`
	IdSkills109 int    `elements:"name:id_skills_109;type:int;size:4;text:未知;" gorm:"column:id_skills_109;type:integer;default:0;not null;" json:"id_skills_109"`
	IdSkills110 int    `elements:"name:id_skills_110;type:int;size:4;text:未知;" gorm:"column:id_skills_110;type:integer;default:0;not null;" json:"id_skills_110"`
	IdSkills111 int    `elements:"name:id_skills_111;type:int;size:4;text:未知;" gorm:"column:id_skills_111;type:integer;default:0;not null;" json:"id_skills_111"`
	IdSkills112 int    `elements:"name:id_skills_112;type:int;size:4;text:未知;" gorm:"column:id_skills_112;type:integer;default:0;not null;" json:"id_skills_112"`
	IdSkills113 int    `elements:"name:id_skills_113;type:int;size:4;text:未知;" gorm:"column:id_skills_113;type:integer;default:0;not null;" json:"id_skills_113"`
	IdSkills114 int    `elements:"name:id_skills_114;type:int;size:4;text:未知;" gorm:"column:id_skills_114;type:integer;default:0;not null;" json:"id_skills_114"`
	IdSkills115 int    `elements:"name:id_skills_115;type:int;size:4;text:未知;" gorm:"column:id_skills_115;type:integer;default:0;not null;" json:"id_skills_115"`
	IdSkills116 int    `elements:"name:id_skills_116;type:int;size:4;text:未知;" gorm:"column:id_skills_116;type:integer;default:0;not null;" json:"id_skills_116"`
	IdSkills117 int    `elements:"name:id_skills_117;type:int;size:4;text:未知;" gorm:"column:id_skills_117;type:integer;default:0;not null;" json:"id_skills_117"`
	IdSkills118 int    `elements:"name:id_skills_118;type:int;size:4;text:未知;" gorm:"column:id_skills_118;type:integer;default:0;not null;" json:"id_skills_118"`
	IdSkills119 int    `elements:"name:id_skills_119;type:int;size:4;text:未知;" gorm:"column:id_skills_119;type:integer;default:0;not null;" json:"id_skills_119"`
	IdSkills120 int    `elements:"name:id_skills_120;type:int;size:4;text:未知;" gorm:"column:id_skills_120;type:integer;default:0;not null;" json:"id_skills_120"`
	IdSkills121 int    `elements:"name:id_skills_121;type:int;size:4;text:未知;" gorm:"column:id_skills_121;type:integer;default:0;not null;" json:"id_skills_121"`
	IdSkills122 int    `elements:"name:id_skills_122;type:int;size:4;text:未知;" gorm:"column:id_skills_122;type:integer;default:0;not null;" json:"id_skills_122"`
	IdSkills123 int    `elements:"name:id_skills_123;type:int;size:4;text:未知;" gorm:"column:id_skills_123;type:integer;default:0;not null;" json:"id_skills_123"`
	IdSkills124 int    `elements:"name:id_skills_124;type:int;size:4;text:未知;" gorm:"column:id_skills_124;type:integer;default:0;not null;" json:"id_skills_124"`
	IdSkills125 int    `elements:"name:id_skills_125;type:int;size:4;text:未知;" gorm:"column:id_skills_125;type:integer;default:0;not null;" json:"id_skills_125"`
	IdSkills126 int    `elements:"name:id_skills_126;type:int;size:4;text:未知;" gorm:"column:id_skills_126;type:integer;default:0;not null;" json:"id_skills_126"`
	IdSkills127 int    `elements:"name:id_skills_127;type:int;size:4;text:未知;" gorm:"column:id_skills_127;type:integer;default:0;not null;" json:"id_skills_127"`
	IdSkills128 int    `elements:"name:id_skills_128;type:int;size:4;text:未知;" gorm:"column:id_skills_128;type:integer;default:0;not null;" json:"id_skills_128"`
	IdDialog    int    `elements:"name:id_dialog;type:int;size:4;text:未知;" gorm:"column:id_dialog;type:integer;default:0;not null;" json:"id_dialog"`
}

func (NpcSkillService) TableName() string {
	return "el_npcskillservice"
}

func (e *NpcSkillService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcSkillService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
