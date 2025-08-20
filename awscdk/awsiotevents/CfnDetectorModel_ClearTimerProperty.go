package awsiotevents


// Information needed to clear the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clearTimerProperty := &ClearTimerProperty{
//   	TimerName: jsii.String("timerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-cleartimer.html
//
type CfnDetectorModel_ClearTimerProperty struct {
	// The name of the timer to clear.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-cleartimer.html#cfn-iotevents-detectormodel-cleartimer-timername
	//
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
}

