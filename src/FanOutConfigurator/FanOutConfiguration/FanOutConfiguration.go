package FanOutConfiguration

import (
	"FanOutConfigurator/ConfigurationFile"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func UpdateFrom(Configuration ConfigurationFile.Configuration) {
	deployFanOut(Configuration.FanOutName)
	updateMappings(Configuration)
}

func deployFanOut(FanOutName string) {
	executeCommand(createDeployCommand(FanOutName))
}

func updateMappings(Configuration ConfigurationFile.Configuration) {
	for _, mapping := range Configuration.Mappings {
		executeCommand(createRegisterCommand(Configuration.FanOutName, mapping))
	}
}

//TODO naming for mapping??
func createRegisterCommand(FanOutName string, Mapping ConfigurationFile.ConfigurationEntry) *exec.Cmd {
	var args []string
	args = append(args, "register")
	args = append(args, "kinesis")
	args = append(args, "--function "+FanOutName)
	args = append(args, "--id "+Mapping.ID)
	args = append(args, "--source-type "+Mapping.SourceType)
	args = append(args, "--source-arn "+Mapping.SourceARN)
	args = append(args, "--destination "+Mapping.DestinationARN)
	if Mapping.DestinationRegion != "" {
		args = append(args, "--destination-region "+Mapping.DestinationRegion)
	}
	args = append(args, "--destination-role-arn "+Mapping.DestinationRoleARN)
	if Mapping.Active {
		args = append(args, "--active true")
	} else {
		args = append(args, "--active false")
	}
	//TODO make configurable:
	cmd := exec.Command("./fanout", args...)
	return cmd
}

func createDeployCommand(FanOutName string) *exec.Cmd {
	var args []string
	args = append(args, "deploy")
	args = append(args, "--function "+FanOutName)
	cmd := exec.Command("./fanout", args...)
	return cmd
}

func executeCommand(cmd *exec.Cmd) {
	fmt.Printf("Executing: %s\n", strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Printf("Output: %s\n", string(output))
	}
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		panic(err)
	}
}
