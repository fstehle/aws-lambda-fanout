package ConfigurationFile

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	FanOutName string `yaml:"FanOutName"`
	Mappings   []ConfigurationEntry `yaml:"Mappings"`
}

type ConfigurationEntry struct {
	ID                 string `yaml:"ID"`
	SourceType         string `yaml:"SourceType"`
	SourceARN          string `yaml:"SourceARN"`
	DestinationARN     string `yaml:"DestinationARN"`
	DestinationRoleARN string `yaml:"DestinationRoleARN"`
	Active             bool `yaml:"Active"`
}

func ReadConfig(FileName string) Configuration {
	config := Configuration{}
	configFileData, err := ioutil.ReadFile(FileName)
	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}

	error := yaml.Unmarshal([]byte(configFileData), &config)

	if (error != nil) {
		fmt.Println(error)
		panic(error)
	}

	return config
}