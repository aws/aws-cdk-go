package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduleEventProperty := &ScheduleEventProperty{
//   	Input: jsii.String("input"),
//   	Schedule: jsii.String("schedule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html
//
type CfnStateMachinePropsMixin_ScheduleEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html#cfn-serverless-statemachine-scheduleevent-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-scheduleevent.html#cfn-serverless-statemachine-scheduleevent-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}

