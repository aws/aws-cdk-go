package awsappmesh


// Interface with properties necessary to import a reusable GatewayRoute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualGateway virtualGateway
//
//   gatewayRouteAttributes := &GatewayRouteAttributes{
//   	GatewayRouteName: jsii.String("gatewayRouteName"),
//   	VirtualGateway: virtualGateway,
//   }
//
// Experimental.
type GatewayRouteAttributes struct {
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `field:"required" json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	// Experimental.
	VirtualGateway IVirtualGateway `field:"required" json:"virtualGateway" yaml:"virtualGateway"`
}

