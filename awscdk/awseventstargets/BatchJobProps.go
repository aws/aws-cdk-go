package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the Batch Job Event Target.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//
//   computeEnvironment := batch.NewFargateComputeEnvironment(this, jsii.String("ComputeEnv"), &FargateComputeEnvironmentProps{
//   	Vpc: Vpc,
//   })
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	Priority: jsii.Number(1),
//   	ComputeEnvironments: []orderedComputeEnvironment{
//   		&orderedComputeEnvironment{
//   			ComputeEnvironment: *ComputeEnvironment,
//   			Order: jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewEcsJobDefinition(this, jsii.String("MyJob"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("Container"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test-repo")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   	}),
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewBatchJob(jobQueue.JobQueueArn, jobQueue, jobDefinition.JobDefinitionArn, jobDefinition, &BatchJobProps{
//   	DeadLetterQueue: queue,
//   	Event: events.RuleTargetInput_FromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	RetryAttempts: jsii.Number(2),
//   	MaxEventAge: awscdk.Duration_*Hours(jsii.Number(2)),
//   }))
//
type BatchJobProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	// Default: no retryStrategy is set.
	//
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	// Default: the entire EventBridge event.
	//
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The name of the submitted job.
	// Default: - Automatically generated.
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The size of the array, if this is an array batch job.
	//
	// Valid values are integers between 2 and 10,000.
	// Default: no arrayProperties are set.
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

