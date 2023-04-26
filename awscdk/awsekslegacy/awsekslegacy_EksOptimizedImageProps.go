package awsekslegacy


// Properties for EksOptimizedImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImageProps := &EksOptimizedImageProps{
//   	KubernetesVersion: jsii.String("kubernetesVersion"),
//   	NodeType: awscdk.Aws_eks_legacy.NodeType_STANDARD,
//   }
//
// Experimental.
type EksOptimizedImageProps struct {
	// The Kubernetes version to use.
	// Experimental.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// What instance type to retrieve the image for (standard or GPU-optimized).
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
}

