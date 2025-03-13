package awscdkeks-v2alpha


// EKS node bootstrapping options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   bootstrapOptions := &BootstrapOptions{
//   	AdditionalArgs: jsii.String("additionalArgs"),
//   	AwsApiRetryAttempts: jsii.Number(123),
//   	DnsClusterIp: jsii.String("dnsClusterIp"),
//   	DockerConfigJson: jsii.String("dockerConfigJson"),
//   	EnableDockerBridge: jsii.Boolean(false),
//   	KubeletExtraArgs: jsii.String("kubeletExtraArgs"),
//   	UseMaxPods: jsii.Boolean(false),
//   }
//
// Experimental.
type BootstrapOptions struct {
	// Additional command line arguments to pass to the `/etc/eks/bootstrap.sh` command.
	// See: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh
	//
	// Default: - none.
	//
	// Experimental.
	AdditionalArgs *string `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// Number of retry attempts for AWS API call (DescribeCluster).
	// Default: 3.
	//
	// Experimental.
	AwsApiRetryAttempts *float64 `field:"optional" json:"awsApiRetryAttempts" yaml:"awsApiRetryAttempts"`
	// Overrides the IP address to use for DNS queries within the cluster.
	// Default: - 10.100.0.10 or 172.20.0.10 based on the IP
	// address of the primary interface.
	//
	// Experimental.
	DnsClusterIp *string `field:"optional" json:"dnsClusterIp" yaml:"dnsClusterIp"`
	// The contents of the `/etc/docker/daemon.json` file. Useful if you want a custom config differing from the default one in the EKS AMI.
	// Default: - none.
	//
	// Experimental.
	DockerConfigJson *string `field:"optional" json:"dockerConfigJson" yaml:"dockerConfigJson"`
	// Restores the docker default bridge network.
	// Default: false.
	//
	// Experimental.
	EnableDockerBridge *bool `field:"optional" json:"enableDockerBridge" yaml:"enableDockerBridge"`
	// Extra arguments to add to the kubelet. Useful for adding labels or taints.
	//
	// For example, `--node-labels foo=bar,goo=far`.
	// Default: - none.
	//
	// Experimental.
	KubeletExtraArgs *string `field:"optional" json:"kubeletExtraArgs" yaml:"kubeletExtraArgs"`
	// Sets `--max-pods` for the kubelet based on the capacity of the EC2 instance.
	// Default: true.
	//
	// Experimental.
	UseMaxPods *bool `field:"optional" json:"useMaxPods" yaml:"useMaxPods"`
}

