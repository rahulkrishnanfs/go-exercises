package bootstrapper

import (
	"customermanage/apputils"
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("appconfig")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config")
	fmt.Print(viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {

		apputils.Logger.Fatal("could not load app config")
	}
	apputils.AppConfig.LogLevel = viper.GetInt("appconfig.LogLevel")
	apputils.AppConfig.Server = viper.GetString("appconfig.Server")
}

func StartUp() {
	apputils.SetLogLevel(apputils.Level(apputils.AppConfig.LogLevel))
}
