package awscloudwatch


// The schedule configuration for the scheduled query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	EndTimeOffset: jsii.Number(123),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	StartTimeOffset: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html
//
type CfnLogAlarmPropsMixin_ScheduleConfigurationProperty struct {
	// The number of seconds into the past to end the query window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-endtimeoffset
	//
	EndTimeOffset *float64 `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The expression that defines when the scheduled query runs, e.g. rate(1 minute).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The number of seconds into the past to start the query window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-scheduleconfiguration.html#cfn-cloudwatch-logalarm-scheduleconfiguration-starttimeoffset
	//
	StartTimeOffset *float64 `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

