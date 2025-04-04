package awseks


// Options for adding an AutoScalingGroup as capacity.
//
// Example:
//   var cluster cluster
//   var asg autoScalingGroup
//
//   cluster.connectAutoScalingGroupCapacity(asg, &AutoScalingGroupOptions{
//   })
//
type AutoScalingGroupOptions struct {
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Default: true.
	//
	BootstrapEnabled *bool `field:"optional" json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// Allows options for node bootstrapping through EC2 user data.
	// Default: - default options.
	//
	BootstrapOptions *BootstrapOptions `field:"optional" json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Allow options to specify different machine image type.
	// Default: MachineImageType.AMAZON_LINUX_2
	//
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Default: - true if the cluster has kubectl enabled (which is the default).
	//
	MapRole *bool `field:"optional" json:"mapRole" yaml:"mapRole"`
	// Installs the AWS spot instance interrupt handler on the cluster if it's not already added.
	//
	// Only relevant if `spotPrice` is configured on the auto-scaling group.
	// Default: true.
	//
	SpotInterruptHandler *bool `field:"optional" json:"spotInterruptHandler" yaml:"spotInterruptHandler"`
}

