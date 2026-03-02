package awsautoscaling


// Deletion protection level for Auto Scaling group.
//
// Example:
//   var vpc Vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	DeletionProtection: autoscaling.DeletionProtection_PREVENT_ALL_DELETION,
//   })
//
type DeletionProtection string

const (
	// No deletion protection.
	DeletionProtection_NONE DeletionProtection = "NONE"
	// Block force delete operations.
	DeletionProtection_PREVENT_FORCE_DELETION DeletionProtection = "PREVENT_FORCE_DELETION"
	// Block all delete operations.
	DeletionProtection_PREVENT_ALL_DELETION DeletionProtection = "PREVENT_ALL_DELETION"
)

