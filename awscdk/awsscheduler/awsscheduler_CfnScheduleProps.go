package awsscheduler


// Properties for defining a `CfnSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnScheduleProps := &cfnScheduleProps{
//   	flexibleTimeWindow: &flexibleTimeWindowProperty{
//   		mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		maximumWindowInMinutes: jsii.Number(123),
//   	},
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   	target: &targetProperty{
//   		arn: jsii.String("arn"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		deadLetterConfig: &deadLetterConfigProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		ecsParameters: &ecsParametersProperty{
//   			taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			capacityProviderStrategy: []interface{}{
//   				&capacityProviderStrategyItemProperty{
//   					capacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					base: jsii.Number(123),
//   					weight: jsii.Number(123),
//   				},
//   			},
//   			enableEcsManagedTags: jsii.Boolean(false),
//   			enableExecuteCommand: jsii.Boolean(false),
//   			group: jsii.String("group"),
//   			launchType: jsii.String("launchType"),
//   			networkConfiguration: &networkConfigurationProperty{
//   				awsvpcConfiguration: &awsVpcConfigurationProperty{
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					assignPublicIp: jsii.String("assignPublicIp"),
//   					securityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   				},
//   			},
//   			placementConstraints: []interface{}{
//   				&placementConstraintProperty{
//   					expression: jsii.String("expression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			placementStrategy: []interface{}{
//   				&placementStrategyProperty{
//   					field: jsii.String("field"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			platformVersion: jsii.String("platformVersion"),
//   			propagateTags: jsii.String("propagateTags"),
//   			referenceId: jsii.String("referenceId"),
//   			tags: tags,
//   			taskCount: jsii.Number(123),
//   		},
//   		eventBridgeParameters: &eventBridgeParametersProperty{
//   			detailType: jsii.String("detailType"),
//   			source: jsii.String("source"),
//   		},
//   		input: jsii.String("input"),
//   		kinesisParameters: &kinesisParametersProperty{
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		retryPolicy: &retryPolicyProperty{
//   			maximumEventAgeInSeconds: jsii.Number(123),
//   			maximumRetryAttempts: jsii.Number(123),
//   		},
//   		sageMakerPipelineParameters: &sageMakerPipelineParametersProperty{
//   			pipelineParameterList: []interface{}{
//   				&sageMakerPipelineParameterProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		sqsParameters: &sqsParametersProperty{
//   			messageGroupId: jsii.String("messageGroupId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	endDate: jsii.String("endDate"),
//   	groupName: jsii.String("groupName"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	name: jsii.String("name"),
//   	scheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	startDate: jsii.String("startDate"),
//   	state: jsii.String("state"),
//   }
//
type CfnScheduleProps struct {
	// `AWS::Scheduler::Schedule.FlexibleTimeWindow`.
	FlexibleTimeWindow interface{} `field:"required" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// `AWS::Scheduler::Schedule.ScheduleExpression`.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// `AWS::Scheduler::Schedule.Target`.
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// `AWS::Scheduler::Schedule.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Scheduler::Schedule.EndDate`.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// `AWS::Scheduler::Schedule.GroupName`.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// `AWS::Scheduler::Schedule.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `AWS::Scheduler::Schedule.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Scheduler::Schedule.ScheduleExpressionTimezone`.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// `AWS::Scheduler::Schedule.StartDate`.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// `AWS::Scheduler::Schedule.State`.
	State *string `field:"optional" json:"state" yaml:"state"`
}

