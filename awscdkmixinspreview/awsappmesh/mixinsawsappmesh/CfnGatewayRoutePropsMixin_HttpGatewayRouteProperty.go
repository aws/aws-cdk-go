package mixinsawsappmesh


// An object that represents an HTTP gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpGatewayRouteProperty := &HttpGatewayRouteProperty{
//   	Action: &HttpGatewayRouteActionProperty{
//   		Rewrite: &HttpGatewayRouteRewriteProperty{
//   			Hostname: &GatewayRouteHostnameRewriteProperty{
//   				DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   			Path: &HttpGatewayRoutePathRewriteProperty{
//   				Exact: jsii.String("exact"),
//   			},
//   			Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   				DefaultPrefix: jsii.String("defaultPrefix"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Target: &GatewayRouteTargetProperty{
//   			Port: jsii.Number(123),
//   			VirtualService: &GatewayRouteVirtualServiceProperty{
//   				VirtualServiceName: jsii.String("virtualServiceName"),
//   			},
//   		},
//   	},
//   	Match: &HttpGatewayRouteMatchProperty{
//   		Headers: []interface{}{
//   			&HttpGatewayRouteHeaderProperty{
//   				Invert: jsii.Boolean(false),
//   				Match: &HttpGatewayRouteHeaderMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   					Range: &GatewayRouteRangeMatchProperty{
//   						End: jsii.Number(123),
//   						Start: jsii.Number(123),
//   					},
//   					Regex: jsii.String("regex"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Hostname: &GatewayRouteHostnameMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Suffix: jsii.String("suffix"),
//   		},
//   		Method: jsii.String("method"),
//   		Path: &HttpPathMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Regex: jsii.String("regex"),
//   		},
//   		Port: jsii.Number(123),
//   		Prefix: jsii.String("prefix"),
//   		QueryParameters: []interface{}{
//   			&QueryParameterProperty{
//   				Match: &HttpQueryParameterMatchProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroute.html
//
type CfnGatewayRoutePropsMixin_HttpGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroute.html#cfn-appmesh-gatewayroute-httpgatewayroute-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroute.html#cfn-appmesh-gatewayroute-httpgatewayroute-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

