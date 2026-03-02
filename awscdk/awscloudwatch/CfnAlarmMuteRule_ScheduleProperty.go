package awscloudwatch


// Schedule for the mute to be active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &ScheduleProperty{
//   	Duration: jsii.String("duration"),
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-schedule.html
//
type CfnAlarmMuteRule_ScheduleProperty struct {
	// The duration of the schedule when it triggers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-schedule.html#cfn-cloudwatch-alarmmuterule-schedule-duration
	//
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// The expression of the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-schedule.html#cfn-cloudwatch-alarmmuterule-schedule-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The timezone of the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-schedule.html#cfn-cloudwatch-alarmmuterule-schedule-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

