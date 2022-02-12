package elements

import (
	"pw_server/utils/structEx"
)

type PlayerLevelexpConfig struct {
	ID     int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name   string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Exp1   int    `elements:"name:exp_1;type:int;size:4;text:未知;" gorm:"column:exp_1;type:integer;default:0;not null;" json:"exp_1"`
	Exp2   int    `elements:"name:exp_2;type:int;size:4;text:未知;" gorm:"column:exp_2;type:integer;default:0;not null;" json:"exp_2"`
	Exp3   int    `elements:"name:exp_3;type:int;size:4;text:未知;" gorm:"column:exp_3;type:integer;default:0;not null;" json:"exp_3"`
	Exp4   int    `elements:"name:exp_4;type:int;size:4;text:未知;" gorm:"column:exp_4;type:integer;default:0;not null;" json:"exp_4"`
	Exp5   int    `elements:"name:exp_5;type:int;size:4;text:未知;" gorm:"column:exp_5;type:integer;default:0;not null;" json:"exp_5"`
	Exp6   int    `elements:"name:exp_6;type:int;size:4;text:未知;" gorm:"column:exp_6;type:integer;default:0;not null;" json:"exp_6"`
	Exp7   int    `elements:"name:exp_7;type:int;size:4;text:未知;" gorm:"column:exp_7;type:integer;default:0;not null;" json:"exp_7"`
	Exp8   int    `elements:"name:exp_8;type:int;size:4;text:未知;" gorm:"column:exp_8;type:integer;default:0;not null;" json:"exp_8"`
	Exp9   int    `elements:"name:exp_9;type:int;size:4;text:未知;" gorm:"column:exp_9;type:integer;default:0;not null;" json:"exp_9"`
	Exp10  int    `elements:"name:exp_10;type:int;size:4;text:未知;" gorm:"column:exp_10;type:integer;default:0;not null;" json:"exp_10"`
	Exp11  int    `elements:"name:exp_11;type:int;size:4;text:未知;" gorm:"column:exp_11;type:integer;default:0;not null;" json:"exp_11"`
	Exp12  int    `elements:"name:exp_12;type:int;size:4;text:未知;" gorm:"column:exp_12;type:integer;default:0;not null;" json:"exp_12"`
	Exp13  int    `elements:"name:exp_13;type:int;size:4;text:未知;" gorm:"column:exp_13;type:integer;default:0;not null;" json:"exp_13"`
	Exp14  int    `elements:"name:exp_14;type:int;size:4;text:未知;" gorm:"column:exp_14;type:integer;default:0;not null;" json:"exp_14"`
	Exp15  int    `elements:"name:exp_15;type:int;size:4;text:未知;" gorm:"column:exp_15;type:integer;default:0;not null;" json:"exp_15"`
	Exp16  int    `elements:"name:exp_16;type:int;size:4;text:未知;" gorm:"column:exp_16;type:integer;default:0;not null;" json:"exp_16"`
	Exp17  int    `elements:"name:exp_17;type:int;size:4;text:未知;" gorm:"column:exp_17;type:integer;default:0;not null;" json:"exp_17"`
	Exp18  int    `elements:"name:exp_18;type:int;size:4;text:未知;" gorm:"column:exp_18;type:integer;default:0;not null;" json:"exp_18"`
	Exp19  int    `elements:"name:exp_19;type:int;size:4;text:未知;" gorm:"column:exp_19;type:integer;default:0;not null;" json:"exp_19"`
	Exp20  int    `elements:"name:exp_20;type:int;size:4;text:未知;" gorm:"column:exp_20;type:integer;default:0;not null;" json:"exp_20"`
	Exp21  int    `elements:"name:exp_21;type:int;size:4;text:未知;" gorm:"column:exp_21;type:integer;default:0;not null;" json:"exp_21"`
	Exp22  int    `elements:"name:exp_22;type:int;size:4;text:未知;" gorm:"column:exp_22;type:integer;default:0;not null;" json:"exp_22"`
	Exp23  int    `elements:"name:exp_23;type:int;size:4;text:未知;" gorm:"column:exp_23;type:integer;default:0;not null;" json:"exp_23"`
	Exp24  int    `elements:"name:exp_24;type:int;size:4;text:未知;" gorm:"column:exp_24;type:integer;default:0;not null;" json:"exp_24"`
	Exp25  int    `elements:"name:exp_25;type:int;size:4;text:未知;" gorm:"column:exp_25;type:integer;default:0;not null;" json:"exp_25"`
	Exp26  int    `elements:"name:exp_26;type:int;size:4;text:未知;" gorm:"column:exp_26;type:integer;default:0;not null;" json:"exp_26"`
	Exp27  int    `elements:"name:exp_27;type:int;size:4;text:未知;" gorm:"column:exp_27;type:integer;default:0;not null;" json:"exp_27"`
	Exp28  int    `elements:"name:exp_28;type:int;size:4;text:未知;" gorm:"column:exp_28;type:integer;default:0;not null;" json:"exp_28"`
	Exp29  int    `elements:"name:exp_29;type:int;size:4;text:未知;" gorm:"column:exp_29;type:integer;default:0;not null;" json:"exp_29"`
	Exp30  int    `elements:"name:exp_30;type:int;size:4;text:未知;" gorm:"column:exp_30;type:integer;default:0;not null;" json:"exp_30"`
	Exp31  int    `elements:"name:exp_31;type:int;size:4;text:未知;" gorm:"column:exp_31;type:integer;default:0;not null;" json:"exp_31"`
	Exp32  int    `elements:"name:exp_32;type:int;size:4;text:未知;" gorm:"column:exp_32;type:integer;default:0;not null;" json:"exp_32"`
	Exp33  int    `elements:"name:exp_33;type:int;size:4;text:未知;" gorm:"column:exp_33;type:integer;default:0;not null;" json:"exp_33"`
	Exp34  int    `elements:"name:exp_34;type:int;size:4;text:未知;" gorm:"column:exp_34;type:integer;default:0;not null;" json:"exp_34"`
	Exp35  int    `elements:"name:exp_35;type:int;size:4;text:未知;" gorm:"column:exp_35;type:integer;default:0;not null;" json:"exp_35"`
	Exp36  int    `elements:"name:exp_36;type:int;size:4;text:未知;" gorm:"column:exp_36;type:integer;default:0;not null;" json:"exp_36"`
	Exp37  int    `elements:"name:exp_37;type:int;size:4;text:未知;" gorm:"column:exp_37;type:integer;default:0;not null;" json:"exp_37"`
	Exp38  int    `elements:"name:exp_38;type:int;size:4;text:未知;" gorm:"column:exp_38;type:integer;default:0;not null;" json:"exp_38"`
	Exp39  int    `elements:"name:exp_39;type:int;size:4;text:未知;" gorm:"column:exp_39;type:integer;default:0;not null;" json:"exp_39"`
	Exp40  int    `elements:"name:exp_40;type:int;size:4;text:未知;" gorm:"column:exp_40;type:integer;default:0;not null;" json:"exp_40"`
	Exp41  int    `elements:"name:exp_41;type:int;size:4;text:未知;" gorm:"column:exp_41;type:integer;default:0;not null;" json:"exp_41"`
	Exp42  int    `elements:"name:exp_42;type:int;size:4;text:未知;" gorm:"column:exp_42;type:integer;default:0;not null;" json:"exp_42"`
	Exp43  int    `elements:"name:exp_43;type:int;size:4;text:未知;" gorm:"column:exp_43;type:integer;default:0;not null;" json:"exp_43"`
	Exp44  int    `elements:"name:exp_44;type:int;size:4;text:未知;" gorm:"column:exp_44;type:integer;default:0;not null;" json:"exp_44"`
	Exp45  int    `elements:"name:exp_45;type:int;size:4;text:未知;" gorm:"column:exp_45;type:integer;default:0;not null;" json:"exp_45"`
	Exp46  int    `elements:"name:exp_46;type:int;size:4;text:未知;" gorm:"column:exp_46;type:integer;default:0;not null;" json:"exp_46"`
	Exp47  int    `elements:"name:exp_47;type:int;size:4;text:未知;" gorm:"column:exp_47;type:integer;default:0;not null;" json:"exp_47"`
	Exp48  int    `elements:"name:exp_48;type:int;size:4;text:未知;" gorm:"column:exp_48;type:integer;default:0;not null;" json:"exp_48"`
	Exp49  int    `elements:"name:exp_49;type:int;size:4;text:未知;" gorm:"column:exp_49;type:integer;default:0;not null;" json:"exp_49"`
	Exp50  int    `elements:"name:exp_50;type:int;size:4;text:未知;" gorm:"column:exp_50;type:integer;default:0;not null;" json:"exp_50"`
	Exp51  int    `elements:"name:exp_51;type:int;size:4;text:未知;" gorm:"column:exp_51;type:integer;default:0;not null;" json:"exp_51"`
	Exp52  int    `elements:"name:exp_52;type:int;size:4;text:未知;" gorm:"column:exp_52;type:integer;default:0;not null;" json:"exp_52"`
	Exp53  int    `elements:"name:exp_53;type:int;size:4;text:未知;" gorm:"column:exp_53;type:integer;default:0;not null;" json:"exp_53"`
	Exp54  int    `elements:"name:exp_54;type:int;size:4;text:未知;" gorm:"column:exp_54;type:integer;default:0;not null;" json:"exp_54"`
	Exp55  int    `elements:"name:exp_55;type:int;size:4;text:未知;" gorm:"column:exp_55;type:integer;default:0;not null;" json:"exp_55"`
	Exp56  int    `elements:"name:exp_56;type:int;size:4;text:未知;" gorm:"column:exp_56;type:integer;default:0;not null;" json:"exp_56"`
	Exp57  int    `elements:"name:exp_57;type:int;size:4;text:未知;" gorm:"column:exp_57;type:integer;default:0;not null;" json:"exp_57"`
	Exp58  int    `elements:"name:exp_58;type:int;size:4;text:未知;" gorm:"column:exp_58;type:integer;default:0;not null;" json:"exp_58"`
	Exp59  int    `elements:"name:exp_59;type:int;size:4;text:未知;" gorm:"column:exp_59;type:integer;default:0;not null;" json:"exp_59"`
	Exp60  int    `elements:"name:exp_60;type:int;size:4;text:未知;" gorm:"column:exp_60;type:integer;default:0;not null;" json:"exp_60"`
	Exp61  int    `elements:"name:exp_61;type:int;size:4;text:未知;" gorm:"column:exp_61;type:integer;default:0;not null;" json:"exp_61"`
	Exp62  int    `elements:"name:exp_62;type:int;size:4;text:未知;" gorm:"column:exp_62;type:integer;default:0;not null;" json:"exp_62"`
	Exp63  int    `elements:"name:exp_63;type:int;size:4;text:未知;" gorm:"column:exp_63;type:integer;default:0;not null;" json:"exp_63"`
	Exp64  int    `elements:"name:exp_64;type:int;size:4;text:未知;" gorm:"column:exp_64;type:integer;default:0;not null;" json:"exp_64"`
	Exp65  int    `elements:"name:exp_65;type:int;size:4;text:未知;" gorm:"column:exp_65;type:integer;default:0;not null;" json:"exp_65"`
	Exp66  int    `elements:"name:exp_66;type:int;size:4;text:未知;" gorm:"column:exp_66;type:integer;default:0;not null;" json:"exp_66"`
	Exp67  int    `elements:"name:exp_67;type:int;size:4;text:未知;" gorm:"column:exp_67;type:integer;default:0;not null;" json:"exp_67"`
	Exp68  int    `elements:"name:exp_68;type:int;size:4;text:未知;" gorm:"column:exp_68;type:integer;default:0;not null;" json:"exp_68"`
	Exp69  int    `elements:"name:exp_69;type:int;size:4;text:未知;" gorm:"column:exp_69;type:integer;default:0;not null;" json:"exp_69"`
	Exp70  int    `elements:"name:exp_70;type:int;size:4;text:未知;" gorm:"column:exp_70;type:integer;default:0;not null;" json:"exp_70"`
	Exp71  int    `elements:"name:exp_71;type:int;size:4;text:未知;" gorm:"column:exp_71;type:integer;default:0;not null;" json:"exp_71"`
	Exp72  int    `elements:"name:exp_72;type:int;size:4;text:未知;" gorm:"column:exp_72;type:integer;default:0;not null;" json:"exp_72"`
	Exp73  int    `elements:"name:exp_73;type:int;size:4;text:未知;" gorm:"column:exp_73;type:integer;default:0;not null;" json:"exp_73"`
	Exp74  int    `elements:"name:exp_74;type:int;size:4;text:未知;" gorm:"column:exp_74;type:integer;default:0;not null;" json:"exp_74"`
	Exp75  int    `elements:"name:exp_75;type:int;size:4;text:未知;" gorm:"column:exp_75;type:integer;default:0;not null;" json:"exp_75"`
	Exp76  int    `elements:"name:exp_76;type:int;size:4;text:未知;" gorm:"column:exp_76;type:integer;default:0;not null;" json:"exp_76"`
	Exp77  int    `elements:"name:exp_77;type:int;size:4;text:未知;" gorm:"column:exp_77;type:integer;default:0;not null;" json:"exp_77"`
	Exp78  int    `elements:"name:exp_78;type:int;size:4;text:未知;" gorm:"column:exp_78;type:integer;default:0;not null;" json:"exp_78"`
	Exp79  int    `elements:"name:exp_79;type:int;size:4;text:未知;" gorm:"column:exp_79;type:integer;default:0;not null;" json:"exp_79"`
	Exp80  int    `elements:"name:exp_80;type:int;size:4;text:未知;" gorm:"column:exp_80;type:integer;default:0;not null;" json:"exp_80"`
	Exp81  int    `elements:"name:exp_81;type:int;size:4;text:未知;" gorm:"column:exp_81;type:integer;default:0;not null;" json:"exp_81"`
	Exp82  int    `elements:"name:exp_82;type:int;size:4;text:未知;" gorm:"column:exp_82;type:integer;default:0;not null;" json:"exp_82"`
	Exp83  int    `elements:"name:exp_83;type:int;size:4;text:未知;" gorm:"column:exp_83;type:integer;default:0;not null;" json:"exp_83"`
	Exp84  int    `elements:"name:exp_84;type:int;size:4;text:未知;" gorm:"column:exp_84;type:integer;default:0;not null;" json:"exp_84"`
	Exp85  int    `elements:"name:exp_85;type:int;size:4;text:未知;" gorm:"column:exp_85;type:integer;default:0;not null;" json:"exp_85"`
	Exp86  int    `elements:"name:exp_86;type:int;size:4;text:未知;" gorm:"column:exp_86;type:integer;default:0;not null;" json:"exp_86"`
	Exp87  int    `elements:"name:exp_87;type:int;size:4;text:未知;" gorm:"column:exp_87;type:integer;default:0;not null;" json:"exp_87"`
	Exp88  int    `elements:"name:exp_88;type:int;size:4;text:未知;" gorm:"column:exp_88;type:integer;default:0;not null;" json:"exp_88"`
	Exp89  int    `elements:"name:exp_89;type:int;size:4;text:未知;" gorm:"column:exp_89;type:integer;default:0;not null;" json:"exp_89"`
	Exp90  int    `elements:"name:exp_90;type:int;size:4;text:未知;" gorm:"column:exp_90;type:integer;default:0;not null;" json:"exp_90"`
	Exp91  int    `elements:"name:exp_91;type:int;size:4;text:未知;" gorm:"column:exp_91;type:integer;default:0;not null;" json:"exp_91"`
	Exp92  int    `elements:"name:exp_92;type:int;size:4;text:未知;" gorm:"column:exp_92;type:integer;default:0;not null;" json:"exp_92"`
	Exp93  int    `elements:"name:exp_93;type:int;size:4;text:未知;" gorm:"column:exp_93;type:integer;default:0;not null;" json:"exp_93"`
	Exp94  int    `elements:"name:exp_94;type:int;size:4;text:未知;" gorm:"column:exp_94;type:integer;default:0;not null;" json:"exp_94"`
	Exp95  int    `elements:"name:exp_95;type:int;size:4;text:未知;" gorm:"column:exp_95;type:integer;default:0;not null;" json:"exp_95"`
	Exp96  int    `elements:"name:exp_96;type:int;size:4;text:未知;" gorm:"column:exp_96;type:integer;default:0;not null;" json:"exp_96"`
	Exp97  int    `elements:"name:exp_97;type:int;size:4;text:未知;" gorm:"column:exp_97;type:integer;default:0;not null;" json:"exp_97"`
	Exp98  int    `elements:"name:exp_98;type:int;size:4;text:未知;" gorm:"column:exp_98;type:integer;default:0;not null;" json:"exp_98"`
	Exp99  int    `elements:"name:exp_99;type:int;size:4;text:未知;" gorm:"column:exp_99;type:integer;default:0;not null;" json:"exp_99"`
	Exp100 int    `elements:"name:exp_100;type:int;size:4;text:未知;" gorm:"column:exp_100;type:integer;default:0;not null;" json:"exp_100"`
	Exp101 int    `elements:"name:exp_101;type:int;size:4;text:未知;" gorm:"column:exp_101;type:integer;default:0;not null;" json:"exp_101"`
	Exp102 int    `elements:"name:exp_102;type:int;size:4;text:未知;" gorm:"column:exp_102;type:integer;default:0;not null;" json:"exp_102"`
	Exp103 int    `elements:"name:exp_103;type:int;size:4;text:未知;" gorm:"column:exp_103;type:integer;default:0;not null;" json:"exp_103"`
	Exp104 int    `elements:"name:exp_104;type:int;size:4;text:未知;" gorm:"column:exp_104;type:integer;default:0;not null;" json:"exp_104"`
	Exp105 int    `elements:"name:exp_105;type:int;size:4;text:未知;" gorm:"column:exp_105;type:integer;default:0;not null;" json:"exp_105"`
	Exp106 int    `elements:"name:exp_106;type:int;size:4;text:未知;" gorm:"column:exp_106;type:integer;default:0;not null;" json:"exp_106"`
	Exp107 int    `elements:"name:exp_107;type:int;size:4;text:未知;" gorm:"column:exp_107;type:integer;default:0;not null;" json:"exp_107"`
	Exp108 int    `elements:"name:exp_108;type:int;size:4;text:未知;" gorm:"column:exp_108;type:integer;default:0;not null;" json:"exp_108"`
	Exp109 int    `elements:"name:exp_109;type:int;size:4;text:未知;" gorm:"column:exp_109;type:integer;default:0;not null;" json:"exp_109"`
	Exp110 int    `elements:"name:exp_110;type:int;size:4;text:未知;" gorm:"column:exp_110;type:integer;default:0;not null;" json:"exp_110"`
	Exp111 int    `elements:"name:exp_111;type:int;size:4;text:未知;" gorm:"column:exp_111;type:integer;default:0;not null;" json:"exp_111"`
	Exp112 int    `elements:"name:exp_112;type:int;size:4;text:未知;" gorm:"column:exp_112;type:integer;default:0;not null;" json:"exp_112"`
	Exp113 int    `elements:"name:exp_113;type:int;size:4;text:未知;" gorm:"column:exp_113;type:integer;default:0;not null;" json:"exp_113"`
	Exp114 int    `elements:"name:exp_114;type:int;size:4;text:未知;" gorm:"column:exp_114;type:integer;default:0;not null;" json:"exp_114"`
	Exp115 int    `elements:"name:exp_115;type:int;size:4;text:未知;" gorm:"column:exp_115;type:integer;default:0;not null;" json:"exp_115"`
	Exp116 int    `elements:"name:exp_116;type:int;size:4;text:未知;" gorm:"column:exp_116;type:integer;default:0;not null;" json:"exp_116"`
	Exp117 int    `elements:"name:exp_117;type:int;size:4;text:未知;" gorm:"column:exp_117;type:integer;default:0;not null;" json:"exp_117"`
	Exp118 int    `elements:"name:exp_118;type:int;size:4;text:未知;" gorm:"column:exp_118;type:integer;default:0;not null;" json:"exp_118"`
	Exp119 int    `elements:"name:exp_119;type:int;size:4;text:未知;" gorm:"column:exp_119;type:integer;default:0;not null;" json:"exp_119"`
	Exp120 int    `elements:"name:exp_120;type:int;size:4;text:未知;" gorm:"column:exp_120;type:integer;default:0;not null;" json:"exp_120"`
	Exp121 int    `elements:"name:exp_121;type:int;size:4;text:未知;" gorm:"column:exp_121;type:integer;default:0;not null;" json:"exp_121"`
	Exp122 int    `elements:"name:exp_122;type:int;size:4;text:未知;" gorm:"column:exp_122;type:integer;default:0;not null;" json:"exp_122"`
	Exp123 int    `elements:"name:exp_123;type:int;size:4;text:未知;" gorm:"column:exp_123;type:integer;default:0;not null;" json:"exp_123"`
	Exp124 int    `elements:"name:exp_124;type:int;size:4;text:未知;" gorm:"column:exp_124;type:integer;default:0;not null;" json:"exp_124"`
	Exp125 int    `elements:"name:exp_125;type:int;size:4;text:未知;" gorm:"column:exp_125;type:integer;default:0;not null;" json:"exp_125"`
	Exp126 int    `elements:"name:exp_126;type:int;size:4;text:未知;" gorm:"column:exp_126;type:integer;default:0;not null;" json:"exp_126"`
	Exp127 int    `elements:"name:exp_127;type:int;size:4;text:未知;" gorm:"column:exp_127;type:integer;default:0;not null;" json:"exp_127"`
	Exp128 int    `elements:"name:exp_128;type:int;size:4;text:未知;" gorm:"column:exp_128;type:integer;default:0;not null;" json:"exp_128"`
	Exp129 int    `elements:"name:exp_129;type:int;size:4;text:未知;" gorm:"column:exp_129;type:integer;default:0;not null;" json:"exp_129"`
	Exp130 int    `elements:"name:exp_130;type:int;size:4;text:未知;" gorm:"column:exp_130;type:integer;default:0;not null;" json:"exp_130"`
	Exp131 int    `elements:"name:exp_131;type:int;size:4;text:未知;" gorm:"column:exp_131;type:integer;default:0;not null;" json:"exp_131"`
	Exp132 int    `elements:"name:exp_132;type:int;size:4;text:未知;" gorm:"column:exp_132;type:integer;default:0;not null;" json:"exp_132"`
	Exp133 int    `elements:"name:exp_133;type:int;size:4;text:未知;" gorm:"column:exp_133;type:integer;default:0;not null;" json:"exp_133"`
	Exp134 int    `elements:"name:exp_134;type:int;size:4;text:未知;" gorm:"column:exp_134;type:integer;default:0;not null;" json:"exp_134"`
	Exp135 int    `elements:"name:exp_135;type:int;size:4;text:未知;" gorm:"column:exp_135;type:integer;default:0;not null;" json:"exp_135"`
	Exp136 int    `elements:"name:exp_136;type:int;size:4;text:未知;" gorm:"column:exp_136;type:integer;default:0;not null;" json:"exp_136"`
	Exp137 int    `elements:"name:exp_137;type:int;size:4;text:未知;" gorm:"column:exp_137;type:integer;default:0;not null;" json:"exp_137"`
	Exp138 int    `elements:"name:exp_138;type:int;size:4;text:未知;" gorm:"column:exp_138;type:integer;default:0;not null;" json:"exp_138"`
	Exp139 int    `elements:"name:exp_139;type:int;size:4;text:未知;" gorm:"column:exp_139;type:integer;default:0;not null;" json:"exp_139"`
	Exp140 int    `elements:"name:exp_140;type:int;size:4;text:未知;" gorm:"column:exp_140;type:integer;default:0;not null;" json:"exp_140"`
	Exp141 int    `elements:"name:exp_141;type:int;size:4;text:未知;" gorm:"column:exp_141;type:integer;default:0;not null;" json:"exp_141"`
	Exp142 int    `elements:"name:exp_142;type:int;size:4;text:未知;" gorm:"column:exp_142;type:integer;default:0;not null;" json:"exp_142"`
	Exp143 int    `elements:"name:exp_143;type:int;size:4;text:未知;" gorm:"column:exp_143;type:integer;default:0;not null;" json:"exp_143"`
	Exp144 int    `elements:"name:exp_144;type:int;size:4;text:未知;" gorm:"column:exp_144;type:integer;default:0;not null;" json:"exp_144"`
	Exp145 int    `elements:"name:exp_145;type:int;size:4;text:未知;" gorm:"column:exp_145;type:integer;default:0;not null;" json:"exp_145"`
	Exp146 int    `elements:"name:exp_146;type:int;size:4;text:未知;" gorm:"column:exp_146;type:integer;default:0;not null;" json:"exp_146"`
	Exp147 int    `elements:"name:exp_147;type:int;size:4;text:未知;" gorm:"column:exp_147;type:integer;default:0;not null;" json:"exp_147"`
	Exp148 int    `elements:"name:exp_148;type:int;size:4;text:未知;" gorm:"column:exp_148;type:integer;default:0;not null;" json:"exp_148"`
	Exp149 int    `elements:"name:exp_149;type:int;size:4;text:未知;" gorm:"column:exp_149;type:integer;default:0;not null;" json:"exp_149"`
	Exp150 int    `elements:"name:exp_150;type:int;size:4;text:未知;" gorm:"column:exp_150;type:integer;default:0;not null;" json:"exp_150"`
}

func (PlayerLevelexpConfig) TableName() string {
	return "el_playerlevelexpconfig"
}

func (e *PlayerLevelexpConfig) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *PlayerLevelexpConfig) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
