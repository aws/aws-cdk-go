package awsecs


// The `ManagedScaling` property specifies the settings for the Auto Scaling group capacity provider.
//
// When managed scaling is enabled, Amazon ECS manages the scale-in and scale-out actions of the Auto Scaling group. Amazon ECS manages a target tracking scaling policy using an Amazon ECS-managed CloudWatch metric with the specified `targetCapacity` value as the target value for the metric. For more information, see [Using Managed Scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling) in the *Amazon Elastic Container Service Developer Guide* .
//
// If managed scaling is disabled, the user must manage the scaling of the Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingProperty := &managedScalingProperty{
//   	instanceWarmupPeriod: jsii.Number(123),
//   	maximumScalingStepSize: jsii.Number(123),
//   	minimumScalingStepSize: jsii.Number(123),
//   	status: jsii.String("status"),
//   	targetCapacity: jsii.Number(123),
//   }
//
type CfnCapacityProvider_ManagedScalingProperty struct {
	// The period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group.
	//
	// If this parameter is omitted, the default value of `300` seconds is used.
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// The maximum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `10000` is used.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// The minimum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `1` is used.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Determines whether to use managed scaling for the capacity provider.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target capacity value for the capacity provider.
	//
	// The specified value must be greater than `0` and less than or equal to `100` . A value of `100` results in the Amazon EC2 instances in your Auto Scaling group being completely used.
	TargetCapacity *float64 `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

