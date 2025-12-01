package interfacesawsautoscaling


// A reference to a AutoScalingGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupReference := &AutoScalingGroupReference{
//   	AutoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   }
//
type AutoScalingGroupReference struct {
	// The ARN of the AutoScalingGroup resource.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The AutoScalingGroupName of the AutoScalingGroup resource.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
}

