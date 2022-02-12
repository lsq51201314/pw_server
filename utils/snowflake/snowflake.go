package snowflake

import (
	"pw_server/global"
	"time"

	"go.uber.org/zap"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init() (err error) {
	var st time.Time
	if st, err = time.Parse("2006-01-02", global.RunCfg.Snowflake.StartTime); err != nil {
		zap.L().Error("无法初始化SnowFlake时间。", zap.Error(err))
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	if node, err = sf.NewNode(global.RunCfg.Snowflake.MachineID); err != nil {
		zap.L().Error("无法初始化SnowFlake。", zap.Error(err))
		return
	}
	return
}

func SFGenID() int64 {
	return node.Generate().Int64()
}
