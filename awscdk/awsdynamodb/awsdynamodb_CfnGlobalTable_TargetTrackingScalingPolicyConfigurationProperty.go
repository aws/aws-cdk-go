package awsdynamodb


// Defines a target tracking scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingScalingPolicyConfigurationProperty := &targetTrackingScalingPolicyConfigurationProperty{
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	disableScaleIn: jsii.Boolean(false),
//   	scaleInCooldown: jsii.Number(123),
//   	scaleOutCooldown: jsii.Number(123),
//   }
//
type CfnGlobalTable_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Defines a target value for the scaling policy.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// The default value is `false` .
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

