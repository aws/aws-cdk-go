package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLogAlarmPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLogAlarmMixinProps := &CfnLogAlarmMixinProps{
//   	ActionLogLineCount: jsii.Number(123),
//   	ActionLogLineRoleArn: jsii.String("actionLogLineRoleArn"),
//   	ActionsEnabled: jsii.Boolean(false),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	QueryResultsToAlarm: jsii.Number(123),
//   	QueryResultsToEvaluate: jsii.Number(123),
//   	ScheduledQueryConfiguration: &ScheduledQueryConfigurationProperty{
//   		AggregationExpression: jsii.String("aggregationExpression"),
//   		LogGroupIdentifiers: []*string{
//   			jsii.String("logGroupIdentifiers"),
//   		},
//   		QueryString: jsii.String("queryString"),
//   		ScheduleConfiguration: &ScheduleConfigurationProperty{
//   			EndTimeOffset: jsii.Number(123),
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			StartTimeOffset: jsii.Number(123),
//   		},
//   		ScheduledQueryRoleArn: jsii.String("scheduledQueryRoleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html
//
type CfnLogAlarmMixinProps struct {
	// The number of log lines to include in alarm notifications.
	//
	// Valid values are 0 to 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-actionloglinecount
	//
	ActionLogLineCount *float64 `field:"optional" json:"actionLogLineCount" yaml:"actionLogLineCount"`
	// The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications.
	//
	// Required when ActionLogLineCount is greater than 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-actionloglinerolearn
	//
	ActionLogLineRoleArn *string `field:"optional" json:"actionLogLineRoleArn" yaml:"actionLogLineRoleArn"`
	// Indicates whether actions should be executed during any changes to the alarm state.
	//
	// The default is TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-actionsenabled
	//
	// Default: - true.
	//
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-alarmactions
	//
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the log alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-alarmdescription
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the log alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-alarmname
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified threshold and the query results.
	//
	// Valid values are GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, and LessThanOrEqualToThreshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-insufficientdataactions
	//
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The actions to execute when this alarm transitions to the OK state from any other state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-okactions
	//
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// The number of query results that must be breaching to trigger the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-queryresultstoalarm
	//
	QueryResultsToAlarm *float64 `field:"optional" json:"queryResultsToAlarm" yaml:"queryResultsToAlarm"`
	// The number of query results over which data is compared to the specified threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-queryresultstoevaluate
	//
	QueryResultsToEvaluate *float64 `field:"optional" json:"queryResultsToEvaluate" yaml:"queryResultsToEvaluate"`
	// The scheduled query configuration for the log alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration
	//
	ScheduledQueryConfiguration interface{} `field:"optional" json:"scheduledQueryConfiguration" yaml:"scheduledQueryConfiguration"`
	// A list of key-value pairs to associate with the log alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The value to compare against the results of the scheduled query evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Sets how this alarm is to handle missing data points.
	//
	// Valid values are breaching, notBreaching, ignore, and missing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html#cfn-cloudwatch-logalarm-treatmissingdata
	//
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

