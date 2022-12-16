package actions

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/fatih/color"
	"log"
)

const (
	PauseAction  = "Pause"
	ReviveAction = "Revive"
)

var Actions = []string{PauseAction, ReviveAction}

func ChoosePauseOrRevive() (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: "Which action would you like to choose?",
		Options: Actions,
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func getClusterNames() []string {
	// Get the ECS cluster names listed
	log.Printf("Get cluster names....")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	ecsClient := ecs.NewFromConfig(cfg)
	//input := &ecs.ListClustersInput{}
	clusterOutput, err := ecsClient.ListClusters(context.TODO(), &ecs.ListClustersInput{})
	return clusterOutput.ClusterArns
}

func ChooseCluster() (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: "Choose an ECS cluster?",
		Options: getClusterNames(),
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func getServiceNamesFromCluster(clusterARN string) []string {
	// Get the ECS cluster names listed
	log.Printf("Get service names....")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	ecsClient := ecs.NewFromConfig(cfg)
	//input := &ecs.ListClustersInput{}
	serviceOutput, err := ecsClient.ListServices(context.TODO(),
		&ecs.ListServicesInput{Cluster: aws.String(clusterARN)})
	return serviceOutput.ServiceArns
}

func ChooseService(clusterARN string) (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: fmt.Sprintf("Choose an ECS service from cluster %s?", clusterARN),
		Options: getServiceNamesFromCluster(clusterARN),
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func PerformAction(actionChoice string, serviceARN string, clusterARN string) (string, error) {
	if actionChoice == PauseAction {
		PauseECSService(serviceARN, clusterARN)
		color.Red("Service %s is %s", serviceARN, PauseAction)
		return PauseAction, nil
	} else {
		ReviveECSService(serviceARN, clusterARN)
		color.Green("Service %s is %s", serviceARN, ReviveAction)
		return ReviveAction, nil
	}
}
