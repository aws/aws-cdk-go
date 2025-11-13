package interfacesawsautoscaling


// A reference to a WarmPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   warmPoolReference := &WarmPoolReference{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   }
//
type WarmPoolReference struct {
	// The AutoScalingGroupName of the WarmPool resource.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
}

