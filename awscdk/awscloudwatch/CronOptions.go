package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions.
// Absence of a field implies '*'.
//
// Example:
//   var alarm1 Alarm
//   var alarm2 Alarm
//
//
//   alarmMuteRule := cloudwatch.NewAlarmMuteRule(this, jsii.String("AlarmMuteRule"), &AlarmMuteRuleProps{
//   	Alarms: []IAlarmRef{
//   		alarm1,
//   	},
//   	// Defines the mute period begins at 0:00 everyday in UTC
//   	Schedule: cloudwatch.ScheduleExpression_Cron(&CronOptions{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("0"),
//   	}),
//   	// Specifies the mute rule lasts 1 hour.
//   	Duration: awscdk.Duration_Hours(jsii.Number(1)),
//   })
//
//   // The mute target can be added after construction.
//   alarmMuteRule.AddAlarm(alarm2)
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-mute-rules.html#defining-alarm-mute-rules
//
type CronOptions struct {
	// The minute to run this rule at.
	Minute *string `field:"required" json:"minute" yaml:"minute"`
	// The day of the month to run this rule at.
	// Default: - Every day of the month.
	//
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Default: - Every hour.
	//
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The month to run this rule at.
	// Default: - Every month.
	//
	Month *string `field:"optional" json:"month" yaml:"month"`
	// The timezone to run the schedule in.
	// Default: - TimeZone.ETC_UTC
	//
	TimeZone awscdk.TimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
}

