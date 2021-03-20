package initialize

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"goblog/config"
	"goblog/consts"
	"os"
)

func ViperInit(path ...string) *viper.Viper {
	var configFile string
	if len(path) == 0 {
		flag.StringVar(&configFile, "c", "", "choose configFile file.")
		flag.Parse()
		if configFile == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(consts.ConfigEnv); configEnv == "" {
				configFile = consts.ConfigFile
				log.Info().Msgf("您正在使用config的默认值,config的路径为%v", consts.ConfigFile)
			} else {
				configFile = configEnv
				log.Info().Msgf("您正在使用GVA_CONFIG环境变量,config的路径为%v", configFile)
			}
		} else {
			log.Info().Msgf("您正在使用命令行的-c参数传递的值,config的路径为%v", configFile)
		}
	} else {
		configFile = path[0]
		log.Info().Msgf("您正在使用func ViperInit()传递的值,config的路径为%v", configFile)
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		log.Panic().Err(err).Msg("Fatal error configFile file")
	}

	if err := v.Unmarshal(&config.ServerConfig); err != nil {
		log.Panic().Err(err).Msg("Fatal error configFile file")
	}
	return v
}
