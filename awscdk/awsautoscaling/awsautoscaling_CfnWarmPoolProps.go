package awsautoscaling


// Properties for defining a `CfnWarmPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWarmPoolProps := &cfnWarmPoolProps{
//   	autoScalingGroupName: jsii.String("autoScalingGroupName"),
//
//   	// the properties below are optional
//   	instanceReusePolicy: &instanceReusePolicyProperty{
//   		reuseOnScaleIn: jsii.Boolean(false),
//   	},
//   	maxGroupPreparedCapacity: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	poolState: jsii.String("poolState"),
//   }
//
type CfnWarmPoolProps struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Indicates whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	//
	// The default is to terminate instances in the Auto Scaling group when the group scales in.
	InstanceReusePolicy interface{} `field:"optional" json:"instanceReusePolicy" yaml:"instanceReusePolicy"`
	// Specifies the maximum number of instances that are allowed to be in the warm pool or in any state except `Terminated` for the Auto Scaling group.
	//
	// This is an optional property. Specify it only if you do not want the warm pool size to be determined by the difference between the group's maximum capacity and its desired capacity.
	//
	// > If a value for `MaxGroupPreparedCapacity` is not specified, Amazon EC2 Auto Scaling launches and maintains the difference between the group's maximum capacity and its desired capacity. If you specify a value for `MaxGroupPreparedCapacity` , Amazon EC2 Auto Scaling uses the difference between the `MaxGroupPreparedCapacity` and the desired capacity instead.
	// >
	// > The size of the warm pool is dynamic. Only when `MaxGroupPreparedCapacity` and `MinSize` are set to the same value does the warm pool have an absolute size.
	//
	// If the desired capacity of the Auto Scaling group is higher than the `MaxGroupPreparedCapacity` , the capacity of the warm pool is 0, unless you specify a value for `MinSize` . To remove a value that you previously set, include the property but specify -1 for the value.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// Specifies the minimum number of instances to maintain in the warm pool.
	//
	// This helps you to ensure that there is always a certain number of warmed instances available to handle traffic spikes. Defaults to 0 if not specified.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Sets the instance state to transition to after the lifecycle actions are complete.
	//
	// Default is `Stopped` .
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
}

