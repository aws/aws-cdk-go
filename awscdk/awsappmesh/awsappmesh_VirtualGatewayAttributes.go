package awsappmesh


// Unterface with properties necessary to import a reusable VirtualGateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//
//   virtualGatewayAttributes := &virtualGatewayAttributes{
//   	mesh: mesh,
//   	virtualGatewayName: jsii.String("virtualGatewayName"),
//   }
//
type VirtualGatewayAttributes struct {
	// The Mesh that the VirtualGateway belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualGateway.
	VirtualGatewayName *string `field:"required" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

