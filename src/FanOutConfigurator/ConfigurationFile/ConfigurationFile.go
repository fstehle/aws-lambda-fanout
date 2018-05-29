package ConfigurationFile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	FanOutName string               `yaml:"FanOutName"`
	Runtime    string               `yaml:"Runtime"`
	Mappings   []ConfigurationEntry `yaml:"Mappings"`
}

type ConfigurationEntry struct {
	ID                 string `yaml:"ID"`
	SourceType         string `yaml:"SourceType"`
	SourceARN          string `yaml:"SourceARN"`
	DestinationARN     string `yaml:"DestinationARN"`
	DestinationRegion  string `yaml:"DestinationRegion"`
	DestinationRoleARN string `yaml:"DestinationRoleARN"`
	Active             bool   `yaml:"Active"`
}

func ReadConfig(FileName string) Configuration {
	config := Configuration{}
	configFileData, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	error := yaml.Unmarshal([]byte(configFileData), &config)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	if config.Runtime == "" {
	  config.Runtime = "nodejs4.3"
	}

	return config
}
