package mailer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"gopkg.in/yaml.v2"
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

func GetConf() *Config {

	var config *Config

	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/Core/mailer/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}

func (mail *Mailer) Start() {
	mail.Config = GetConf()
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
