package mixinsawsiotevents


// Information required to reset the timer.
//
// The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resetTimerProperty := &ResetTimerProperty{
//   	TimerName: jsii.String("timerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-resettimer.html
//
type CfnDetectorModelPropsMixin_ResetTimerProperty struct {
	// The name of the timer to reset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-resettimer.html#cfn-iotevents-detectormodel-resettimer-timername
	//
	TimerName *string `field:"optional" json:"timerName" yaml:"timerName"`
}

