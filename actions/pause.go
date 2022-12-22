package actions

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"log"
)

func (ecsClient *EcsAction) PauseECSService(serviceARN string, clusterARN string) {
	log.Printf("Pausing service")
	serviceName := serviceARN
	//update service's desired count to 0
	ecsServiceOutput, _ := ecsClient.UpdateService(context.TODO(),
		&ecs.UpdateServiceInput{Service: aws.String(serviceName),
			Cluster: aws.String(clusterARN), DesiredCount: aws.Int32(0)})
	fmt.Println("ECS service update: %s ", ecsServiceOutput.ResultMetadata)
	// stop the running tasks
	listTasksOutput, _ := ecsClient.ListTasks(context.TODO(),
		&ecs.ListTasksInput{Cluster: aws.String(clusterARN),
			ServiceName: aws.String(serviceName)})
	for _, taskARN := range listTasksOutput.TaskArns {
		log.Printf("Killing tasks: %s", taskARN)
		ecsClient.StopTask(context.TODO(),
			&ecs.StopTaskInput{Cluster: aws.String(clusterARN),
				Task: aws.String(taskARN)})
	}

}
