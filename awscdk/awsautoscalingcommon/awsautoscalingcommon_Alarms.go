package awsautoscalingcommon


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarms := &alarms{
//   	lowerAlarmIntervalIndex: jsii.Number(123),
//   	upperAlarmIntervalIndex: jsii.Number(123),
//   }
//
type Alarms struct {
	LowerAlarmIntervalIndex *float64 `field:"optional" json:"lowerAlarmIntervalIndex" yaml:"lowerAlarmIntervalIndex"`
	UpperAlarmIntervalIndex *float64 `field:"optional" json:"upperAlarmIntervalIndex" yaml:"upperAlarmIntervalIndex"`
}

