package awsredshift


// A reference to a ScheduledAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledActionReference := &ScheduledActionReference{
//   	ScheduledActionName: jsii.String("scheduledActionName"),
//   }
//
type ScheduledActionReference struct {
	// The ScheduledActionName of the ScheduledAction resource.
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
}

