package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	Brokers      string `mapstructure:"brokers"`
	Topic        string `mapstructure:"topic"`
	MailSmtp     string `mapstructure:"gmail_smtp_server"`
	MailUser     string `mapstructure:"gmail_smtp_user"`
	MailPassword string `mapstructure:"gmail_smtp_pass"`
	MailPort     uint   `mapstructure:"gmail_smtp_port"`
	ServerPort   string `mapstructure:"server_port"`
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
