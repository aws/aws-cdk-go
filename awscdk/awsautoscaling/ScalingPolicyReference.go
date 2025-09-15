package awsautoscaling


// A reference to a ScalingPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingPolicyReference := &ScalingPolicyReference{
//   	ScalingPolicyArn: jsii.String("scalingPolicyArn"),
//   }
//
type ScalingPolicyReference struct {
	// The Arn of the ScalingPolicy resource.
	ScalingPolicyArn *string `field:"required" json:"scalingPolicyArn" yaml:"scalingPolicyArn"`
}

