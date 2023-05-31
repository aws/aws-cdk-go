package awsappmesh


// An object that represents a gateway route target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteTargetProperty := &GatewayRouteTargetProperty{
//   	VirtualService: &GatewayRouteVirtualServiceProperty{
//   		VirtualServiceName: jsii.String("virtualServiceName"),
//   	},
//
//   	// the properties below are optional
//   	Port: jsii.Number(123),
//   }
//
type CfnGatewayRoute_GatewayRouteTargetProperty struct {
	// An object that represents a virtual service gateway route target.
	VirtualService interface{} `field:"required" json:"virtualService" yaml:"virtualService"`
	// The port number of the gateway route target.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

