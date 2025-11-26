package previewawsioteventsmixins


// Information needed to clear the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clearTimerProperty := &ClearTimerProperty{
//   	TimerName: jsii.String("timerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-cleartimer.html
//
type CfnDetectorModelPropsMixin_ClearTimerProperty struct {
	// The name of the timer to clear.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-cleartimer.html#cfn-iotevents-detectormodel-cleartimer-timername
	//
	TimerName *string `field:"optional" json:"timerName" yaml:"timerName"`
}

