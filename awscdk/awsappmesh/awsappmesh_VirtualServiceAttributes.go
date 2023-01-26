package awsappmesh


// Interface with properties ncecessary to import a reusable VirtualService.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//
//   virtualServiceAttributes := &virtualServiceAttributes{
//   	mesh: mesh,
//   	virtualServiceName: jsii.String("virtualServiceName"),
//   }
//
type VirtualServiceAttributes struct {
	// The Mesh which the VirtualService belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualService, it is recommended this follows the fully-qualified domain name format.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
}

