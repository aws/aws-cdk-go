package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an AWS Batch Job / Queue as an event rule target.
//
// Most likely the code will look something like this:
// `new BatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition)`
//
// In the future this API will be improved to be fully typed.
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
type BatchJob interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger queue this batch job as a result from an EventBridge event.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for BatchJob
type jsiiProxy_BatchJob struct {
	internal.Type__awseventsIRuleTarget
}

func NewBatchJob(jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) BatchJob {
	_init_.Initialize()

	if err := validateNewBatchJobParameters(jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		&j,
	)

	return &j
}

func NewBatchJob_Override(b BatchJob, jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		b,
	)
}

func (b *jsiiProxy_BatchJob) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := b.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

