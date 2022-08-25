package awsiotevents


// Specifies whether to get notified for alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acknowledgeFlowProperty := &acknowledgeFlowProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnAlarmModel_AcknowledgeFlowProperty struct {
	// The value must be `TRUE` or `FALSE` .
	//
	// If `TRUE` , you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to `NORMAL` . If `FALSE` , you won't receive notifications. The alarm automatically changes to the `NORMAL` state when the input property value returns to the specified range.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

