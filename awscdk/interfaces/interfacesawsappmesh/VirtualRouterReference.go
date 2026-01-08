package interfacesawsappmesh


// A reference to a VirtualRouter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterReference := &VirtualRouterReference{
//   	VirtualRouterArn: jsii.String("virtualRouterArn"),
//   }
//
type VirtualRouterReference struct {
	// The Arn of the VirtualRouter resource.
	VirtualRouterArn *string `field:"required" json:"virtualRouterArn" yaml:"virtualRouterArn"`
}

