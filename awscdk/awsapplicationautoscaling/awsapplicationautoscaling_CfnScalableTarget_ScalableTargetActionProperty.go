package awsapplicationautoscaling


// `ScalableTargetAction` specifies the minimum and maximum capacity for the `ScalableTargetAction` property of the [AWS::ApplicationAutoScaling::ScalableTarget ScheduledAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalableTargetActionProperty := &scalableTargetActionProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
type CfnScalableTarget_ScalableTargetActionProperty struct {
	// The maximum capacity.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

