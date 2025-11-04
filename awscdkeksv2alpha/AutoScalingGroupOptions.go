package awscdkeksv2alpha


// Options for adding an AutoScalingGroup as capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   autoScalingGroupOptions := &AutoScalingGroupOptions{
//   	BootstrapEnabled: jsii.Boolean(false),
//   	BootstrapOptions: &BootstrapOptions{
//   		AdditionalArgs: jsii.String("additionalArgs"),
//   		AwsApiRetryAttempts: jsii.Number(123),
//   		DnsClusterIp: jsii.String("dnsClusterIp"),
//   		DockerConfigJson: jsii.String("dockerConfigJson"),
//   		EnableDockerBridge: jsii.Boolean(false),
//   		KubeletExtraArgs: jsii.String("kubeletExtraArgs"),
//   		UseMaxPods: jsii.Boolean(false),
//   	},
//   	MachineImageType: eks_v2_alpha.MachineImageType_AMAZON_LINUX_2,
//   }
//
// Experimental.
type AutoScalingGroupOptions struct {
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Default: true.
	//
	// Experimental.
	BootstrapEnabled *bool `field:"optional" json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// Allows options for node bootstrapping through EC2 user data.
	// Default: - default options.
	//
	// Experimental.
	BootstrapOptions *BootstrapOptions `field:"optional" json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Allow options to specify different machine image type.
	// Default: MachineImageType.AMAZON_LINUX_2
	//
	// Experimental.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
}

