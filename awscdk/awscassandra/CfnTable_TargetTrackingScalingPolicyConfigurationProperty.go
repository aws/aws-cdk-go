package awscassandra


// Amazon Keyspaces supports the `target tracking` auto scaling policy for a provisioned table.
//
// This policy scales a table based on the ratio of consumed to provisioned capacity. The auto scaling target is a percentage of the provisioned capacity of the table.
//
// - `targetTrackingScalingPolicyConfiguration` : To define the target tracking policy, you must define the target value.
//
// - `targetValue` : The target utilization rate of the table. Amazon Keyspaces auto scaling ensures that the ratio of consumed capacity to provisioned capacity stays at or near this value. You define `targetValue` as a percentage. A `double` between 20 and 90. (Required)
// - `disableScaleIn` : A `boolean` that specifies if `scale-in` is disabled or enabled for the table. This parameter is disabled by default. To turn on `scale-in` , set the `boolean` value to `FALSE` . This means that capacity for a table can be automatically scaled down on your behalf. (Optional)
// - `scaleInCooldown` : A cooldown period in seconds between scaling activities that lets the table stabilize before another scale in activity starts. If no value is provided, the default is 0. (Optional)
// - `scaleOutCooldown` : A cooldown period in seconds between scaling activities that lets the table stabilize before another scale out activity starts. If no value is provided, the default is 0. (Optional)
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.html
//
type CfnTable_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Specifies the target value for the target tracking auto scaling policy.
	//
	// Amazon Keyspaces auto scaling scales up capacity automatically when traffic exceeds this target utilization rate, and then back down when it falls below the target. This ensures that the ratio of consumed capacity to provisioned capacity stays at or near this value. You define `targetValue` as a percentage. An `integer` between 20 and 90.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.html#cfn-cassandra-table-targettrackingscalingpolicyconfiguration-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Specifies if `scale-in` is enabled.
	//
	// When auto scaling automatically decreases capacity for a table, the table *scales in* . When scaling policies are set, they can't scale in the table lower than its minimum capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.html#cfn-cassandra-table-targettrackingscalingpolicyconfiguration-disablescalein
	//
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Specifies a `scale-in` cool down period.
	//
	// A cooldown period in seconds between scaling activities that lets the table stabilize before another scaling activity starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.html#cfn-cassandra-table-targettrackingscalingpolicyconfiguration-scaleincooldown
	//
	// Default: - 0.
	//
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Specifies a scale out cool down period.
	//
	// A cooldown period in seconds between scaling activities that lets the table stabilize before another scaling activity starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.html#cfn-cassandra-table-targettrackingscalingpolicyconfiguration-scaleoutcooldown
	//
	// Default: - 0.
	//
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

