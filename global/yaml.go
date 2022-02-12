package global

import "github.com/spf13/viper"

type RunConfig struct {
	*ServerCfg `mapstructure:"server"`
	*LogCfg    `mapstructure:"log"`
	*Snowflake `mapstructure:"snowflake"`
	*MySQLCfg  `mapstructure:"mysql"`
}

type ServerCfg struct {
	Port int `mapstructure:"port"`
}

type LogCfg struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxAge     int    `mapstructure:"maxage"`
	MaxBackups int    `mapstructure:"maxbackups"`
}

type Snowflake struct {
	StartTime string `mapstructure:"starttime"`
	MachineID int64  `mapstructure:"machineid"`
}

type MySQLCfg struct {
	Host    string `mapstructure:"host"`
	User    string `mapstructure:"user"`
	Passwd  string `mapstructure:"passwd"`
	DbName  string `mapstructure:"dbname"`
	Port    int    `mapstructure:"port"`
	MaxOpen int    `mapstructure:"maxopen"`
	MaxIdle int    `mapstructure:"maxidle"`
}

func (e *RunConfig) Init() (err error) {
	viper.SetConfigFile("./config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(e); err != nil {
		return
	}
	return
}
