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
//   targetProperty := &targetProperty{
//   	arn: jsii.String("arn"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	deadLetterConfig: &deadLetterConfigProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	ecsParameters: &ecsParametersProperty{
//   		taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		capacityProviderStrategy: []interface{}{
//   			&capacityProviderStrategyItemProperty{
//   				capacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				base: jsii.Number(123),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   		enableEcsManagedTags: jsii.Boolean(false),
//   		enableExecuteCommand: jsii.Boolean(false),
//   		group: jsii.String("group"),
//   		launchType: jsii.String("launchType"),
//   		networkConfiguration: &networkConfigurationProperty{
//   			awsvpcConfiguration: &awsVpcConfigurationProperty{
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				assignPublicIp: jsii.String("assignPublicIp"),
//   				securityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   			},
//   		},
//   		placementConstraints: []interface{}{
//   			&placementConstraintProperty{
//   				expression: jsii.String("expression"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		placementStrategy: []interface{}{
//   			&placementStrategyProperty{
//   				field: jsii.String("field"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		platformVersion: jsii.String("platformVersion"),
//   		propagateTags: jsii.String("propagateTags"),
//   		referenceId: jsii.String("referenceId"),
//   		tags: []interface{}{
//   			tags,
//   		},
//   		taskCount: jsii.Number(123),
//   	},
//   	eventBridgeParameters: &eventBridgeParametersProperty{
//   		detailType: jsii.String("detailType"),
//   		source: jsii.String("source"),
//   	},
//   	input: jsii.String("input"),
//   	kinesisParameters: &kinesisParametersProperty{
//   		partitionKey: jsii.String("partitionKey"),
//   	},
//   	retryPolicy: &retryPolicyProperty{
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	sageMakerPipelineParameters: &sageMakerPipelineParametersProperty{
//   		pipelineParameterList: []interface{}{
//   			&sageMakerPipelineParameterProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	sqsParameters: &sqsParametersProperty{
//   		messageGroupId: jsii.String("messageGroupId"),
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

