//go:build !dev

package config

import (
	"os"
	"strconv"
)

type App struct {
	Brokers      string
	Topic        string
	MailSmtp     string
	MailUser     string
	MailPassword string
	MailPort     uint
	ServerPort   string
}

var conf *App

func Init() {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	conf = &App{
		Brokers:      os.Getenv("KAFKA_BROKERS"),
		Topic:        os.Getenv("KAFKA_TOPIC_READER"),
		MailSmtp:     os.Getenv("SMTP_SERVER"),
		MailUser:     os.Getenv("SMTP_USER"),
		MailPassword: os.Getenv("STMP_PASS"),
		MailPort:     uint(port),
	}
}

func GetConfig() *App {
	return conf
}
