package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for an event rule target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//   var role role
//   var ruleTargetInput ruleTargetInput
//
//   ruleTargetConfig := &ruleTargetConfig{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	batchParameters: &batchParametersProperty{
//   		jobDefinition: jsii.String("jobDefinition"),
//   		jobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		arrayProperties: &batchArrayPropertiesProperty{
//   			size: jsii.Number(123),
//   		},
//   		retryStrategy: &batchRetryStrategyProperty{
//   			attempts: jsii.Number(123),
//   		},
//   	},
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
//   			awsVpcConfiguration: &awsVpcConfigurationProperty{
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
//   		placementStrategies: []interface{}{
//   			&placementStrategyProperty{
//   				field: jsii.String("field"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		platformVersion: jsii.String("platformVersion"),
//   		propagateTags: jsii.String("propagateTags"),
//   		referenceId: jsii.String("referenceId"),
//   		tagList: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		taskCount: jsii.Number(123),
//   	},
//   	httpParameters: &httpParametersProperty{
//   		headerParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		pathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		queryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	input: ruleTargetInput,
//   	kinesisParameters: &kinesisParametersProperty{
//   		partitionKeyPath: jsii.String("partitionKeyPath"),
//   	},
//   	retryPolicy: &retryPolicyProperty{
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	role: role,
//   	runCommandParameters: &runCommandParametersProperty{
//   		runCommandTargets: []interface{}{
//   			&runCommandTargetProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	sqsParameters: &sqsParametersProperty{
//   		messageGroupId: jsii.String("messageGroupId"),
//   	},
//   	targetResource: construct,
//   }
//
// Experimental.
type RuleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the target.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Parameters used when the rule invokes Amazon AWS Batch Job/Queue.
	// Experimental.
	BatchParameters *CfnRule_BatchParametersProperty `field:"optional" json:"batchParameters" yaml:"batchParameters"`
	// Contains information about a dead-letter queue configuration.
	// Experimental.
	DeadLetterConfig *CfnRule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The Amazon ECS task definition and task count to use, if the event target is an Amazon ECS task.
	// Experimental.
	EcsParameters *CfnRule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge API destination.
	// Experimental.
	HttpParameters *CfnRule_HttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// A unique, user-defined identifier for the target.
	//
	// Acceptable values
	// include alphanumeric characters, periods (.), hyphens (-), and
	// underscores (_).
	// Deprecated: no replacement. we will always use an autogenerated id.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// What input to send to the event target.
	// Experimental.
	Input RuleTargetInput `field:"optional" json:"input" yaml:"input"`
	// Settings that control shard assignment, when the target is a Kinesis stream.
	//
	// If you don't include this parameter, eventId is used as the
	// partition key.
	// Experimental.
	KinesisParameters *CfnRule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// A RetryPolicy object that includes information about the retry policy settings.
	// Experimental.
	RetryPolicy *CfnRule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Role to use to invoke this event target.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Parameters used when the rule invokes Amazon EC2 Systems Manager Run Command.
	// Experimental.
	RunCommandParameters *CfnRule_RunCommandParametersProperty `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// Parameters used when the FIFO sqs queue is used an event target by the rule.
	// Experimental.
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
	// Experimental.
	TargetResource awscdk.IConstruct `field:"optional" json:"targetResource" yaml:"targetResource"`
}

