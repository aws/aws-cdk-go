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
//   cfnScheduleProps := &CfnScheduleProps{
//   	FlexibleTimeWindow: &FlexibleTimeWindowProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		MaximumWindowInMinutes: jsii.Number(123),
//   	},
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	Target: &TargetProperty{
//   		Arn: jsii.String("arn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		EcsParameters: &EcsParametersProperty{
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					CapacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					Base: jsii.Number(123),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   				},
//   			},
//   			PlacementConstraints: []interface{}{
//   				&PlacementConstraintProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlacementStrategy: []interface{}{
//   				&PlacementStrategyProperty{
//   					Field: jsii.String("field"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlatformVersion: jsii.String("platformVersion"),
//   			PropagateTags: jsii.String("propagateTags"),
//   			ReferenceId: jsii.String("referenceId"),
//   			Tags: tags,
//   			TaskCount: jsii.Number(123),
//   		},
//   		EventBridgeParameters: &EventBridgeParametersProperty{
//   			DetailType: jsii.String("detailType"),
//   			Source: jsii.String("source"),
//   		},
//   		Input: jsii.String("input"),
//   		KinesisParameters: &KinesisParametersProperty{
//   			PartitionKey: jsii.String("partitionKey"),
//   		},
//   		RetryPolicy: &RetryPolicyProperty{
//   			MaximumEventAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   		},
//   		SageMakerPipelineParameters: &SageMakerPipelineParametersProperty{
//   			PipelineParameterList: []interface{}{
//   				&SageMakerPipelineParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SqsParameters: &SqsParametersProperty{
//   			MessageGroupId: jsii.String("messageGroupId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EndDate: jsii.String("endDate"),
//   	GroupName: jsii.String("groupName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	StartDate: jsii.String("startDate"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html
//
type CfnScheduleProps struct {
	// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-flexibletimewindow
	//
	FlexibleTimeWindow interface{} `field:"required" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// The expression that defines when the schedule runs. The following formats are supported.
	//
	// - `at` expression - `at(yyyy-mm-ddThh:mm:ss)`
	// - `rate` expression - `rate(value unit)`
	// - `cron` expression - `cron(fields)`
	//
	// You can use `at` expressions to create one-time schedules that invoke a target once, at the time and in the time zone, that you specify. You can use `rate` and `cron` expressions to create recurring schedules. Rate-based schedules are useful when you want to invoke a target at regular intervals, such as every 15 minutes or every five days. Cron-based schedules are useful when you want to invoke a target periodically at a specific time, such as at 8:00 am (UTC+0) every 1st day of the month.
	//
	// A `cron` expression consists of six fields separated by white spaces: `(minutes hours day_of_month month day_of_week year)` .
	//
	// A `rate` expression consists of a *value* as a positive integer, and a *unit* with the following options: `minute` | `minutes` | `hour` | `hours` | `day` | `days`
	//
	// For more information and examples, see [Schedule types on EventBridge Scheduler](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html) in the *EventBridge Scheduler User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The schedule's target details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-target
	//
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The description you specify for the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date, in UTC, before which the schedule can invoke its target.
	//
	// Depending on the schedule's recurrence expression, invocations might stop on, or before, the `EndDate` you specify.
	// EventBridge Scheduler ignores `EndDate` for one-time schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-enddate
	//
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The name of the schedule group associated with this schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The Amazon Resource Name (ARN) for the customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The timezone in which the scheduling expression is evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-scheduleexpressiontimezone
	//
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// The date, in UTC, after which the schedule can begin invoking its target.
	//
	// Depending on the schedule's recurrence expression, invocations might occur on, or after, the `StartDate` you specify.
	// EventBridge Scheduler ignores `StartDate` for one-time schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-startdate
	//
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Specifies whether the schedule is enabled or disabled.
	//
	// *Allowed Values* : `ENABLED` | `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

