package awsappmesh


// An object that represents the gateway route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteRewriteProperty := &GrpcGatewayRouteRewriteProperty{
//   	Hostname: &GatewayRouteHostnameRewriteProperty{
//   		DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayrouterewrite.html
//
type CfnGatewayRoute_GrpcGatewayRouteRewriteProperty struct {
	// The host name of the gateway route to rewrite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayrouterewrite.html#cfn-appmesh-gatewayroute-grpcgatewayrouterewrite-hostname
	//
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
}

