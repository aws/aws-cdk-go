package awsautoscaling


// The strategies for when launches fail in an Availability Zone.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	AzCapacityDistributionStrategy: autoscaling.CapacityDistributionStrategy_BALANCED_ONLY,
//   })
//
type CapacityDistributionStrategy string

const (
	// If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.
	CapacityDistributionStrategy_BALANCED_ONLY CapacityDistributionStrategy = "BALANCED_ONLY"
	// If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.
	CapacityDistributionStrategy_BALANCED_BEST_EFFORT CapacityDistributionStrategy = "BALANCED_BEST_EFFORT"
)

