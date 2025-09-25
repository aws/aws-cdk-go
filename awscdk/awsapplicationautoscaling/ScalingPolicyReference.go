package awsapplicationautoscaling


// A reference to a ScalingPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingPolicyReference := &ScalingPolicyReference{
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ScalingPolicyArn: jsii.String("scalingPolicyArn"),
//   }
//
type ScalingPolicyReference struct {
	// The ScalableDimension of the ScalingPolicy resource.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The Arn of the ScalingPolicy resource.
	ScalingPolicyArn *string `field:"required" json:"scalingPolicyArn" yaml:"scalingPolicyArn"`
}

