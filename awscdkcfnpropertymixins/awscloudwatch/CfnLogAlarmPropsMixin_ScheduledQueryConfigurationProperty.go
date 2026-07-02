package awscloudwatch


// The scheduled query configuration for the log alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduledQueryConfigurationProperty := &ScheduledQueryConfigurationProperty{
//   	AggregationExpression: jsii.String("aggregationExpression"),
//   	LogGroupIdentifiers: []*string{
//   		jsii.String("logGroupIdentifiers"),
//   	},
//   	QueryString: jsii.String("queryString"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		EndTimeOffset: jsii.Number(123),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		StartTimeOffset: jsii.Number(123),
//   	},
//   	ScheduledQueryRoleArn: jsii.String("scheduledQueryRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html
//
type CfnLogAlarmPropsMixin_ScheduledQueryConfigurationProperty struct {
	// The aggregation expression for the scheduled query, e.g. count(*) or avg(latency) by host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-aggregationexpression
	//
	AggregationExpression *string `field:"optional" json:"aggregationExpression" yaml:"aggregationExpression"`
	// The log groups to query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-loggroupidentifiers
	//
	LogGroupIdentifiers *[]*string `field:"optional" json:"logGroupIdentifiers" yaml:"logGroupIdentifiers"`
	// The query string to execute against the specified log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The schedule configuration for the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-scheduleconfiguration
	//
	ScheduleConfiguration interface{} `field:"optional" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN of the IAM role that grants permissions to execute the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-scheduledqueryrolearn
	//
	ScheduledQueryRoleArn *string `field:"optional" json:"scheduledQueryRoleArn" yaml:"scheduledQueryRoleArn"`
}

