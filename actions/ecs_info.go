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

type EcsAction struct {
	*ecs.Client
}

func (e *EcsAction) NewECSClient() (*EcsAction, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	ecsClient := ecs.NewFromConfig(cfg)
	return &EcsAction{ecsClient}, err
}

func ChoosePauseOrRevive() (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: "Which action would you like to choose?",
		Options: Actions,
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func getClusterNames(ecsClient *ecs.Client) []string {
	// Get the ECS cluster names listed
	log.Printf("Get cluster names....")
	clusterOutput, _ := ecsClient.ListClusters(context.TODO(), &ecs.ListClustersInput{})
	return clusterOutput.ClusterArns
}

func (ecsClient *EcsAction) ChooseCluster() (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: "Choose an ECS cluster?",
		Options: getClusterNames(ecsClient.Client),
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func getServiceNamesFromCluster(ecsClient *ecs.Client, clusterARN string) []string {
	// Get the ECS cluster names listed
	log.Printf("Get service names....")
	//input := &ecs.ListClustersInput{}
	serviceOutput, _ := ecsClient.ListServices(context.TODO(),
		&ecs.ListServicesInput{Cluster: aws.String(clusterARN)})
	return serviceOutput.ServiceArns
}

func (ecsClient *EcsAction) ChooseService(clusterARN string) (string, error) {
	var choice string
	var desiredAction = &survey.Select{
		Message: fmt.Sprintf("Choose an ECS service from cluster %s?", clusterARN),
		Options: getServiceNamesFromCluster(ecsClient.Client, clusterARN),
	}
	err := survey.AskOne(desiredAction, &choice)
	return choice, err
}

func (e *EcsAction) PerformAction(actionChoice string, serviceARN string, clusterARN string) (string, error) {

	if actionChoice == PauseAction {
		e.PauseECSService(serviceARN, clusterARN)
		color.Red("Service %s is %s", serviceARN, PauseAction)
		return PauseAction, nil
	} else {
		e.ReviveECSService(serviceARN, clusterARN)
		color.Green("Service %s is %s", serviceARN, ReviveAction)
		return ReviveAction, nil
	}
}
