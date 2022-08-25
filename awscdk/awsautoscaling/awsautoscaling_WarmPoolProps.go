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
//   warmPoolProps := &warmPoolProps{
//   	autoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	maxGroupPreparedCapacity: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	poolState: awscdk.Aws_autoscaling.poolState_HIBERNATED,
//   	reuseOnScaleIn: jsii.Boolean(false),
//   }
//
type WarmPoolProps struct {
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
	// The Auto Scaling group to add the warm pool to.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

