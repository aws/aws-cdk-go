package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slidingWindow interface{}
//
//   evaluationWindowProperty := &EvaluationWindowProperty{
//   	SlidingWindow: slidingWindow,
//   	WallClockWindow: &WallClockWindowProperty{
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html
//
type CfnAlarm_EvaluationWindowProperty struct {
	// Configuration for sliding evaluation window (default behavior).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html#cfn-cloudwatch-alarm-evaluationwindow-slidingwindow
	//
	SlidingWindow interface{} `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html#cfn-cloudwatch-alarm-evaluationwindow-wallclockwindow
	//
	WallClockWindow interface{} `field:"optional" json:"wallClockWindow" yaml:"wallClockWindow"`
}

