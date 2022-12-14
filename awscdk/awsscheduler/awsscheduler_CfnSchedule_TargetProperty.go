package awsscheduler


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
//   		tags: tags,
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
	// `CfnSchedule.TargetProperty.Arn`.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// `CfnSchedule.TargetProperty.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnSchedule.TargetProperty.DeadLetterConfig`.
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// `CfnSchedule.TargetProperty.EcsParameters`.
	EcsParameters interface{} `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// `CfnSchedule.TargetProperty.EventBridgeParameters`.
	EventBridgeParameters interface{} `field:"optional" json:"eventBridgeParameters" yaml:"eventBridgeParameters"`
	// `CfnSchedule.TargetProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnSchedule.TargetProperty.KinesisParameters`.
	KinesisParameters interface{} `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// `CfnSchedule.TargetProperty.RetryPolicy`.
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// `CfnSchedule.TargetProperty.SageMakerPipelineParameters`.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// `CfnSchedule.TargetProperty.SqsParameters`.
	SqsParameters interface{} `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

