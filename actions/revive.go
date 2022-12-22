package actions

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"log"
)

func (ecsClient *EcsAction) ReviveECSService(serviceARN string, clusterARN string) {
	log.Printf("Reviving service....")
	serviceName := serviceARN
	ecsServiceOutput, _ := ecsClient.UpdateService(context.TODO(),
		&ecs.UpdateServiceInput{Service: aws.String(serviceName),
			Cluster: aws.String(clusterARN), DesiredCount: aws.Int32(1)})
	log.Printf("ecsServiceOutput: %s", ecsServiceOutput)
}
