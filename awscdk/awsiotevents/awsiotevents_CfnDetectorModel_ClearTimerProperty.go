package awsiotevents


// Information needed to clear the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clearTimerProperty := &clearTimerProperty{
//   	timerName: jsii.String("timerName"),
//   }
//
type CfnDetectorModel_ClearTimerProperty struct {
	// The name of the timer to clear.
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
}

