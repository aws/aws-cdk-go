package awsautoscalingplans


// A reference to a ScalingPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingPlanReference := &ScalingPlanReference{
//   	ScalingPlanId: jsii.String("scalingPlanId"),
//   }
//
type ScalingPlanReference struct {
	// The Id of the ScalingPlan resource.
	ScalingPlanId *string `field:"required" json:"scalingPlanId" yaml:"scalingPlanId"`
}

