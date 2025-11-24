package mixinsawsappmesh


// An object that represents the action to take if a match is determined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpGatewayRouteActionProperty := &HttpGatewayRouteActionProperty{
//   	Rewrite: &HttpGatewayRouteRewriteProperty{
//   		Hostname: &GatewayRouteHostnameRewriteProperty{
//   			DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   		},
//   		Path: &HttpGatewayRoutePathRewriteProperty{
//   			Exact: jsii.String("exact"),
//   		},
//   		Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   			DefaultPrefix: jsii.String("defaultPrefix"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: &GatewayRouteTargetProperty{
//   		Port: jsii.Number(123),
//   		VirtualService: &GatewayRouteVirtualServiceProperty{
//   			VirtualServiceName: jsii.String("virtualServiceName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteaction.html
//
type CfnGatewayRoutePropsMixin_HttpGatewayRouteActionProperty struct {
	// The gateway route action to rewrite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteaction.html#cfn-appmesh-gatewayroute-httpgatewayrouteaction-rewrite
	//
	Rewrite interface{} `field:"optional" json:"rewrite" yaml:"rewrite"`
	// An object that represents the target that traffic is routed to when a request matches the gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteaction.html#cfn-appmesh-gatewayroute-httpgatewayrouteaction-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

