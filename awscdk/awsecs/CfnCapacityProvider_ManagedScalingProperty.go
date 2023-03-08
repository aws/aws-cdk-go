package awsecs


// The managed scaling settings for the Auto Scaling group capacity provider.
//
// When managed scaling is enabled, Amazon ECS manages the scale-in and scale-out actions of the Auto Scaling group. Amazon ECS manages a target tracking scaling policy using an Amazon ECS managed CloudWatch metric with the specified `targetCapacity` value as the target value for the metric. For more information, see [Using managed scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling) in the *Amazon Elastic Container Service Developer Guide* .
//
// If managed scaling is off, the user must manage the scaling of the Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingProperty := &ManagedScalingProperty{
//   	InstanceWarmupPeriod: jsii.Number(123),
//   	MaximumScalingStepSize: jsii.Number(123),
//   	MinimumScalingStepSize: jsii.Number(123),
//   	Status: jsii.String("status"),
//   	TargetCapacity: jsii.Number(123),
//   }
//
type CfnCapacityProvider_ManagedScalingProperty struct {
	// The period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group.
	//
	// If this parameter is omitted, the default value of `300` seconds is used.
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// The maximum number of Amazon EC2 instances that Amazon ECS will scale out at one time.
	//
	// The scale in process is not affected by this parameter. If this parameter is omitted, the default value of `1` is used.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// The minimum number of Amazon EC2 instances that Amazon ECS will scale out at one time.
	//
	// The scale in process is not affected by this parameter If this parameter is omitted, the default value of `1` is used.
	//
	// When additional capacity is required, Amazon ECS will scale up the minimum scaling step size even if the actual demand is less than the minimum scaling step size.
	//
	// If you use a capacity provider with an Auto Scaling group configured with more than one Amazon EC2 instance type or Availability Zone, Amazon ECS will scale up by the exact minimum scaling step size value and will ignore both the maximum scaling step size as well as the capacity demand.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Determines whether to use managed scaling for the capacity provider.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target capacity value for the capacity provider.
	//
	// The specified value must be greater than `0` and less than or equal to `100` . A value of `100` results in the Amazon EC2 instances in your Auto Scaling group being completely used.
	TargetCapacity *float64 `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

