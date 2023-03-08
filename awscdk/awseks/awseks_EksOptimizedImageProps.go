package awseks


// Properties for EksOptimizedImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImageProps := &eksOptimizedImageProps{
//   	cpuArch: awscdk.Aws_eks.cpuArch_ARM_64,
//   	kubernetesVersion: jsii.String("kubernetesVersion"),
//   	nodeType: awscdk.*Aws_eks.nodeType_STANDARD,
//   }
//
// Experimental.
type EksOptimizedImageProps struct {
	// What cpu architecture to retrieve the image for (arm64 or x86_64).
	// Experimental.
	CpuArch CpuArch `field:"optional" json:"cpuArch" yaml:"cpuArch"`
	// The Kubernetes version to use.
	// Experimental.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// What instance type to retrieve the image for (standard or GPU-optimized).
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
}

