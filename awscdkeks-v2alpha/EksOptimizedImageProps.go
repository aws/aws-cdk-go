package awscdkeks-v2alpha


// Properties for EksOptimizedImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   eksOptimizedImageProps := &EksOptimizedImageProps{
//   	CpuArch: eks_v2_alpha.CpuArch_ARM_64,
//   	KubernetesVersion: jsii.String("kubernetesVersion"),
//   	NodeType: eks_v2_alpha.NodeType_STANDARD,
//   }
//
// Experimental.
type EksOptimizedImageProps struct {
	// What cpu architecture to retrieve the image for (arm64 or x86_64).
	// Default: CpuArch.X86_64
	//
	// Experimental.
	CpuArch CpuArch `field:"optional" json:"cpuArch" yaml:"cpuArch"`
	// The Kubernetes version to use.
	// Default: - The latest version.
	//
	// Experimental.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// What instance type to retrieve the image for (standard or GPU-optimized).
	// Default: NodeType.STANDARD
	//
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
}

