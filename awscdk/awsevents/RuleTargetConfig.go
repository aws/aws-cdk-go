package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for an event rule target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var role role
//   var ruleTargetInput ruleTargetInput
//
//   ruleTargetConfig := &RuleTargetConfig{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	BatchParameters: &BatchParametersProperty{
//   		JobDefinition: jsii.String("jobDefinition"),
//   		JobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		ArrayProperties: &BatchArrayPropertiesProperty{
//   			Size: jsii.Number(123),
//   		},
//   		RetryStrategy: &BatchRetryStrategyProperty{
//   			Attempts: jsii.Number(123),
//   		},
//   	},
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
//   			AwsVpcConfiguration: &AwsVpcConfigurationProperty{
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
//   		PlacementStrategies: []interface{}{
//   			&PlacementStrategyProperty{
//   				Field: jsii.String("field"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlatformVersion: jsii.String("platformVersion"),
//   		PropagateTags: jsii.String("propagateTags"),
//   		ReferenceId: jsii.String("referenceId"),
//   		TagList: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TaskCount: jsii.Number(123),
//   	},
//   	HttpParameters: &HttpParametersProperty{
//   		HeaderParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		PathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		QueryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	Input: ruleTargetInput,
//   	KinesisParameters: &KinesisParametersProperty{
//   		PartitionKeyPath: jsii.String("partitionKeyPath"),
//   	},
//   	RetryPolicy: &RetryPolicyProperty{
//   		MaximumEventAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   	},
//   	Role: role,
//   	RunCommandParameters: &RunCommandParametersProperty{
//   		RunCommandTargets: []interface{}{
//   			&RunCommandTargetProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	SqsParameters: &SqsParametersProperty{
//   		MessageGroupId: jsii.String("messageGroupId"),
//   	},
//   	TargetResource: construct,
//   }
//
type RuleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the target.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Parameters used when the rule invokes Amazon AWS Batch Job/Queue.
	// Default: no parameters set.
	//
	BatchParameters *CfnRule_BatchParametersProperty `field:"optional" json:"batchParameters" yaml:"batchParameters"`
	// Contains information about a dead-letter queue configuration.
	// Default: no dead-letter queue set.
	//
	DeadLetterConfig *CfnRule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The Amazon ECS task definition and task count to use, if the event target is an Amazon ECS task.
	EcsParameters *CfnRule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge API destination.
	// Default: - None.
	//
	HttpParameters *CfnRule_HttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// What input to send to the event target.
	// Default: the entire event.
	//
	Input RuleTargetInput `field:"optional" json:"input" yaml:"input"`
	// Settings that control shard assignment, when the target is a Kinesis stream.
	//
	// If you don't include this parameter, eventId is used as the
	// partition key.
	KinesisParameters *CfnRule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// A RetryPolicy object that includes information about the retry policy settings.
	// Default: EventBridge default retry policy.
	//
	RetryPolicy *CfnRule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Role to use to invoke this event target.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Parameters used when the rule invokes Amazon EC2 Systems Manager Run Command.
	RunCommandParameters *CfnRule_RunCommandParametersProperty `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// Parameters used when the FIFO sqs queue is used an event target by the rule.
	SqsParameters *CfnRule_SqsParametersProperty `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
	// The resource that is backing this target.
	//
	// This is the resource that will actually have some action performed on it when used as a target
	// (for example, start a build for a CodeBuild project).
	// We need it to determine whether the rule belongs to a different account than the target -
	// if so, we generate a more complex setup,
	// including an additional stack containing the EventBusPolicy.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
	//
	// Default: the target is not backed by any resource.
	//
	TargetResource constructs.IConstruct `field:"optional" json:"targetResource" yaml:"targetResource"`
}

