package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleEventProperty := &ScheduleEventProperty{
//   	Schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	Input: jsii.String("input"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html
//
type CfnStateMachine_ScheduleEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html#cfn-serverless-statemachine-scheduleevent-schedule
	//
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html#cfn-serverless-statemachine-scheduleevent-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
}

