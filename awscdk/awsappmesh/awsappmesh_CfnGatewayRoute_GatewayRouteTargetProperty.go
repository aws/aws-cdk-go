package awsappmesh


// An object that represents a gateway route target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteTargetProperty := &gatewayRouteTargetProperty{
//   	virtualService: &gatewayRouteVirtualServiceProperty{
//   		virtualServiceName: jsii.String("virtualServiceName"),
//   	},
//
//   	// the properties below are optional
//   	port: jsii.Number(123),
//   }
//
type CfnGatewayRoute_GatewayRouteTargetProperty struct {
	// An object that represents a virtual service gateway route target.
	VirtualService interface{} `field:"required" json:"virtualService" yaml:"virtualService"`
	// `CfnGatewayRoute.GatewayRouteTargetProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

