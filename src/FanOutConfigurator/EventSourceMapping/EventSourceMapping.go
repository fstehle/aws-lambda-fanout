package EventSourceMapping

import (
	"gopkg.in/fatih/set.v0"
	"github.com/aws/aws-sdk-go/service/lambda"
	"FanOutConfigurator/ConfigurationFile"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"os"
)

func UpdateEventSourceMappings(config ConfigurationFile.Configuration) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lambdaClient := lambda.New(sess, &aws.Config{Region: aws.String("eu-central-1")})

	getFunctionArgs := &lambda.ListEventSourceMappingsInput{
		FunctionName: &config.FanOutName,
	}

	functionData, err := lambdaClient.ListEventSourceMappings(getFunctionArgs)

	if err != nil {
		fmt.Println("Cannot configure function for notifications")
		fmt.Println(err)
		os.Exit(0)
	}

	eventSourcesFromConfig := set.New()
	eventSourcesActive := set.New()

	for _, configurationMapping := range config.Mappings {
		eventSourcesFromConfig.Add(configurationMapping.SourceARN)
	}

	for _, existingMapping := range functionData.EventSourceMappings {
		eventSourcesActive.Add(*existingMapping.EventSourceArn)
	}

	for _, requiredEventSource := range set.Difference(eventSourcesFromConfig, eventSourcesActive).List() {
		addEventSourceMappingFor(requiredEventSource.(string), config.FanOutName, lambdaClient)

	}

}

func addEventSourceMappingFor(eventSourceARN string, fanOutName string, lambdaClient *lambda.Lambda) {
	fmt.Printf("add mapping for %v to lambda\n", eventSourceARN)

	createEventSourceMappingInput := &lambda.CreateEventSourceMappingInput{
		FunctionName: &fanOutName,
		EventSourceArn: &eventSourceARN,
		StartingPosition: aws.String("TRIM_HORIZON"),
	}

	functionData, err := lambdaClient.CreateEventSourceMapping(createEventSourceMappingInput)

	if err != nil {
		fmt.Println("failed to add eventsourcemapping to fanout-lambda")
		panic(err)
	}

	fmt.Println(functionData)
}