package awseks


// EKS node bootstrapping options.
//
// Example:
//   var cluster Cluster
//
//   cluster.AddAutoScalingGroupCapacity(jsii.String("spot"), &AutoScalingGroupCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   	MinCapacity: jsii.Number(2),
//   	BootstrapOptions: &BootstrapOptions{
//   		KubeletExtraArgs: jsii.String("--node-labels foo=bar,goo=far"),
//   		AwsApiRetryAttempts: jsii.Number(5),
//   	},
//   })
//
type BootstrapOptions struct {
	// Additional command line arguments to pass to the `/etc/eks/bootstrap.sh` command.
	// See: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh
	//
	// Default: - none.
	//
	AdditionalArgs *string `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// Number of retry attempts for AWS API call (DescribeCluster).
	// Default: 3.
	//
	AwsApiRetryAttempts *float64 `field:"optional" json:"awsApiRetryAttempts" yaml:"awsApiRetryAttempts"`
	// Overrides the IP address to use for DNS queries within the cluster.
	// Default: - 10.100.0.10 or 172.20.0.10 based on the IP
	// address of the primary interface.
	//
	DnsClusterIp *string `field:"optional" json:"dnsClusterIp" yaml:"dnsClusterIp"`
	// The contents of the `/etc/docker/daemon.json` file. Useful if you want a custom config differing from the default one in the EKS AMI.
	// Default: - none.
	//
	DockerConfigJson *string `field:"optional" json:"dockerConfigJson" yaml:"dockerConfigJson"`
	// Restores the docker default bridge network.
	// Default: false.
	//
	EnableDockerBridge *bool `field:"optional" json:"enableDockerBridge" yaml:"enableDockerBridge"`
	// Extra arguments to add to the kubelet. Useful for adding labels or taints.
	//
	// For example, `--node-labels foo=bar,goo=far`.
	// Default: - none.
	//
	KubeletExtraArgs *string `field:"optional" json:"kubeletExtraArgs" yaml:"kubeletExtraArgs"`
	// Sets `--max-pods` for the kubelet based on the capacity of the EC2 instance.
	// Default: true.
	//
	UseMaxPods *bool `field:"optional" json:"useMaxPods" yaml:"useMaxPods"`
}

