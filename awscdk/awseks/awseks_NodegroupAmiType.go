package awseks


// The AMI type for your node group.
//
// GPU instance types should use the `AL2_x86_64_GPU` AMI type, which uses the
// Amazon EKS-optimized Linux AMI with GPU support. Non-GPU instances should use the `AL2_x86_64` AMI type, which
// uses the Amazon EKS-optimized Linux AMI.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacity: jsii.Number(0),
//   })
//
//   cluster.addNodegroupCapacity(jsii.String("custom-node-group"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	minSize: jsii.Number(4),
//   	diskSize: jsii.Number(100),
//   	amiType: eks.nodegroupAmiType_AL2_X86_64_GPU,
//   })
//
type NodegroupAmiType string

const (
	// Amazon Linux 2 (x86-64).
	NodegroupAmiType_AL2_X86_64 NodegroupAmiType = "AL2_X86_64"
	// Amazon Linux 2 with GPU support.
	NodegroupAmiType_AL2_X86_64_GPU NodegroupAmiType = "AL2_X86_64_GPU"
	// Amazon Linux 2 (ARM-64).
	NodegroupAmiType_AL2_ARM_64 NodegroupAmiType = "AL2_ARM_64"
	// Bottlerocket Linux(ARM-64).
	NodegroupAmiType_BOTTLEROCKET_ARM_64 NodegroupAmiType = "BOTTLEROCKET_ARM_64"
	// Bottlerocket(x86-64).
	NodegroupAmiType_BOTTLEROCKET_X86_64 NodegroupAmiType = "BOTTLEROCKET_X86_64"
)

