package awscdkeks-v2alpha


// The AMI type for your node group.
//
// GPU instance types should use the `AL2_x86_64_GPU` AMI type, which uses the
// Amazon EKS-optimized Linux AMI with GPU support or the `BOTTLEROCKET_ARM_64_NVIDIA` or `BOTTLEROCKET_X86_64_NVIDIA`
// AMI types, which uses the Amazon EKS-optimized Linux AMI with Nvidia-GPU support.
//
// Non-GPU instances should use the `AL2_x86_64` AMI type, which uses the Amazon EKS-optimized Linux AMI.
// Experimental.
type NodegroupAmiType string

const (
	// Amazon Linux 2 (x86-64).
	// Experimental.
	NodegroupAmiType_AL2_X86_64 NodegroupAmiType = "AL2_X86_64"
	// Amazon Linux 2 with GPU support.
	// Experimental.
	NodegroupAmiType_AL2_X86_64_GPU NodegroupAmiType = "AL2_X86_64_GPU"
	// Amazon Linux 2 (ARM-64).
	// Experimental.
	NodegroupAmiType_AL2_ARM_64 NodegroupAmiType = "AL2_ARM_64"
	// Bottlerocket Linux (ARM-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_ARM_64 NodegroupAmiType = "BOTTLEROCKET_ARM_64"
	// Bottlerocket (x86-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_X86_64 NodegroupAmiType = "BOTTLEROCKET_X86_64"
	// Bottlerocket Linux with Nvidia-GPU support (ARM-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_ARM_64_NVIDIA NodegroupAmiType = "BOTTLEROCKET_ARM_64_NVIDIA"
	// Bottlerocket with Nvidia-GPU support (x86-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_X86_64_NVIDIA NodegroupAmiType = "BOTTLEROCKET_X86_64_NVIDIA"
	// Bottlerocket Linux (ARM-64) with FIPS enabled.
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_ARM_64_FIPS NodegroupAmiType = "BOTTLEROCKET_ARM_64_FIPS"
	// Bottlerocket (x86-64) with FIPS enabled.
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_X86_64_FIPS NodegroupAmiType = "BOTTLEROCKET_X86_64_FIPS"
	// Windows Core 2019 (x86-64).
	// Experimental.
	NodegroupAmiType_WINDOWS_CORE_2019_X86_64 NodegroupAmiType = "WINDOWS_CORE_2019_X86_64"
	// Windows Core 2022 (x86-64).
	// Experimental.
	NodegroupAmiType_WINDOWS_CORE_2022_X86_64 NodegroupAmiType = "WINDOWS_CORE_2022_X86_64"
	// Windows Full 2019 (x86-64).
	// Experimental.
	NodegroupAmiType_WINDOWS_FULL_2019_X86_64 NodegroupAmiType = "WINDOWS_FULL_2019_X86_64"
	// Windows Full 2022 (x86-64).
	// Experimental.
	NodegroupAmiType_WINDOWS_FULL_2022_X86_64 NodegroupAmiType = "WINDOWS_FULL_2022_X86_64"
	// Amazon Linux 2023 (x86-64).
	// Experimental.
	NodegroupAmiType_AL2023_X86_64_STANDARD NodegroupAmiType = "AL2023_X86_64_STANDARD"
	// Amazon Linux 2023 with AWS Neuron drivers (x86-64).
	// Experimental.
	NodegroupAmiType_AL2023_X86_64_NEURON NodegroupAmiType = "AL2023_X86_64_NEURON"
	// Amazon Linux 2023 with NVIDIA drivers (x86-64).
	// Experimental.
	NodegroupAmiType_AL2023_X86_64_NVIDIA NodegroupAmiType = "AL2023_X86_64_NVIDIA"
	// Amazon Linux 2023 (ARM-64).
	// Experimental.
	NodegroupAmiType_AL2023_ARM_64_STANDARD NodegroupAmiType = "AL2023_ARM_64_STANDARD"
)

