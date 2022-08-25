package awsautoscaling


// InstancesDistribution is a subproperty of MixedInstancesPolicy that describes an instances distribution for an Auto Scaling group.
//
// The instances distribution specifies the distribution of On-Demand Instances
// and Spot Instances, the maximum price to pay for Spot Instances, and how the Auto Scaling group allocates
// instance types to fulfill On-Demand and Spot capacities.
//
// For more information and example configurations, see Auto Scaling groups with multiple instance types
// and purchase options in the Amazon EC2 Auto Scaling User Guide:
//
// https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html
//
// Example:
//   var vpc vpc
//   var launchTemplate1 launchTemplate
//   var launchTemplate2 launchTemplate
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	mixedInstancesPolicy: &mixedInstancesPolicy{
//   		instancesDistribution: &instancesDistribution{
//   			onDemandPercentageAboveBaseCapacity: jsii.Number(50),
//   		},
//   		launchTemplate: launchTemplate1,
//   		launchTemplateOverrides: []launchTemplateOverrides{
//   			&launchTemplateOverrides{
//   				instanceType: ec2.NewInstanceType(jsii.String("t3.micro")),
//   			},
//   			&launchTemplateOverrides{
//   				instanceType: ec2.NewInstanceType(jsii.String("t3a.micro")),
//   			},
//   			&launchTemplateOverrides{
//   				instanceType: ec2.NewInstanceType(jsii.String("t4g.micro")),
//   				launchTemplate: launchTemplate2,
//   			},
//   		},
//   	},
//   })
//
type InstancesDistribution struct {
	// Indicates how to allocate instance types to fulfill On-Demand capacity.
	//
	// The only valid value is prioritized,
	// which is also the default value.
	OnDemandAllocationStrategy OnDemandAllocationStrategy `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances.
	//
	// This
	// base portion is provisioned first as your group scales. Defaults to 0 if not specified. If you specify weights
	// for the instance types in the overrides, set the value of OnDemandBaseCapacity in terms of the number of
	// capacity units, and not the number of instances.
	OnDemandBaseCapacity *float64 `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.
	//
	// Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances).
	// Defaults to 100 if not specified. If set to 100, only On-Demand Instances are provisioned.
	OnDemandPercentageAboveBaseCapacity *float64 `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// If the allocation strategy is lowest-price, the Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify.
	//
	// Defaults to
	// lowest-price if not specified.
	//
	// If the allocation strategy is capacity-optimized (recommended), the Auto Scaling group launches instances using Spot
	// pools that are optimally chosen based on the available Spot capacity. Alternatively, you can use capacity-optimized-prioritized
	// and set the order of instance types in the list of launch template overrides from highest to lowest priority
	// (from first to last in the list). Amazon EC2 Auto Scaling honors the instance type priorities on a best-effort basis but
	// optimizes for capacity first.
	SpotAllocationStrategy SpotAllocationStrategy `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// The number of Spot Instance pools to use to allocate your Spot capacity.
	//
	// The Spot pools are determined from the different instance
	// types in the overrides. Valid only when the Spot allocation strategy is lowest-price. Value must be in the range of 1 to 20.
	// Defaults to 2 if not specified.
	SpotInstancePools *float64 `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// If you leave the value at its default (empty),
	// Amazon EC2 Auto Scaling uses the On-Demand price as the maximum Spot price. To remove a value that you previously set, include
	// the property but specify an empty string ("") for the value.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
}

