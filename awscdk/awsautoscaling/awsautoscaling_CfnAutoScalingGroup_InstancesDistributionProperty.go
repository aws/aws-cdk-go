package awsautoscaling


// `InstancesDistribution` is a property of the [AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html) property type that describes an instances distribution for an Auto Scaling group. All properties have a default value, which is the value that is used or assumed when the property is not specified.
//
// The instances distribution specifies the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances, and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities.
//
// For more information and example configurations, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instancesDistributionProperty := &instancesDistributionProperty{
//   	onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   	onDemandBaseCapacity: jsii.Number(123),
//   	onDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   	spotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   	spotInstancePools: jsii.Number(123),
//   	spotMaxPrice: jsii.String("spotMaxPrice"),
//   }
//
type CfnAutoScalingGroup_InstancesDistributionProperty struct {
	// The order of the launch template overrides to use in fulfilling On-Demand capacity.
	//
	// If you specify `lowest-price` , Amazon EC2 Auto Scaling uses price to determine the order, launching the lowest price first.
	//
	// If you specify `prioritized` , Amazon EC2 Auto Scaling uses the priority that you assigned to each launch template override, launching the highest priority first. If all your On-Demand capacity cannot be fulfilled using your highest priority instance, then Amazon EC2 Auto Scaling launches the remaining capacity using the second priority instance type, and so on.
	//
	// Default: `lowest-price` for Auto Scaling groups that specify `InstanceRequirements` in the overrides and `prioritized` for Auto Scaling groups that don't.
	//
	// Valid values: `lowest-price` | `prioritized`.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances.
	//
	// This base portion is launched first as your group scales.
	//
	// If you specify weights for the instance types in the overrides, the base capacity is measured in the same unit of measurement as the instance types. If you specify `InstanceRequirements` in the overrides, the base capacity is measured in the same unit of measurement as your group's desired capacity.
	//
	// Default: `0`
	//
	// > An update to this setting means a gradual replacement of instances to adjust the current On-Demand Instance levels. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
	OnDemandBaseCapacity *float64 `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond `OnDemandBaseCapacity` .
	//
	// Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). If set to 100, only On-Demand Instances are used.
	//
	// Default: `100`
	//
	// > An update to this setting means a gradual replacement of instances to adjust the current On-Demand and Spot Instance levels for your additional capacity higher than the base capacity. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
	OnDemandPercentageAboveBaseCapacity *float64 `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// Indicates how to allocate instances across Spot Instance pools.
	//
	// If the allocation strategy is `lowest-price` , the Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify.
	//
	// If the allocation strategy is `capacity-optimized` (recommended), the Auto Scaling group launches instances using Spot pools that are optimally chosen based on the available Spot capacity. Alternatively, you can use `capacity-optimized-prioritized` and set the order of instance types in the list of launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling honors the instance type priorities on a best-effort basis but optimizes for capacity first.
	//
	// Default: `lowest-price`
	//
	// Valid values: `lowest-price` | `capacity-optimized` | `capacity-optimized-prioritized`.
	SpotAllocationStrategy *string `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// The number of Spot Instance pools across which to allocate your Spot Instances.
	//
	// The Spot pools are determined from the different instance types in the overrides. Valid only when the Spot allocation strategy is `lowest-price` . Value must be in the range of 1–20.
	//
	// Default: `2`.
	SpotInstancePools *float64 `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// If you keep the value at its default (unspecified), Amazon EC2 Auto Scaling uses the On-Demand price as the maximum Spot price. To remove a value that you previously set, include the property but specify an empty string ("") for the value.
	//
	// > If your maximum price is lower than the Spot price for the instance types that you selected, your Spot Instances are not launched.
	//
	// Valid Range: Minimum value of 0.001
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
}

