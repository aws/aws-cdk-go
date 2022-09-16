package awsautoscaling


// Options for a warm pool.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   autoScalingGroup.addWarmPool(&warmPoolOptions{
//   	minSize: jsii.Number(1),
//   	reuseOnScaleIn: jsii.Boolean(true),
//   })
//
type WarmPoolOptions struct {
	// The maximum number of instances that are allowed to be in the warm pool or in any state except Terminated for the Auto Scaling group.
	//
	// If the value is not specified, Amazon EC2 Auto Scaling launches and maintains
	// the difference between the group's maximum capacity and its desired capacity.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// The minimum number of instances to maintain in the warm pool.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The instance state to transition to after the lifecycle actions are complete.
	PoolState PoolState `field:"optional" json:"poolState" yaml:"poolState"`
	// Indicates whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	//
	// If the value is not specified, instances in the Auto Scaling group will be terminated
	// when the group scales in.
	ReuseOnScaleIn *bool `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}

