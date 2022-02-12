package mailer

import (
	"fmt"
	"net/smtp"
	"root/core/yml"

	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Host     string
	Port     string
	Sender   string
	Password string
}

type Mailer struct {
	Auth   smtp.Auth
	Config *Config
}

func (mail *Mailer) Start() {

	config := &Config{}

	data := yml.Read("/core/mailer/config.yml")
	mapstructure.Decode(data, config)

	mail.Config = config
	mail.Auth = smtp.PlainAuth("", mail.Config.Sender, mail.Config.Password, mail.Config.Host)
}

func (mail *Mailer) Send(to []string, message []byte) {
	err := smtp.SendMail(mail.Config.Host+":"+mail.Config.Port, mail.Auth, mail.Config.Sender, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
