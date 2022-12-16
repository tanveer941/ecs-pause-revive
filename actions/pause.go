package actions

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"log"
)

func PauseECSService(serviceARN string, clusterARN string) {
	log.Printf("Pausing service")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	ecsClient := ecs.NewFromConfig(cfg)

	serviceName := serviceARN
	//update service's desired count to 0
	ecsServiceOutput, err := ecsClient.UpdateService(context.TODO(),
		&ecs.UpdateServiceInput{Service: aws.String(serviceName),
			Cluster: aws.String(clusterARN), DesiredCount: aws.Int32(0)})
	fmt.Println("ECS service update: %s ", ecsServiceOutput.ResultMetadata)
	// stop the running tasks
	listTasksOutput, err := ecsClient.ListTasks(context.TODO(),
		&ecs.ListTasksInput{Cluster: aws.String(clusterARN),
			ServiceName: aws.String(serviceName)})
	for _, taskARN := range listTasksOutput.TaskArns {
		log.Printf("Killing tasks: %s", taskARN)
		ecsClient.StopTask(context.TODO(),
			&ecs.StopTaskInput{Cluster: aws.String(clusterARN),
				Task: aws.String(taskARN)})
	}

}
