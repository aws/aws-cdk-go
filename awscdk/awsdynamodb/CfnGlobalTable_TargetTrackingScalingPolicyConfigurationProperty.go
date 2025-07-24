package awsdynamodb


// Defines a target tracking scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingScalingPolicyConfigurationProperty := &TargetTrackingScalingPolicyConfigurationProperty{
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	DisableScaleIn: jsii.Boolean(false),
//   	ScaleInCooldown: jsii.Number(123),
//   	ScaleOutCooldown: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.html
//
type CfnGlobalTable_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Defines a target value for the scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.html#cfn-dynamodb-globaltable-targettrackingscalingpolicyconfiguration-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.html#cfn-dynamodb-globaltable-targettrackingscalingpolicyconfiguration-disablescalein
	//
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.html#cfn-dynamodb-globaltable-targettrackingscalingpolicyconfiguration-scaleincooldown
	//
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.html#cfn-dynamodb-globaltable-targettrackingscalingpolicyconfiguration-scaleoutcooldown
	//
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

