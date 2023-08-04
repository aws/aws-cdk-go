package awsautoscaling


// Properties for a warm pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   warmPoolProps := &WarmPoolProps{
//   	AutoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	MaxGroupPreparedCapacity: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	PoolState: awscdk.Aws_autoscaling.PoolState_HIBERNATED,
//   	ReuseOnScaleIn: jsii.Boolean(false),
//   }
//
type WarmPoolProps struct {
	// The maximum number of instances that are allowed to be in the warm pool or in any state except Terminated for the Auto Scaling group.
	//
	// If the value is not specified, Amazon EC2 Auto Scaling launches and maintains
	// the difference between the group's maximum capacity and its desired capacity.
	// Default: - max size of the Auto Scaling group.
	//
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// The minimum number of instances to maintain in the warm pool.
	// Default: 0.
	//
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The instance state to transition to after the lifecycle actions are complete.
	// Default: PoolState.STOPPED
	//
	PoolState PoolState `field:"optional" json:"poolState" yaml:"poolState"`
	// Indicates whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	//
	// If the value is not specified, instances in the Auto Scaling group will be terminated
	// when the group scales in.
	// Default: false.
	//
	ReuseOnScaleIn *bool `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
	// The Auto Scaling group to add the warm pool to.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

