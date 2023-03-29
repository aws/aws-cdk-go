package awsscheduler


// The schedule's target.
//
// EventBridge Scheduler supports templated target that invoke common API operations, as well as universal targets that you can customize to invoke over 6,000 API operations across more than 270 services. You can only specify one templated or universal target for a schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   targetProperty := &TargetProperty{
//   	Arn: jsii.String("arn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	EcsParameters: &EcsParametersProperty{
//   		TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		CapacityProviderStrategy: []interface{}{
//   			&CapacityProviderStrategyItemProperty{
//   				CapacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				Base: jsii.Number(123),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		EnableEcsManagedTags: jsii.Boolean(false),
//   		EnableExecuteCommand: jsii.Boolean(false),
//   		Group: jsii.String("group"),
//   		LaunchType: jsii.String("launchType"),
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   			},
//   		},
//   		PlacementConstraints: []interface{}{
//   			&PlacementConstraintProperty{
//   				Expression: jsii.String("expression"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlacementStrategy: []interface{}{
//   			&PlacementStrategyProperty{
//   				Field: jsii.String("field"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlatformVersion: jsii.String("platformVersion"),
//   		PropagateTags: jsii.String("propagateTags"),
//   		ReferenceId: jsii.String("referenceId"),
//   		Tags: tags,
//   		TaskCount: jsii.Number(123),
//   	},
//   	EventBridgeParameters: &EventBridgeParametersProperty{
//   		DetailType: jsii.String("detailType"),
//   		Source: jsii.String("source"),
//   	},
//   	Input: jsii.String("input"),
//   	KinesisParameters: &KinesisParametersProperty{
//   		PartitionKey: jsii.String("partitionKey"),
//   	},
//   	RetryPolicy: &RetryPolicyProperty{
//   		MaximumEventAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   	},
//   	SageMakerPipelineParameters: &SageMakerPipelineParametersProperty{
//   		PipelineParameterList: []interface{}{
//   			&SageMakerPipelineParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SqsParameters: &SqsParametersProperty{
//   		MessageGroupId: jsii.String("messageGroupId"),
//   	},
//   }
//
type CfnSchedule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the target.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The Amazon Resource Name (ARN) of the IAM role that EventBridge Scheduler will use for this target when the schedule is invoked.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule.
	//
	// If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The templated target type for the Amazon ECS [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) API operation.
	EcsParameters interface{} `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// The templated target type for the EventBridge [`PutEvents`](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) API operation.
	EventBridgeParameters interface{} `field:"optional" json:"eventBridgeParameters" yaml:"eventBridgeParameters"`
	// The text, or well-formed JSON, passed to the target.
	//
	// If you are configuring a templated Lambda , AWS Step Functions , or Amazon EventBridge target, the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, Amazon EventBridge Scheduler delivers a default notification to the target.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The templated target type for the Amazon Kinesis [`PutRecord`](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html) API operation.
	KinesisParameters interface{} `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// A `RetryPolicy` object that includes information about the retry policy settings, including the maximum age of an event, and the maximum number of times EventBridge Scheduler will try to deliver the event to a target.
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// The templated target type for the Amazon SageMaker [`StartPipelineExecution`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StartPipelineExecution.html) API operation.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// The templated target type for the Amazon SQS [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) API operation. Contains the message group ID to use when the target is a FIFO queue. If you specify an Amazon SQS FIFO queue as a target, the queue must have content-based deduplication enabled. For more information, see [Using the Amazon SQS message deduplication ID](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html) in the *Amazon SQS Developer Guide* .
	SqsParameters interface{} `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

