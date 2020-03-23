package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Config(configPath string) map[string]string {
	config, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var yamlMap map[interface{}]interface{}
	err = yaml.Unmarshal(config, &yamlMap)
	var serverInfoMap map[string]string = map[string]string{
		"IP":   yamlMap["server"].(map[interface{}]interface{})["serverIP"].(string),
		"PORT": yamlMap["server"].(map[interface{}]interface{})["serverPort"].(string),
	}

	return serverInfoMap
}
