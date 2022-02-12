package yml

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string
	Port     string
	Sender   string
	Password string
}

func Read(path string) interface{} {

	var data interface{}

	pwd, _ := os.Getwd()

	yamlFile, err := ioutil.ReadFile(pwd + path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return data

}
