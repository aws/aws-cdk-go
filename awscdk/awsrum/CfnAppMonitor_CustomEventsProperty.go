package awsrum


// This structure specifies whether this app monitor allows the web client to define and send custom events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customEventsProperty := &CustomEventsProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-customevents.html
//
type CfnAppMonitor_CustomEventsProperty struct {
	// Set this to `ENABLED` to allow the web client to send custom events for this app monitor.
	//
	// Valid values are `ENABLED` and `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-customevents.html#cfn-rum-appmonitor-customevents-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

