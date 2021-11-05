//go:build dev

package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	Brokers      string `mapstructure:"kafka_brokers"`
	Topic        string `mapstructure:"kafka_topic_reader"`
	MailSmtp     string `mapstructure:"smtp_server"`
	MailUser     string `mapstructure:"smtp_user"`
	MailPassword string `mapstructure:"smtp_pass"`
	MailPort     uint   `mapstructure:"smtp_port"`
}

var envConfig *viper.Viper
var conf *App

func Init() {

	envConfig = viper.New()
	envConfig.AddConfigPath(".")
	envConfig.AddConfigPath("../")
	envConfig.SetConfigType("env")
	envConfig.SetConfigName(`.env`)

	if err := envConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error on reading the envConfig file: %v", err)
	}
	marshallErr := envConfig.Unmarshal(&conf)
	if marshallErr != nil {
		log.Fatalf("Error on unmarshalling the envConfig file: %v", marshallErr)
	}

}

func GetConfig() *App {
	return conf
}
