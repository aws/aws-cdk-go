package interfacesawsautoscalingplans


// A reference to a ScalingPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingPlanReference := &ScalingPlanReference{
//   	ScalingPlanArn: jsii.String("scalingPlanArn"),
//   	ScalingPlanName: jsii.String("scalingPlanName"),
//   	ScalingPlanVersion: jsii.String("scalingPlanVersion"),
//   }
//
type ScalingPlanReference struct {
	// The ARN of the ScalingPlan resource.
	ScalingPlanArn *string `field:"required" json:"scalingPlanArn" yaml:"scalingPlanArn"`
	// The ScalingPlanName of the ScalingPlan resource.
	ScalingPlanName *string `field:"required" json:"scalingPlanName" yaml:"scalingPlanName"`
	// The ScalingPlanVersion of the ScalingPlan resource.
	ScalingPlanVersion *string `field:"required" json:"scalingPlanVersion" yaml:"scalingPlanVersion"`
}

