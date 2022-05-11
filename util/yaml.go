package util

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlVal struct {
	DBName string `yaml:"dbname"`
}

var Yamlvalue *YamlVal

func init() {
	yfile, err := ioutil.ReadFile("config/local.yaml")

	if err != nil {

		log.Fatal(err)
	}
	value := &YamlVal{}

	err2 := yaml.Unmarshal(yfile, &value)

	if err2 != nil {

		log.Fatal(err2)
	}
	Yamlvalue = value
}
