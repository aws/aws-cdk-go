package awsappmesh


// A reference to a VirtualRouter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterReference := &VirtualRouterReference{
//   	VirtualRouterArn: jsii.String("virtualRouterArn"),
//   	VirtualRouterId: jsii.String("virtualRouterId"),
//   }
//
type VirtualRouterReference struct {
	// The ARN of the VirtualRouter resource.
	VirtualRouterArn *string `field:"required" json:"virtualRouterArn" yaml:"virtualRouterArn"`
	// The Id of the VirtualRouter resource.
	VirtualRouterId *string `field:"required" json:"virtualRouterId" yaml:"virtualRouterId"`
}

