package awsappmesh


// Interface with properties ncecessary to import a reusable VirtualRouter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//
//   virtualRouterAttributes := &virtualRouterAttributes{
//   	mesh: mesh,
//   	virtualRouterName: jsii.String("virtualRouterName"),
//   }
//
type VirtualRouterAttributes struct {
	// The Mesh which the VirtualRouter belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `field:"required" json:"virtualRouterName" yaml:"virtualRouterName"`
}

