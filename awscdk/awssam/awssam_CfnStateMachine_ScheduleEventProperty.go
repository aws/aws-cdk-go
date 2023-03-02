package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleEventProperty := &scheduleEventProperty{
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   }
//
type CfnStateMachine_ScheduleEventProperty struct {
	// `CfnStateMachine.ScheduleEventProperty.Schedule`.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// `CfnStateMachine.ScheduleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
}

