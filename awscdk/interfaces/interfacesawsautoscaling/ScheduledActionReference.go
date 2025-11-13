package interfacesawsautoscaling


// A reference to a ScheduledAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledActionReference := &ScheduledActionReference{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	ScheduledActionName: jsii.String("scheduledActionName"),
//   }
//
type ScheduledActionReference struct {
	// The AutoScalingGroupName of the ScheduledAction resource.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// The ScheduledActionName of the ScheduledAction resource.
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
}

