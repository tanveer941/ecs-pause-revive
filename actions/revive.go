package actions

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"log"
)

func ReviveECSService(serviceARN string, clusterARN string) {
	log.Printf("Reviving service....")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	ecsClient := ecs.NewFromConfig(cfg)
	serviceName := serviceARN
	ecsServiceOutput, err := ecsClient.UpdateService(context.TODO(),
		&ecs.UpdateServiceInput{Service: aws.String(serviceName),
			Cluster: aws.String(clusterARN), DesiredCount: aws.Int32(1)})
	log.Printf("ecsServiceOutput: %s", ecsServiceOutput)
}
