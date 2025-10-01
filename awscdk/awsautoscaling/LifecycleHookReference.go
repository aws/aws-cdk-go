package awsautoscaling


// A reference to a LifecycleHook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleHookReference := &LifecycleHookReference{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	LifecycleHookName: jsii.String("lifecycleHookName"),
//   }
//
type LifecycleHookReference struct {
	// The AutoScalingGroupName of the LifecycleHook resource.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// The LifecycleHookName of the LifecycleHook resource.
	LifecycleHookName *string `field:"required" json:"lifecycleHookName" yaml:"lifecycleHookName"`
}

