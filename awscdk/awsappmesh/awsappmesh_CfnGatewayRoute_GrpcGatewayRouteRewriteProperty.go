package awsappmesh


// An object that represents the gateway route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteRewriteProperty := &grpcGatewayRouteRewriteProperty{
//   	hostname: &gatewayRouteHostnameRewriteProperty{
//   		defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   	},
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteRewriteProperty struct {
	// The host name of the gateway route to rewrite.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
}

