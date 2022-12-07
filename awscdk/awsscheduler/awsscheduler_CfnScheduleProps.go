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
//   			tags: []interface{}{
//   				tags,
//   			},
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
	// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
	FlexibleTimeWindow interface{} `field:"required" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// The expression that defines when the schedule runs. The following formats are supported.
	//
	// - `at` expression - `at(yyyy-mm-ddThh:mm:ss)`
	// - `rate` expression - `rate(unit value)`
	// - `cron` expression - `cron(fields)`
	//
	// You can use `at` expressions to create one-time schedules that invoke a target once, at the time and in the time zone, that you specify. You can use `rate` and `cron` expressions to create recurring schedules. Rate-based schedules are useful when you want to invoke a target at regular intervals, such as every 15 minutes or every five days. Cron-based schedules are useful when you want to invoke a target periodically at a specific time, such as at 8:00 am (UTC+0) every 1st day of the month.
	//
	// A `cron` expression consists of six fields separated by white spaces: `(minutes hours day_of_month month day_of_week year)` .
	//
	// A `rate` expression consists of a *value* as a positive integer, and a *unit* with the following options: `minute` | `minutes` | `hour` | `hours` | `day` | `days`
	//
	// For more information and examples, see [Schedule types on EventBridge Scheduler](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html) in the *EventBridge Scheduler User Guide* .
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The schedule's target details.
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The description you specify for the schedule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date, in UTC, before which the schedule can invoke its target.
	//
	// Depending on the schedule's recurrence expression, invocations might stop on, or before, the `EndDate` you specify.
	// EventBridge Scheduler ignores `EndDate` for one-time schedules.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The name of the schedule group associated with this schedule.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The Amazon Resource Name (ARN) for the customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the schedule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The timezone in which the scheduling expression is evaluated.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// The date, in UTC, after which the schedule can begin invoking its target.
	//
	// Depending on the schedule's recurrence expression, invocations might occur on, or after, the `StartDate` you specify.
	// EventBridge Scheduler ignores `StartDate` for one-time schedules.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Specifies whether the schedule is enabled or disabled.
	State *string `field:"optional" json:"state" yaml:"state"`
}

