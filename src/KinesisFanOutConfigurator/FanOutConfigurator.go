package main



import (
	"KinesisFanOutConfigurator/ConfigurationFile"
	"KinesisFanOutConfigurator/EventSourceMapping"
	"KinesisFanOutConfigurator/FanOutConfiguration"
	"os"
)

func main() {
	config := ConfigurationFile.ReadConfig(getConfigFilePath())
	FanOutConfiguration.UpdateFrom(config)
	EventSourceMapping.UpdateEventSourceMappings(config)
}


func getConfigFilePath() string {
	if (len(os.Args) != 2) {
		panic("usage: updateFanoutConfig pathToConfigFile")
	}
	path := os.Args[1]

	return path
}
