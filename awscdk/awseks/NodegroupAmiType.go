package awseks


// The AMI type for your node group.
//
// GPU instance types should use the `AL2_x86_64_GPU` AMI type, which uses the
// Amazon EKS-optimized Linux AMI with GPU support. Non-GPU instances should use the `AL2_x86_64` AMI type, which
// uses the Amazon EKS-optimized Linux AMI.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_27(),
//   	DefaultCapacity: jsii.Number(0),
//   })
//
//   cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
//   	InstanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	MinSize: jsii.Number(4),
//   	DiskSize: jsii.Number(100),
//   	AmiType: eks.NodegroupAmiType_AL2_X86_64_GPU,
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
	// Windows Core 2019 (x86-64).
	NodegroupAmiType_WINDOWS_CORE_2019_X86_64 NodegroupAmiType = "WINDOWS_CORE_2019_X86_64"
	// Windows Core 2022 (x86-64).
	NodegroupAmiType_WINDOWS_CORE_2022_X86_64 NodegroupAmiType = "WINDOWS_CORE_2022_X86_64"
	// Windows Full 2019 (x86-64).
	NodegroupAmiType_WINDOWS_FULL_2019_X86_64 NodegroupAmiType = "WINDOWS_FULL_2019_X86_64"
	// Windows Full 2022 (x86-64).
	NodegroupAmiType_WINDOWS_FULL_2022_X86_64 NodegroupAmiType = "WINDOWS_FULL_2022_X86_64"
)

