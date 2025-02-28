package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

// Config of a Schedule Target used during initialization of Schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import scheduler_alpha "github.com/aws/aws-cdk-go/awscdkscheduleralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var scheduleTargetInput scheduleTargetInput
//   var tags interface{}
//
//   scheduleTargetConfig := &ScheduleTargetConfig{
//   	Arn: jsii.String("arn"),
//   	Role: role,
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
//   	Input: scheduleTargetInput,
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
// Experimental.
type ScheduleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the target.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Role to use to invoke this event target.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule.
	//
	// If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.\
	// Experimental.
	DeadLetterConfig *awsscheduler.CfnSchedule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The templated target type for the Amazon ECS RunTask API Operation.
	// Experimental.
	EcsParameters *awsscheduler.CfnSchedule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// The templated target type for the EventBridge PutEvents API operation.
	// Experimental.
	EventBridgeParameters *awsscheduler.CfnSchedule_EventBridgeParametersProperty `field:"optional" json:"eventBridgeParameters" yaml:"eventBridgeParameters"`
	// What input to pass to the target.
	// Experimental.
	Input ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The templated target type for the Amazon Kinesis PutRecord API operation.
	// Experimental.
	KinesisParameters *awsscheduler.CfnSchedule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// A `RetryPolicy` object that includes information about the retry policy settings, including the maximum age of an event, and the maximum number of times EventBridge Scheduler will try to deliver the event to a target.
	// Experimental.
	RetryPolicy *awsscheduler.CfnSchedule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// The templated target type for the Amazon SageMaker StartPipelineExecution API operation.
	// Experimental.
	SageMakerPipelineParameters *awsscheduler.CfnSchedule_SageMakerPipelineParametersProperty `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// The templated target type for the Amazon SQS SendMessage API Operation.
	// Experimental.
	SqsParameters *awsscheduler.CfnSchedule_SqsParametersProperty `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

