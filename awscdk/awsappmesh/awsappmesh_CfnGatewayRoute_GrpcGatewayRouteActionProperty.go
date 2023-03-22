package awsappmesh


// An object that represents the action to take if a match is determined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteActionProperty := &grpcGatewayRouteActionProperty{
//   	target: &gatewayRouteTargetProperty{
//   		virtualService: &gatewayRouteVirtualServiceProperty{
//   			virtualServiceName: jsii.String("virtualServiceName"),
//   		},
//
//   		// the properties below are optional
//   		port: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	rewrite: &grpcGatewayRouteRewriteProperty{
//   		hostname: &gatewayRouteHostnameRewriteProperty{
//   			defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   		},
//   	},
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteActionProperty struct {
	// An object that represents the target that traffic is routed to when a request matches the gateway route.
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The gateway route action to rewrite.
	Rewrite interface{} `field:"optional" json:"rewrite" yaml:"rewrite"`
}

