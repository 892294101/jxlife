package initialize

import (
	"fmt"
	"github.com/892294101/jxlife/server/global"
	"github.com/892294101/jxutils"
	"github.com/spf13/viper"
	"path"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	home, err := jxutils.GetProgramHome()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	viper.AddConfigPath(path.Join(home, "config"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error resources file: %s \n", err.Error())
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct %s \n", err.Error())
	}
}
