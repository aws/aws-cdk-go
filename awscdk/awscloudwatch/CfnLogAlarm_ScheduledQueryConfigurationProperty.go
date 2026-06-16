package awscloudwatch


// The scheduled query configuration for the log alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledQueryConfigurationProperty := &ScheduledQueryConfigurationProperty{
//   	AggregationExpression: jsii.String("aggregationExpression"),
//   	LogGroupIdentifiers: []*string{
//   		jsii.String("logGroupIdentifiers"),
//   	},
//   	QueryLanguage: jsii.String("queryLanguage"),
//   	QueryString: jsii.String("queryString"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		EndTimeOffset: jsii.Number(123),
//   		StartTimeOffset: jsii.Number(123),
//   	},
//   	ScheduledQueryRoleArn: jsii.String("scheduledQueryRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html
//
type CfnLogAlarm_ScheduledQueryConfigurationProperty struct {
	// The aggregation expression for the scheduled query, e.g. count(*) or avg(latency) by host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-aggregationexpression
	//
	AggregationExpression *string `field:"required" json:"aggregationExpression" yaml:"aggregationExpression"`
	// The log groups to query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-loggroupidentifiers
	//
	LogGroupIdentifiers *[]*string `field:"required" json:"logGroupIdentifiers" yaml:"logGroupIdentifiers"`
	// The query language to use for the scheduled query (CWLI or SQL).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-querylanguage
	//
	QueryLanguage *string `field:"required" json:"queryLanguage" yaml:"queryLanguage"`
	// The query string to execute against the specified log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The schedule configuration for the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-scheduleconfiguration
	//
	ScheduleConfiguration interface{} `field:"required" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN of the IAM role that grants permissions to execute the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduledqueryconfiguration.html#cfn-cloudwatch-logalarm-scheduledqueryconfiguration-scheduledqueryrolearn
	//
	ScheduledQueryRoleArn *string `field:"required" json:"scheduledQueryRoleArn" yaml:"scheduledQueryRoleArn"`
}

