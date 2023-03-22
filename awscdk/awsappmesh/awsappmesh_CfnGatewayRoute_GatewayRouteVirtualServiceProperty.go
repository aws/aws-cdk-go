package awsappmesh


// An object that represents the virtual service that traffic is routed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteVirtualServiceProperty := &gatewayRouteVirtualServiceProperty{
//   	virtualServiceName: jsii.String("virtualServiceName"),
//   }
//
type CfnGatewayRoute_GatewayRouteVirtualServiceProperty struct {
	// The name of the virtual service that traffic is routed to.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
}

