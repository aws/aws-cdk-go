package awsekslegacy


// Example:
//   // up to ten spot instances
//   var cluster cluster
//
//   cluster.AddCapacity(jsii.String("spot"), &CapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   	DesiredCapacity: jsii.Number(2),
//   	BootstrapOptions: &BootstrapOptions{
//   		KubeletExtraArgs: jsii.String("--node-labels foo=bar,goo=far"),
//   		AwsApiRetryAttempts: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type BootstrapOptions struct {
	// Additional command line arguments to pass to the `/etc/eks/bootstrap.sh` command.
	// See: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh
	//
	// Experimental.
	AdditionalArgs *string `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// Number of retry attempts for AWS API call (DescribeCluster).
	// Experimental.
	AwsApiRetryAttempts *float64 `field:"optional" json:"awsApiRetryAttempts" yaml:"awsApiRetryAttempts"`
	// The contents of the `/etc/docker/daemon.json` file. Useful if you want a custom config differing from the default one in the EKS AMI.
	// Experimental.
	DockerConfigJson *string `field:"optional" json:"dockerConfigJson" yaml:"dockerConfigJson"`
	// Restores the docker default bridge network.
	// Experimental.
	EnableDockerBridge *bool `field:"optional" json:"enableDockerBridge" yaml:"enableDockerBridge"`
	// Extra arguments to add to the kubelet. Useful for adding labels or taints.
	//
	// For example, `--node-labels foo=bar,goo=far`.
	// Experimental.
	KubeletExtraArgs *string `field:"optional" json:"kubeletExtraArgs" yaml:"kubeletExtraArgs"`
	// Sets `--max-pods` for the kubelet based on the capacity of the EC2 instance.
	// Experimental.
	UseMaxPods *bool `field:"optional" json:"useMaxPods" yaml:"useMaxPods"`
}

