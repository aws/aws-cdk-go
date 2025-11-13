package interfacesawsappmesh


// A reference to a GatewayRoute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteReference := &GatewayRouteReference{
//   	GatewayRouteArn: jsii.String("gatewayRouteArn"),
//   	GatewayRouteId: jsii.String("gatewayRouteId"),
//   }
//
type GatewayRouteReference struct {
	// The ARN of the GatewayRoute resource.
	GatewayRouteArn *string `field:"required" json:"gatewayRouteArn" yaml:"gatewayRouteArn"`
	// The Id of the GatewayRoute resource.
	GatewayRouteId *string `field:"required" json:"gatewayRouteId" yaml:"gatewayRouteId"`
}

