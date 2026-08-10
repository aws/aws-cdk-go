package awscloudwatch


// The schedule configuration for the scheduled query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	StartTimeOffset: jsii.Number(123),
//
//   	// the properties below are optional
//   	EndTimeOffset: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html
//
type CfnLogAlarm_ScheduleConfigurationProperty struct {
	// The expression that defines when the scheduled query runs, e.g. rate(1 minute).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The number of seconds into the past to start the query window.
	//
	// Must be a positive value and cannot exceed 2592000 seconds (30 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-starttimeoffset
	//
	StartTimeOffset *float64 `field:"required" json:"startTimeOffset" yaml:"startTimeOffset"`
	// The number of seconds into the past to end the query window.
	//
	// Must be a non-negative value and cannot exceed 2592000 seconds (30 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-endtimeoffset
	//
	EndTimeOffset *float64 `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
}

